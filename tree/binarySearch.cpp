#include<cstdio>
#include<cstdlib>
#include<string>
#include<iostream>
using namespace std;

struct Node {
  int key
  Node *right, *left, *parent
};

Node *root, *NIL

void insert(int k) {
  Node *y = NIL;
  NOde *x = root;
  Node *z;

  z = (Node *)malloc(sizeof(Node));
  z -> key = k;
  z -> left = NIL;
  z -> right = NIL;

  while (x != NIL) {
    y = x;
    if (z -> key < x -> key) {
      x = x -> left;
    } else {
      x = x -> right;
    }
  }

  z -> parent = y;
  if (y == NIL) {
    root = z;
  } else {
    if (z -> key < y -> key) {
      y -> left = z;
    } else {
      y -> right = z;
    }
  }
}
