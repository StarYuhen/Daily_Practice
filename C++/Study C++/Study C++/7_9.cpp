
#include <iostream>


using namespace std;


typedef double (*Operation) (double, double);
// 涉及指向函数指针的使用
double calculate(double x, double y, Operation operation);
double add(double x, double y);

//int main() {
//	double test = calculate(2.5, 10.4, add);
//	cout << "内容： " << test << endl;
//}





double calculate(double x, double y, Operation operation) {
	return operation(x, y);
}

double add(double x, double y) {
	return x + y;
}