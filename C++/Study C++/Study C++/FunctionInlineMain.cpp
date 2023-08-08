// 学习弟八章的内容


#include <iostream>

using namespace std;

// 声明一个函数内联
inline double square(double x) { return x * x; }
// 原始实现函数内联
#define SQUARE(x) x * x

// 结构指针的领用
struct name {
	int age;
	string name;
};

const name& use(name& Name);

// 带有默认值的函数
void SetName(string name = "yuhen");

// 函数模版
//template <class Any>
//void Swap(Any& a, Any& b);

// 函数模版
template <class Any>
void Swap(Any& a, Any& b) {
	Any temp = a;
	a = b;
	b = temp;
}

// 函数模版重载,用于交换数组的元素
template <typename T>
void Swap(T a[], T b[], int n) {
	T temp;
	for (int i = 0; i < n; i++) {
		temp = a[i];
		a[i] = b[i];
		b[i] = temp;
	}
}

// 函数模版具体化
template <> void Swap<name>(name& a, name& b) {
	int t1;
	string t2;
	t2 = a.name;
	a.name = b.name;
	b.name = t2;
	t1 = a.age;
	a.age = b.age;
	a.age = t1;

}


//int main() {
//	double a = square(10);
//	//cout << a << endl;
//	//double b = SQUARE(10);
//	//cout << b << endl;
//	// 当变量变成指针时，能够被其他变量修改，无论有多少层，依旧会被修改。
//	double* c = &a;
//	double* d = c;
//	a = 200;
//	cout << *d << endl;
//
//	name Name = {
//		10,
//		"yuhen"
//	};
//	use(Name);
//	cout << "name的age: " << Name.age << endl;
//	// 使用默认值函数
//	SetName();
//	SetName("StarYuhen");
//	int j = 10, k = 20;
//	int j1[2] = { 1,2 };
//	int k1[2] = { 3,4 };
//	Swap(j, k);
//	Swap(j1, k1, 1);
//	cout << "函数模版: " << j << " " << k << endl;
//	cout << "函数模版重载: " << j1[0] << " " << k1[0] << endl;
//
//	name Name2 = {
//		15,
//		"StarYuhen"
//	};
//	Swap(Name, Name2);
//	cout << "利用显式化交换结构体： " << Name.name << " " << Name2.name << endl;
//
//}

// 使用const临时变量
void refcube(const double& rs) {
	cout << rs * rs * rs << endl;
}

const name& use(name& Name) {
	cout << "name的age: " << Name.age << endl;
	Name.age++;
	return Name;
}

void SetName(string name) {
	cout << "Hello! " << name << endl;
}




