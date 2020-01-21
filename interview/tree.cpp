#include <iostream>
#include <queue>
using namespace std;

struct Node {
    int val;
    Node * left = nullptr;
    Node * right = nullptr;

    Node(int val_in) {
        val = val_in;
    }

    ~Node() {
        if (left) delete left;
        if (right) delete right;
    }
};

Node * initTree() {
    Node * root = new Node(0);
    Node * num1 = new Node(1);
    Node * num2 = new Node(2);
    Node * num3 = new Node(3);
    Node * num4 = new Node(4);
    Node * num5 = new Node(5);
    Node * num6 = new Node(6);

    root->left = num1;
    num1->left = num2;
    num1->right = num3;
    num2->left = num4;

    root->right = num5;
    num5->right = num6;

    /*      0
     *   1     5
     *  2  3     6
     * 4
    */

    return root;
}

void inorder(Node* root) {
    if (!root) {
        return;
    }
    inorder(root->left);
    cout << root->val << " ";
    inorder(root->right);
    return;
}

void preorder(Node* root) {
    if (!root) {
        return;
    }
    cout << root->val << " ";
    preorder(root->left);
    preorder(root->right);
    return;
}

void postorder(Node* root) {
    if (!root) {
        return;
    }
    postorder(root->left);
    postorder(root->right);
    cout << root->val << " ";
    return;
}

void bfs(Node * root) {
    queue<Node *> q;
    q.push(root);
    while (!q.empty()) {
        Node * cur = q.front();
        cout << cur->val << " ";
        if (cur->left) q.push(cur->left);
        if (cur->right) q.push(cur->right);
        q.pop();
    }
    cout << endl;
    return;
}

int main() {
    Node * root = initTree();
    inorder(root);
    cout << endl;
    preorder(root);
    cout << endl;
    postorder(root);
    cout << endl;
    bfs(root);
    delete root;
}
