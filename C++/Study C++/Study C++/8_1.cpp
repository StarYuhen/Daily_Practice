#include <iostream>


using namespace std;

static int functionCount = 0;

inline void StringLog(string name) {
	cout << name << endl;
}

// ���ش˺���
void StringLog(string name, int count) {
	if (functionCount != 0) {
		for (int j = 0; j < functionCount; j++) {
			cout << name << endl;
		}
	}
	cout << "���������ù�" << functionCount << "��" << endl;
	++functionCount;
}

//int main() {
//	StringLog("StarYuhen");
//	StringLog("StarYuhen", 10);
//	StringLog("StarYuhen", 10);
//}

