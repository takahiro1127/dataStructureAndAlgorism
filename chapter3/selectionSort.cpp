#include<stdio.h>

int selectionSort(int A[], int N) {
  int i, j, t, sw = 0, minj = j;
  for (i = 0; i < N - 1; i++ ) {
    minj = i
    for (j = i; j < N; j++) {
      if (A[j] < A[minj]) {
        minj = j
      }
    }
    t = A[i]; A[i] = A[minj]; A[minj] = t;
    if ( i != minj) sw++;
  }
  return sw;
}
