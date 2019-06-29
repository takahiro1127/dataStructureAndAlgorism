# include<iostream>
# include<algorithm>
using namespace std;
static const int MAX = 200000;

int() {
  int R[MAX], n;
  cin >> n; //配列の長さを与えてくれている
  // iostream 使って標準入力からnに代入している
  for ( int i =0; i < n; i++ ) cin >> R[i]; //配列に順番に入れている

  int maxv = -2000000000; //十分小さい値に設定
  int minv = R[0]

  for (int i = 1; i < n; i++) {
    maxv = max(maxv, R[i] - minv);//直前までの差と今回の値との差の間で大きい方を取ってくる
    minv = min(minv, R[i]);//minの更新、今回値とそれまでの値で一番小さいのを取ってくる
  }

  cout << maxv << endl;//標準出力

  return 0;
}
