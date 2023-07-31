// 作业6：内容如上，但要包含3个元素的结构数组，每个元素都是CandyBar结构。


#include <iostream>
#include <string>

using namespace std;

struct CandyBar {
	string  brand;
	double weight;
	int calorie;
};


//int main() {
//	// 先用new创建结构数组，让其动态化
//	CandyBar* snack = new CandyBar[3];
//
//	// 然后再初始化
//	// 利用指针的偏移来初始化
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
//		"玉 衡",
//		3.8,
//		679
//	};
//
//	// 输出
//	for (int i = 0; i < 3; i++) {
//		cout << "内容：" << (snack + i)->brand << ", " << (snack + i)->weight << ", " << (snack + i)->calorie << endl;
//	}
//
//
//
//}