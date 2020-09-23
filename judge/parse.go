package judge

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var suffix = map[string]string{
	"0": ".c",
	"1": ".cpp",
	"2": ".py",
}

var beforeScript = map[string]string{
	"0": "gcc $%path%$$%full%$ -o $%path%$$%file%$.exe -std=c11",
	"1": "c++ $%path%$$%full%$ -o $%path%$$%file%$.exe -std=c++11",
	"2": "",
}

var script = map[string]string{
	"0": "$%path%$$%file%$.exe",
	"1": "$%path%$$%file%$.exe",
	"2": "python3 $%path%$$%file%$.py",
}

var afterScript = map[string]string{
	"0": "rm $%path%$$%file%$.exe",
	"1": "rm $%path%$$%file%$.exe",
	"2": "",
}

func getSampleID(ProID string) (samples []string) {
	Path := "./data/" + ProID
	paths, err := ioutil.ReadDir(Path)
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`(\d+).in`)
	for _, path := range paths {
		name := path.Name()
		tmp := reg.FindSubmatch([]byte(name))
		if tmp != nil {
			idx := string(tmp[1])
			ans := Path + "/" + fmt.Sprintf("%v.out", idx)
			if _, err := os.Stat(ans); err == nil {
				samples = append(samples, idx)
			}
		}
	}
	return
}

func plain(raw []byte) string {
	buf := bufio.NewScanner(bytes.NewReader(raw))
	var b bytes.Buffer
	newline := []byte{'\n'}
	for buf.Scan() {
		b.Write(bytes.TrimSpace(buf.Bytes()))
		b.Write(newline)
	}
	return b.String()
}

func splitCmd(s string) (res []string) {
	// https://github.com/vrischmann/shlex/blob/master/shlex.go
	var buf bytes.Buffer
	insideQuotes := false
	for _, r := range s {
		switch {
		case unicode.IsSpace(r) && !insideQuotes:
			if buf.Len() > 0 {
				res = append(res, buf.String())
				buf.Reset()
			}
		case r == '"' || r == '\'':
			if insideQuotes {
				res = append(res, buf.String())
				buf.Reset()
				insideQuotes = false
				continue
			}
			insideQuotes = true
		default:
			buf.WriteRune(r)
		}
	}
	if buf.Len() > 0 {
		res = append(res, buf.String())
	}
	return
}

// Parse and start test
func Parse(ProID string, lang string, filename string) (string, string, uint64, uint64) {
	return Test(beforeScript[lang], script[lang], afterScript[lang], filename, ProID)
}

func getLimit(ProID string) (memory, time uint64) { //需要加入数据库
	memory = 1024 * 1024 * 256
	time = 1
	file, err := os.Open("./data/limits/" + ProID + ".lim")
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(file)
	file.Close()
	if err != nil {
		return
	}
	s := strings.Split(string(data), " ")
	m, _ := strconv.Atoi(s[0])
	t, _ := strconv.Atoi(s[1])
	memory = uint64(m) * 1024 * 1024
	time = uint64(t)
	return
}

//ParseMemory for formating the memory using
func ParseMemory(memory uint64) string {
	if memory > 1024*1024 {
		return fmt.Sprintf("%.3fMB", float64(memory)/1024.0/1024.0)
	} else if memory > 1024 {
		return fmt.Sprintf("%.3fKB", float64(memory)/1024.0)
	}
	return fmt.Sprintf("%vB", memory)
}
