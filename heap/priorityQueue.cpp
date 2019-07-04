#include<cstdio>
#include<cstring>
#include<algorithm>
using namespace std;
#define MAX 2000000
#define INFTY (1<<30)
int H, A[MAX + 1];

int extraction() {
  int maxv;
  if (H < 1) return -INFTY;
  maxv = A[1];
  A[1] = A[H--];
  maxHeapify(1);
  return maxv;
}

void increaseKey(int i, int key) {
  if ( key < A[i]) return;
  A[i] = key;
  while (i > 1 && A[i/2] < A[i]) {
    swap(A[i], A[i /2]);
    i = i / 2;
  }
}

void insert(int key) {
  H++;
  A[H] = - INFTY;
  increaseKey(H, key);
}
