// ��ҵ7�������ṹ�壬������cout��cin��ֵ��������ṹ�壺���������ƣ�ֱ��������
// ��ҵ7��ͬʱ�������8,9������

#include <iostream>
#include <string>

using namespace std;

struct Pizza {
	string name; // ����
	int diameter; // ֱ��
	int weight; // ����
};



//int main() {
//	// ����ʹ��vector����̬������������ֻ��new����̬��
//	Pizza* pizza = new Pizza[4];
//	// ��pizza��ʹ��*����*pizza����˼��ȡpizzaָ��ָ��Ķ�����Ԫ��
//	for (int i = 0; i <= 3; i++) {
//		cout << "������������ֱ����";
//		//string diameter;
//		//getline(cin, diameter);
//		//// ת��Ϊint
//		//pizza[i].diameter = stoi(diameter);
//		cin >> pizza[i].diameter;
//		// ������ֱ�����������getline������Ҫ��cin.ignore(),��ջ�����
//		cin.ignore();
//		cout << "������������˾�����֣�";
//		getline(cin, pizza[i].name);
//		cout << "������������������";
//		cin >> pizza[i].weight;
//		cin.ignore();
//		// �Ƿ��˳�
//		cout << "�Ƿ��˳����룿��y/n��";
//		if (cin.get() == 'y') {
//			// �����ڴ�
//			delete [] pizza;
//			return 0;
//		}
//		// �����ʱ��pizza
//		cout << "������˾�����֣�" << pizza[i].name << endl;
//		cout << "������ֱ����" << pizza[i].diameter << endl;
//		cout << "������������" << pizza[i].weight << endl;
//	}
//}
