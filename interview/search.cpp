#include <iostream>
#include <stdlib.h>
#include <time.h>
#include <vector>
#include <cmath>
#include <algorithm>
using namespace std;

int bsearch(const vector<int>& v, int val) {
    int left = 0;
    int right = v.size() - 1;
    while (left <= right) {
        int cur_idx = floor( (left + right) / 2.0);
        int cur = v[cur_idx];
        if (cur == val) return cur_idx;
        if (cur < val) left = cur_idx + 1;
        if (cur > val) right = cur_idx - 1;
    }
    return -1;
}

int main() {
    srand(time(NULL));
    vector<int> v;

    for (int i = 0; i < 10; ++i) {
        int r = rand() % 10;
        v.push_back(r);
    }
    cout << endl;
    sort(v.begin(), v.end());
    for (const auto& val : v) {
        cout << val << " ";
    }
    cout << endl;

    int search;
    while (cin >> search) {
        cout << bsearch(v, search);
        cout << endl;
    }
}
