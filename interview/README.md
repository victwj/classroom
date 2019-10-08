# Interview

This section contains my preparation for CS technical interviews, designed
to be a quick reference.

## Data Structures

- Vector

```
#include <vector>
std::vector<int> vec;
O(1)	- vec[n]		
AO(1)	- vec.push_back(n)		
O(1)	- vec.pop_back()
O(n)	- vec.insert(it pos, n) 
O(n)	- vec.erase(it pos); vec.erase(it first, it last)
O(1)	- vec.size()
O(n)	- vec.clear()
```

- Doubly Linked List

```
#include <list>
std::list<T> list;
O(1)	- list.front()
O(1)	- list.back()
O(1)	- list.push_front(n)
O(1)	- list.push_back(n)
O(1)	- list.size()
O(1)	- list.erase()
O(1)	- list.insert(it pos, n) 
O(1)	- list.erase(it pos); vec.erase(it first, it last)
```

- Queue

```
#include <queue>
std::queue<T> q;
O(1)		- q.front()
O(1)		- q.back()
AO(1)/O(n)	- q.push(); (push_back on underlying container)
O(1)		- q.pop(); (pop_front on underlying container)
O(1)		- q.size()
```

- Stack

```
STACK
#include <stack>
std::stack<T> s;
O(1)		- q.top()
AO(1)/O(n)	- q.push()
O(1)		- q.pop()
O(1)		- q.size()
```

- Hash table

```
#include <unordered_map> 
std::unordered_map<int, double> ht;
O(1)	- ht[n] = m
```

## C++

- Range based for loop:

```
// Read-only:
for (auto x : v) // By value, makes a copy
for (const auto& x : v) // More efficient, const reference

// Write:
for (auto& x : v)
```
