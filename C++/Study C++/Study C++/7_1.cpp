// 作业1：要求用户输入两个数，知道其中一个为0，对于两个数会调用函数求调和平均数，并返回main。
// 调数平均数指的是倒数平均数的倒数，2.0*x*y/(x+y)

#include <iostream>

using namespace std;

double average(double x, double y);

//int main() {
//	while (true) {
//		double x, y;
//		cout << "请输入两个数字，输入0退出程序" << endl;
//		cout << "第一个数字： ";
//		cin >> x;
//		cin.ignore();
//		cout << "第二个数字： ";
//		cin >> y;
//		if (x == 0 || y == 0) {
//			cout << "检测到你输入了0，程序退出" << endl;
//			break;
//		}
//
//		cout << "调和平均数为：" << average(x, y) << endl;
//	}
//}


double average(double x, double y) {
	return 2.0 * x * y / (x + y);
}
