
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

int n;
int a, b;

int main(void) {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    #ifdef ACM_DEBUG
    freopen("acm.txt", "r", stdin);
    #endif  

    // gcd
    cin >> n;
    while(n--) {
        cin >> a >> b;

        while(b) {
            int c = a%b;
            a = b;
            b = c;
        }
        cout << a << endl;
    }

    return 0;
}