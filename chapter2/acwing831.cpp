#pragma GCC optimize(2)

#include <bits/stdc++.h>
using namespace std;

#define __FASTIO__ ios::sync_with_stdio(false);cin.tie(nullptr);cout.tie(nullptr);

#define int long long   // 性能影响不大
#define ll long long
#define ull unsigned long long

#define all(a) a.begin(), a.end()
#define fori(i, a, b) for(ll i=a;i<b;i++)
#define endl "\n"
#define umap unordered_map
#define pq priority_queue
#define pb push_back
#define pii pair<int,int>
#define vi vector<int>
#define vll vector<long long>
#define vvi vector<vector<int>>
#define len(x) (int)((x).size())

#define inf 1e18

template<typename T>void coutvc(vector<T> nums){for(auto x:nums){cout<<x<<" ";}cout<<endl;}
void printb(int a){cout<<bitset<sizeof(a)*8>(a)<<endl;}  // 打印数字对应的二进制

struct node {
    // int who;
    // int score;
    // int time;
    int a;
    int b;
};



// kmp 算法获取子串在母串中出现的所有下标
vector<int> kmp(string &text, string &pattern) {
    int m = pattern.length();
    vector<int> pi(m);
    int c = 0;
    for (int i = 1; i < m; i++) {
        char v = pattern[i];
        while (c && pattern[c] != v) {
            c = pi[c - 1];
        }
        if (pattern[c] == v) {
            c++;
        }
        pi[i] = c;
    }

    vector<int> res;
    c = 0;
    for (int i = 0; i < text.length(); i++) {
        char v = text[i];
        while (c && pattern[c] != v) {
            c = pi[c - 1];
        }
        if (pattern[c] == v) {
            c++;
        }
        if (c == m) {
            res.push_back(i - m + 1);
            c = pi[c - 1];
        }
    }
    return res;
}


void solve() {
    int n1; cin >> n1;
    string s1; cin >> s1;
    int n2; cin >> n2;
    string s2; cin >> s2;

    auto ans = kmp(s2, s1);

    coutvc(ans);

}

signed main(void) {
    __FASTIO__
    #ifdef ACM_DEBUG
    freopen("acm.in", "r", stdin); 
    #endif

    solve();
    // int t; cin >> t; while(t--) { solve(); }

    return 0;
}
