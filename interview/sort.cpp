#include <iostream>
#include <stdlib.h>
#include <time.h>
#include <vector>
#include <cmath>
#include <algorithm>
#include <cassert>

using namespace std;

void insertion(vector<int>& v) {
    for (size_t i = 1; i < v.size(); ++i) {
        for (size_t j = i; j > 0; --j) {
            if (v[j] < v[j - 1]) swap(v[j], v[j - 1]);
        }
    }

}

void selection(vector<int>& v) {
    size_t cur_idx = 0;
    while (cur_idx < v.size()) {
        size_t min_idx = cur_idx;
        for (size_t j = cur_idx + 1; j < v.size(); ++j) {
            if (v[min_idx] > v[j]) {
                min_idx = j;
            }
        }
        if (min_idx != cur_idx) swap(v[cur_idx], v[min_idx]);
        cur_idx++;
    }
}

void bubble(vector<int>& v) {
    bool swapped = true;
    while (swapped) {
        swapped = false;
        for (size_t i = 0; i < v.size() - 1; ++i) {
            if (v[i] > v[i + 1]) {
                swap(v[i], v[i + 1]);
                swapped = true;
            }
        }
    }
}

void merge(vector<int>& v, int start, int end) {
    vector<int> vcopy;
    vcopy.resize(v.size());
    copy(v.begin(), v.end(), vcopy.begin());
    int left = start;
    int middle = (start + end) / 2;
    int right = middle;
    int current = start;
    while (left < middle && right < end) {
        if (v[left] < v[right]) {
            vcopy[current] = v[left];
            left++;
            current++;
        }
        else {
            vcopy[current] = v[right];
            right++;
            current++;
        }
    }
    while (left < middle) {
        vcopy[current] = v[left];
        left++;
        current++;
    }
    while (right < end) {
        vcopy[current] = v[right];
        right++;
        current++;
    }
    v = vcopy;
    for (auto x : v) cout << x;
    cout << endl;
}

void mergesort(vector<int>& v, int start, int end) {
    if (start >= end - 1) {
        // end - 1 because if start 0 and end 1, the size of array is 1
        return;
    }
    cout << "mergesort on " << start << " " << end << endl;
    int middle = (start + end) / 2;
    mergesort(v, start, middle);
    mergesort(v, middle, end);
    merge(v, start, end);
}

void mergesort(vector<int>& v) {
    mergesort(v, 0, v.size());
}

int main() {
    srand(time(NULL));
    vector<int> v;
    vector<int> v_sorted;
    int size = rand() % 20;
    for (int i = 0; i < size; ++i) {
        int r = rand() % 20;
        v.push_back(r);
        v_sorted.push_back(r);
    }

    sort(v_sorted.begin(), v_sorted.end());
    mergesort(v);

    for (auto x : v_sorted) {
        cout << x << " ";
    }
    cout << endl;
    for (auto x : v) {
        cout << x << " ";
    }
    cout << endl;

    assert(v == v_sorted);


}
