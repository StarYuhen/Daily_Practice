// ��ҵ6�� һ��car�Ľṹ�����������������̣�������ݣ���������Ϣ��Ȼ����ʾ����

#include <iostream>
#include <string>

using namespace std;


struct Car {
	string make;
	int year;
};

//int main() {
//	int length;
//	cout << "How many cars do you wish to catalog?";
//	cin >> length;
//	cin.ignore();
//	Car* car = new Car[length];
//	for (int j = 0; j < length; j++) {
//		cout << "Car #" << j + 1 << endl;
//		cout << "Please enter the make: ";
//		getline(cin, car[j].make);
//		cout << "Please enter the year made: ";
//		cin >> car[j].year;
//		cin.ignore();
//	}
//	cout << "Here is your collection:" << endl;
//	for (int i = 0; i < length; i++) {
//		cout << car[i].year << " " << car[i].make << endl;
//	}
//}