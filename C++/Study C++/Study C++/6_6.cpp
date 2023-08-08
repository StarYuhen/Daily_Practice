// 作业6：记录捐的钱，并且分配每个捐钱的姓名和款项，程序展示所有捐款超过10000的对象，同时还会计算Patrons的标记

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
//	// 计算Patrons的标记
//	cout << "捐款大于10000的人:  " << endl;
//	for (const DonateMoney& donate : donates) {
//		if (donate.money > 10000) {
//			cout << donate.name << "  " << donate.money << endl;
//		}
//	}
//
//
//
//}