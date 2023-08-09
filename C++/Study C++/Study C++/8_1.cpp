#include <iostream>


using namespace std;

static int functionCount = 0;

inline void StringLog(string name) {
	cout << name << endl;
}

// 重载此函数
void StringLog(string name, int count) {
	if (functionCount != 0) {
		for (int j = 0; j < functionCount; j++) {
			cout << name << endl;
		}
	}
	cout << "函数被调用过" << functionCount << "次" << endl;
	++functionCount;
}

//int main() {
//	StringLog("StarYuhen");
//	StringLog("StarYuhen", 10);
//	StringLog("StarYuhen", 10);
//}

