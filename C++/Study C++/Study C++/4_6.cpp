// ��ҵ6���������ϣ���Ҫ����3��Ԫ�صĽṹ���飬ÿ��Ԫ�ض���CandyBar�ṹ��


#include <iostream>
#include <string>

using namespace std;

struct CandyBar {
	string  brand;
	double weight;
	int calorie;
};


//int main() {
//	// ����new�����ṹ���飬���䶯̬��
//	CandyBar* snack = new CandyBar[3];
//
//	// Ȼ���ٳ�ʼ��
//	// ����ָ���ƫ������ʼ��
//	*snack = {
//		"Mochar Munch",
//		2.3,
//		350
//	};
//	*(snack + 1) = {
//		"Star Yuhen",
//		3.2,
//		450
//	};
//	*(snack + 2) = {
//		"�� ��",
//		3.8,
//		679
//	};
//
//	// ���
//	for (int i = 0; i < 3; i++) {
//		cout << "���ݣ�" << (snack + i)->brand << ", " << (snack + i)->weight << ", " << (snack + i)->calorie << endl;
//	}
//
//
//
//}