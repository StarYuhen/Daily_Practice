// 作业7：声明结构体，而后用cout和cin赋值和输出，结构体：披萨的名称，直径，重量
// 作业7的同时，完成了8,9的内容

#include <iostream>
#include <string>

using namespace std;

struct Pizza {
	string name; // 名字
	int diameter; // 直径
	int weight; // 重量
};



//int main() {
//	// 可以使用vector来动态化，但是这里只用new来动态化
//	Pizza* pizza = new Pizza[4];
//	// 在pizza中使用*，即*pizza，意思是取pizza指针指向的对象首元素
//	for (int i = 0; i <= 3; i++) {
//		cout << "请输入披萨的直径：";
//		//string diameter;
//		//getline(cin, diameter);
//		//// 转化为int
//		//pizza[i].diameter = stoi(diameter);
//		cin >> pizza[i].diameter;
//		// 这样会直接跳过下面的getline，所以要用cin.ignore(),清空缓存区
//		cin.ignore();
//		cout << "请输入披萨公司的名字：";
//		getline(cin, pizza[i].name);
//		cout << "请输入披萨的重量：";
//		cin >> pizza[i].weight;
//		cin.ignore();
//		// 是否退出
//		cout << "是否退出输入？（y/n）";
//		if (cin.get() == 'y') {
//			// 回收内存
//			delete [] pizza;
//			return 0;
//		}
//		// 输出此时的pizza
//		cout << "披萨公司的名字：" << pizza[i].name << endl;
//		cout << "披萨的直径：" << pizza[i].diameter << endl;
//		cout << "披萨的重量：" << pizza[i].weight << endl;
//	}
//}
