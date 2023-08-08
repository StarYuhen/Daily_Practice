// 作业4： 太长

#include <iostream>
#include <vector>

using namespace std;


const int strsize = 20;

struct Bop {
	char fullname[strsize];
	char title[strsize];
	char bopname[strsize];
	int preference;
};


//int main() {
//	vector<Bop>bop = {
//		{"Wimp Macho", "CEO", "WM", 0},
//		{"Raki Rhodes", "Junior Programmer", "RR", 1},
//		{"Celia Laiter", "Art Director", "MIPS", 2},
//		{"Hoppy Hipman", "Analyst Trainee", "HH", 1},
//		{"Pat Hand", "Junior Programmer", "LOOPY", 2}
//	};
//	while (true) {
//		cout << "Benevolent Order of Programmers Report\n"
//			<< "A. display by name\t\tB. display by title\n"
//			<< "C. display by bopname\t\tD. display by preference\n"
//			<< "Q. quit\n"
//			<< "Enter your choice: ";
//		char choice;
//		cin >> choice;
//		switch (choice) {
//		case 65:
//			// 通过vector解决过滤问题,循环太麻烦了
//			find_if(bop.begin(), bop.end(), [](Bop bop) {
//				if (bop.preference == 0) {
//					cout << bop.fullname << endl;
//				}
//			
//				});
//			break;
//		case 66:
//			find_if(bop.begin(), bop.end(), [](Bop bop) {
//				if (bop.preference == 1) {
//					cout << bop.title << endl;
//				}
//				
//				});
//			break;
//		case 67:
//			find_if(bop.begin(), bop.end(), [](Bop bop) {
//				if (bop.preference == 2) {
//					cout << bop.bopname << endl;
//				}
//				});
//			break;
//		case 81:
//			cout << "Bye!" << endl;
//			return 0;
//		}
//	}
//
//}


