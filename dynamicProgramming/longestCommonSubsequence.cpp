#include<iostream>
#include<string>
#include<algorithm>
using namespace std;
static const int N = 1000;

int lcs(string X, string Y) {
  int c[N + 1][N + 1];
  int m = X.size();
  int n = Y.size();
  int maxl  = 0;
  X = ' ' + X;
  Y = ' ' + Y;

  for (int i = 0; i <= m; i++) {
    for (int j = 1; j <= n; j++) {
      if (X[i] == Y[j]) {
        c[i][j] = c[i - 1][j - 1] + 1;
      } else {
        c[i][j] = max(c[i - 1][j], c[i][j - 1]);
      }
      maxl = max(maxl, c[i][j]);
    }
  }
  return maxl;
}
