// ѧϰ�ܰ��µ�����


#include <iostream>

using namespace std;

// ����һ����������
inline double square(double x) { return x * x; }
// ԭʼʵ�ֺ�������
#define SQUARE(x) x * x

// �ṹָ�������
struct name {
	int age;
	string name;
};

const name& use(name& Name);

// ����Ĭ��ֵ�ĺ���
void SetName(string name = "yuhen");

// ����ģ��
//template <class Any>
//void Swap(Any& a, Any& b);

// ����ģ��
template <class Any>
void Swap(Any& a, Any& b) {
	Any temp = a;
	a = b;
	b = temp;
}

// ����ģ������,���ڽ��������Ԫ��
template <typename T>
void Swap(T a[], T b[], int n) {
	T temp;
	for (int i = 0; i < n; i++) {
		temp = a[i];
		a[i] = b[i];
		b[i] = temp;
	}
}

// ����ģ����廯
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
//	// ���������ָ��ʱ���ܹ������������޸ģ������ж��ٲ㣬���ɻᱻ�޸ġ�
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
//	cout << "name��age: " << Name.age << endl;
//	// ʹ��Ĭ��ֵ����
//	SetName();
//	SetName("StarYuhen");
//	int j = 10, k = 20;
//	int j1[2] = { 1,2 };
//	int k1[2] = { 3,4 };
//	Swap(j, k);
//	Swap(j1, k1, 1);
//	cout << "����ģ��: " << j << " " << k << endl;
//	cout << "����ģ������: " << j1[0] << " " << k1[0] << endl;
//
//	name Name2 = {
//		15,
//		"StarYuhen"
//	};
//	Swap(Name, Name2);
//	cout << "������ʽ�������ṹ�壺 " << Name.name << " " << Name2.name << endl;
//
//}

// ʹ��const��ʱ����
void refcube(const double& rs) {
	cout << rs * rs * rs << endl;
}

const name& use(name& Name) {
	cout << "name��age: " << Name.age << endl;
	Name.age++;
	return Name;
}

void SetName(string name) {
	cout << "Hello! " << name << endl;
}




