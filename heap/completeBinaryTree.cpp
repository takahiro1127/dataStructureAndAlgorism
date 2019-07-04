#include<iostream>
using namespace std;
#define MAX 100000

int parent(int i) {return 1/2;}
int left(int i) {return 2 * i;}
int right(int i) {return 2* i +1;}
