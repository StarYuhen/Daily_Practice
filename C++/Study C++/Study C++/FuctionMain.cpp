// ���ڼ�¼�����º�����ѧϰ�ļ�

#include <iostream>

using namespace std;

// ���庯��
void simple();
void cheers(int);
string simple(string);
int sumArray(int[], int);
string sumArray(int*, int, string);

//// ����һ���򵥺���
//int main() {
//	// ���壬�ṩԭ�ͣ�����
//	simple();
//	simple("����");
//	int cookie[8] = { 1,2,4,6,8,16,32,64 };
//	cout << "ǰ��λ�����Ϊ�� " << sumArray(cookie, 8) << endl;
//	int* cookies = new int[8] { 1, 2, 4, 6, 8, 16, 32, 64};
//	sumArray(cookies, 8, "test");
//	// ����һ������ָ��
//	// �����(*open)(int[], int)��Ϊ������ֻ�뺯����ָ�룬��*open (int[], int) ��������Ƿ���ָ��ĺ�����
//	int (*open)(int[], int) = sumArray;
//	cout << "���ú���ָ�������Ϊ�� " << (*open)(cookies, 6) << endl;
//
//}

// ��������
void simple() {
	cout << "���Ǻ���simple" << endl;
	cheers(10);
	// ʹ��arr
}

// ���ݺ�����ע������
void cheers(int n) {
	for (int i = 0; i < n; i++) {
		cout << "Cheers!  ";
	}
	cout << endl;
}

// ���Է��غ�������������simple
string simple(string name) {
	cout << "�ҵ�������: " << name << endl;
	return name;
}


// ��������ǰn��Ԫ�صĺ�
int sumArray(int arr[], int n) {
	int sum = 0;
	for (int j = 0; j < n; j++) {
		sum += arr[j];
	}
	return sum;
}

// ���غ�����ͬʱʹ��ָ��
string sumArray(int* arr, int n, string name) {
	int sum = 0;
	for (int j = 0; j < n; j++) {
		sum += *(arr + j);
	}
	cout << "ָ������ļ�����: " << sum << endl;
	// cout << "ָ������ĳ���: " << sizeof(arr) << endl;
	return name;
}