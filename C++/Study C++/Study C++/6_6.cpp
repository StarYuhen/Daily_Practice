// ��ҵ6����¼���Ǯ�����ҷ���ÿ����Ǯ�������Ϳ������չʾ���о���10000�Ķ���ͬʱ�������Patrons�ı��

#include <iostream>
#include <string>
#include <vector>

using namespace std;

struct  DonateMoney {
	string name;
	double money;
};


//int main() {
//	vector<DonateMoney> donates = {
//		{ "Yuhen",1000 },
//		{"StarYuhen",10800},
//		{"YuhenStar",10080},
//		{"Your",10002}
//	};
//
//	// ����Patrons�ı��
//	cout << "������10000����:  " << endl;
//	for (const DonateMoney& donate : donates) {
//		if (donate.money > 10000) {
//			cout << donate.name << "  " << donate.money << endl;
//		}
//	}
//
//
//
//}