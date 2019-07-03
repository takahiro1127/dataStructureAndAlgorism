#include<stdio.h>
#define MAX 10000
#define NIL - 1

struct Node { int p, l, r;};
struct Node T[MAX];
int n;

void preParse(int u) {
  if (u== NIL) return;
  printf(" %d", u);
  preParse(T[u].l);
  preParse(T[u].r);
}

void inParse(int u) {
  if (u== NIL) return
  inParse(T[u].l);
  printf(" %d", u);
  inParse(T[u].r)
}

void postParse(int u){
  if (u == NIL) return;
  postParse(T[u].l);
  postParse(T[u].r);
  printf(" %d", u);
}

