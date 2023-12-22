#include <bits/stdc++.h>

using namespace std;

int main() {
    ios::sync_with_stdio(false); cin.tie(nullptr); cout.tie(nullptr);
    
    int t;
    long long a, b, mod;
    
    cin >> t;
    while(t--) {
        cin >> a >> b >> mod;
        
        long long ans = 1;
        for(; b; b /= 2) {
            if(b%2) ans = ans * a % mod;
            a = a * a % mod;
        }
        cout << ans << endl;
    }
    
    return 0;
}