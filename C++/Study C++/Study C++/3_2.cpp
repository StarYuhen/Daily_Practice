// 作业2： 让用户输入英尺和英寸单位的身高，并输入体重，而后计算BMI指数

#include <iostream>
#include <cmath>

//int main() {
//	using namespace std;
//	const int HeightPer = 12;
//	// 身高计算，1英寸=0,0254m，英寸转身高
//	const int HeightPerM = 0.0254;
//	// 体重计算，1kg=2.2磅，磅转体重
//	const int WeightPerKg = 2.2;
//
//	// 英尺，英寸，体重
//	int HeightFeet, HeightInch, Weight;
//
//
//	// 输入身高
//	cout << "请输入您的身高为(英尺): __\b\b\b";
//	cin >> HeightFeet;
//	cout << "还剩(英寸): __\b\b\b";
//	cin >> HeightInch;
//	// 输入体重
//	cout << "请输入您的体重为(磅): __\b\b\b";
//	cin >> Weight;
//	int height = HeightFeet * HeightPer * HeightPerM + HeightInch * HeightPerM;
//
//
//
//	// 计算结果
//	cout << "您的身高为：" << height << "米" << endl;
//	cout << "您的BMI指数： " << pow(Weight / WeightPerKg, 2) << endl;
//
//}