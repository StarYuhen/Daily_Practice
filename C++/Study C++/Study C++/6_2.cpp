// 作业2： 编写一个程序将donation值存入一个doble数组中，遇到非数字输入时结束输出，并报告平均值和多少个数字大于平均值

#include <iostream>
#include <cctype>

using namespace std;

//int main() {
//	double donation[10];
//	int number = 0;
//	int sum = 0, average = 0, count = 0, sumcount = 0;
//	cout << "请输入数值(遇到非数字输入自动结束输入): ";
//	while (cin >> number && sumcount < 10) {
//		sum += number;
//		donation[sumcount] = number;
//		sumcount++;
//	}
//
//	// 计算数组的平均值和大于平均值的个数
//	average = sum / 10;
//
//
//
//	for (int j = 0; j < 10; j++) {
//		if (donation[j] > average) {
//			count++;
//		}
//	}
//
//
//	cout << "总数为: " << sum << endl;
//	cout << "平均数为: " << average << endl;
//	cout << "大于平均数的个数为: " << count << endl;
//}