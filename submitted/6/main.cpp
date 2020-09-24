#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
const int maxn=6e6+5;
struct EDGE{
    int v,nex;ll w;
}edge[maxn];
int head[4000006],ecnt;
inline void add_edge(int u,int v,ll w){
    edge[++ecnt]={v,head[u],w};head[u]=ecnt;
    edge[++ecnt]={u,head[v],w};head[v]=ecnt;
}
ll dis[maxn],vcnt;
const ll INF=0x3f3f3f3f3f3f3f3f;
bool vis[maxn];
inline void Dijkstra(int s){
    priority_queue<pair<ll,int>,vector<pair<ll,int> >,greater<pair<ll,int> > > PQ;
    for(int i=1;i<=vcnt;i++) dis[i]=INF;dis[s]=0;PQ.push({0,s});
    while(!PQ.empty()){
        int u=PQ.top().second;ll w=PQ.top().first;vis[u]=1;PQ.pop();
        for(int i=head[u];i;i=edge[i].nex){
            int v=edge[i].v;
            if(!vis[v]&&edge[i].w+w<dis[v]){dis[v]=edge[i].w+w;PQ.push({dis[v],v});}
        }
    }
}
int main() {
    int n,m;cin>>n>>m;vcnt=(n-1)*(m-1)*2+2;
    if(n==1){ll ans=INF,w;for(int i=1;i<=m-1;i++){scanf("%lld",&w);ans=min(ans,w);}cout<<ans<<endl;return 0;}
    if(m==1){ll ans=INF,w;for(int i=1;i<=n-1;i++){scanf("%lld",&w);ans=min(ans,w);}cout<<ans<<endl;return 0;}
    int s=vcnt-1,t=vcnt;ll w;
    for(int i=1;i<m;i++){
        scanf("%lld",&w);add_edge(i*2,t,w);
    }
    for(int i=1;i<n-1;i++){
        for(int j=1;j<m;j++){
            scanf("%lld",&w);
            add_edge((i-1)*(m-1)*2+j*2-1,i*(m-1)*2+j*2,w);
        }
    }
    for(int i=1;i<m;i++){
        scanf("%lld",&w);add_edge((n-2)*(m-1)*2+i*2-1,s,w);
    }
    for(int i=1;i<n;i++){
        scanf("%lld",&w);add_edge(s,(i-1)*(m-1)*2+1,w);
        for(int j=2;j<m;j++) {
            scanf("%lld",&w);
            add_edge((i-1)*(m-1)*2+j*2-1,(i-1)*(m-1)*2+j*2-2,w);
        }
        if(m>1) {scanf("%lld",&w);add_edge(t,i*(m-1)*2,w);}
    }
    for(int i=1;i<n;i++){
        for(int j=1;j<m;j++){
            scanf("%lld",&w);add_edge((i-1)*(m-1)*2+j*2-1,(i-1)*(m-1)*2+j*2,w);
        }
    }
    Dijkstra(s);
    cout<<dis[t]<<endl;
    return 0;
}