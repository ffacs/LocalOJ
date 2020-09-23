package judge

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/shirou/gopsutil/process"
	"github.com/xalanq/cf-tool/util"
)

func judge(ProID, sampleID, command string, timeLimit, memLimit uint64) (string, bool, uint64, uint64) {
	inPath := fmt.Sprintf("./data/%v/%v.in", ProID, sampleID)
	ansPath := fmt.Sprintf("./data/%v/%v.out", ProID, sampleID)
	input, err := os.Open(inPath)
	if err != nil {
		fmt.Println("input file not found!")
		return "UKE", false, 0, 0
	}
	var o bytes.Buffer
	output := io.Writer(&o)

	var er bytes.Buffer
	errput := io.Writer(&er)

	cmds := splitCmd(command)

	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stdin = input   //输入重定向
	cmd.Stdout = output //输出重定向
	cmd.Stderr = errput //错误输出到buffer

	if err := cmd.Start(); err != nil {
		return fmt.Sprintf("RE #%v ... %v\n%v", sampleID, err.Error(), er.String()), false, 0, 0
	}

	pid := int32(cmd.Process.Pid)
	maxMemory := uint64(0)
	ch := make(chan error)
	go func() {
		ch <- cmd.Wait()
	}()
	running := true
	after := time.After(time.Duration(timeLimit*1000+10) * time.Millisecond)
	for running {
		select {
		case err := <-ch:
			if err != nil { //输出RE
				return fmt.Sprintf("RE #%v ... %v\n%v", sampleID, err.Error(), er.String()), false, uint64(cmd.ProcessState.UserTime().Milliseconds()), maxMemory
			}
			running = false
		case <-after:
			cmd.Process.Kill()
			for { // waiting process exit
				if cmd.ProcessState != nil {
					break
				}
				time.Sleep(1 * time.Nanosecond)
			}
			anal := fmt.Sprintf("%.3fs %v\n", cmd.ProcessState.UserTime().Seconds(), ParseMemory(maxMemory))
			// return fmt.Sprintf("TLE #%v", sampleID), false, timeLimit, maxMemory
			return fmt.Sprintf("TLE #%v ...%v", sampleID, anal), false, uint64(cmd.ProcessState.UserTime().Milliseconds()), maxMemory
		default:
			p, err := process.NewProcess(pid)
			if err == nil {
				m, err := p.MemoryInfo()
				if err == nil && m.RSS > maxMemory {
					maxMemory = m.RSS
				}
				if maxMemory > memLimit {
					cmd.Process.Kill()
					for { // waiting process exit
						if cmd.ProcessState != nil {
							break
						}
						time.Sleep(1 * time.Nanosecond)
					}
					anal := fmt.Sprintf("%.3fs %v\n", cmd.ProcessState.UserTime().Seconds(), ParseMemory(maxMemory))
					return fmt.Sprintf("MLE #%v ...%v", sampleID, anal), false, uint64(cmd.ProcessState.UserTime().Milliseconds()), maxMemory

					// return fmt.Sprintf("MLE #%v", sampleID), false, uint64(cmd.ProcessState.UserTime().Milliseconds()), maxMemory
				}
			}
		}
	}
	b, err := ioutil.ReadFile(ansPath)
	if err != nil {
		b = []byte{}
	}
	ans := plain(b)
	out := plain(o.Bytes())

	anal := fmt.Sprintf("%.3fs %v\n", cmd.ProcessState.UserTime().Seconds(), ParseMemory(maxMemory))
	if out != ans {
		return fmt.Sprintf("WA #%v ...%v", sampleID, anal), false, uint64(cmd.ProcessState.UserTime().Milliseconds()), maxMemory
	}
	return fmt.Sprintf("Passed #%v ...%v", sampleID, anal), true, uint64(cmd.ProcessState.UserTime().Milliseconds()), maxMemory
}

//Test for starting judging and return judging status
func Test(BeforeScript, Script, AfterScript, filename, ProID string) (status, logs string, maxTime uint64, maxMem uint64) {
	path, full := filepath.Split(filename)
	samples := getSampleID(ProID)
	ext := filepath.Ext(filename)
	file := full[:len(full)-len(ext)]
	rand := util.RandString(8)

	filter := func(cmd string) string {
		cmd = strings.ReplaceAll(cmd, "$%rand%$", rand)
		cmd = strings.ReplaceAll(cmd, "$%path%$", path)
		cmd = strings.ReplaceAll(cmd, "$%full%$", full)
		cmd = strings.ReplaceAll(cmd, "$%file%$", file)
		return cmd
	}

	var cmdout bytes.Buffer //Compile Imformation
	var cmderr bytes.Buffer //Compile Error
	ceo := io.Writer(&cmdout)
	cee := io.Writer(&cmderr)

	run := func(script string) error {
		if s := filter(script); len(s) > 0 {
			cmds := splitCmd(s)
			cmd := exec.Command(cmds[0], cmds[1:]...)
			cmd.Stdout = ceo
			cmd.Stderr = cee
			return cmd.Run()
		}
		return nil
	}

	if err := run(BeforeScript); err != nil {
		status = "CE"
		logs = cmderr.String()
		return
	}

	memLimit, timeLimit := getLimit(ProID)

	if s := filter(Script); len(s) > 0 {
		for _, i := range samples {
			sta, pass, runtime, runmem := judge(ProID, i, s, timeLimit, memLimit)
			logs += sta
			if runtime > maxTime {
				maxTime = runtime
			}
			if runmem > maxMem {
				maxMem = runmem
			}
			if !pass {
				status = strings.Split(sta, " ")[0]
				if e := run(AfterScript); e != nil {
					fmt.Println(e)
				}
				return
			}
		}
	}
	if err := run(AfterScript); err != nil {
		fmt.Println(err)
	}
	status = "AC"
	return
}
