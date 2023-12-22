
// #define ACM_DEBUG   // comment this line when upload !!!

#pragma GCC optimize(2)

#include <bits/stdc++.h>
using namespace std;

#define LL long long
#define ULL unsigned long long

#define PII pair<int,int>
#define all(a) a.begin(), a.end()

#define umap unordered_map
#define pq priority_queue

#define vi vector<int>
#define vvi vector<vector<int>>
#define pb push_back

#define inf 0x3f3f3f3f

auto printvector = [](vector<int> nums) { for(auto x:nums) {cout << x << " ";} cout << endl;};
void printb(int a) { cout << bitset<sizeof(a)*8>(a) << endl; }  // 打印数字对应的二进制

struct node {
    // int from;
    int to;
    int val;
};

// 判断质数
int n;
int a;

bool isprime(int x) {
    if(x < 2) return false;
    for(int i = 2; i <= x / i; i++) {
        if(x % i == 0) {
            return false;
        }
    }
    return true;
}

int main(void) {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    cin >> n;
    
    while(n--) {
        cin >> a;
        if(isprime(a)) {
            cout << "Yes" << endl;
        } else {
            cout << "No" << endl;
        }
    }

    return 0;
}