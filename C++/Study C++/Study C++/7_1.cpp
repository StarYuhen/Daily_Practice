// ��ҵ1��Ҫ���û�������������֪������һ��Ϊ0����������������ú��������ƽ������������main��
// ����ƽ����ָ���ǵ���ƽ�����ĵ�����2.0*x*y/(x+y)

#include <iostream>

using namespace std;

double average(double x, double y);

//int main() {
//	while (true) {
//		double x, y;
//		cout << "�������������֣�����0�˳�����" << endl;
//		cout << "��һ�����֣� ";
//		cin >> x;
//		cin.ignore();
//		cout << "�ڶ������֣� ";
//		cin >> y;
//		if (x == 0 || y == 0) {
//			cout << "��⵽��������0�������˳�" << endl;
//			break;
//		}
//
//		cout << "����ƽ����Ϊ��" << average(x, y) << endl;
//	}
//}


double average(double x, double y) {
	return 2.0 * x * y / (x + y);
}
