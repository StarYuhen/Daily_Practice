// 用于记录第七章函数的学习文件

#include <iostream>

using namespace std;

// 定义函数
void simple();
void cheers(int);
string simple(string);
int sumArray(int[], int);
string sumArray(int*, int, string);

//// 测试一个简单函数
//int main() {
//	// 定义，提供原型，调用
//	simple();
//	simple("张三");
//	int cookie[8] = { 1,2,4,6,8,16,32,64 };
//	cout << "前几位数组和为： " << sumArray(cookie, 8) << endl;
//	int* cookies = new int[8] { 1, 2, 4, 6, 8, 16, 32, 64};
//	sumArray(cookies, 8, "test");
//	// 声明一个函数指针
//	// 如果是(*open)(int[], int)以为着他是只想函数的指针，而*open (int[], int) 则代表他是返回指针的函数。
//	int (*open)(int[], int) = sumArray;
//	cout << "调用函数指针数组和为： " << (*open)(cookies, 6) << endl;
//
//}

// 函数本身
void simple() {
	cout << "我是函数simple" << endl;
	cheers(10);
	// 使用arr
}

// 传递函数，注意重载
void cheers(int n) {
	for (int i = 0; i < n; i++) {
		cout << "Cheers!  ";
	}
	cout << endl;
}

// 测试返回函数，并且重载simple
string simple(string name) {
	cout << "我的名字是: " << name << endl;
	return name;
}


// 计算数组前n个元素的和
int sumArray(int arr[], int n) {
	int sum = 0;
	for (int j = 0; j < n; j++) {
		sum += arr[j];
	}
	return sum;
}

// 重载函数，同时使用指针
string sumArray(int* arr, int n, string name) {
	int sum = 0;
	for (int j = 0; j < n; j++) {
		sum += *(arr + j);
	}
	cout << "指针数组的计算结果: " << sum << endl;
	// cout << "指针数组的长度: " << sizeof(arr) << endl;
	return name;
}