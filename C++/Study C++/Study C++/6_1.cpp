// 作业1：读取键盘输入，知道遇到@符号，并回显输入，并将大写字符转化为小写，小写字符转化为大写

#include <iostream>
#include <string>
#include <cctype>

using namespace std;

// 我是废物，没有独立完成
//int main() {
//	char name;
//	char CharName[10000];
//	cout << "请输入数据(遇到@会自动停止): ";
//	while (cin >> name) {
//		// 64就是@的ASCII码
//		if (name == 64) {
//			// 大写字母转小写
//			switch (islower(name)) {
//			case 0:
//				// 判断为小写字母，自动转大写
//				name = toupper(name);
//				strcat(CharName, name);
//				break;
//			case 1:
//				// 判断为大写字母，自动转小写
//				name = tolower(name);
//				cout << name;
//				break;
//			}
//			break;
//		}
//		// 如果是数字，直接跳过
//		if (isdigit(name)) {
//			continue;
//		}
//	}
//}

//int main() {
//	char ch;
//
//	std::cout << "请输入字符，遇到@符号停止输入：" << std::endl;
//	// 原理，利用cin将录入数据存入ch，然后判断ch是否为@，如果不是@，则继续循环
//	// 而后判断ch是否为大写，如果是大写，则转换为小写，如果是小写，则转换为大写
//	// 等于@终止过后，会直接输出结果，虽然这里有一个while循环和std::cout << ch;输出内容的语句，但是并不会影响结果
//	// 因为他们在@终止后，才会执行
//
//
//    while (std::cin >> ch && ch != '@') {
//        std::cout << "输入的字符: " << ch << std::endl;
//
//        if (std::isupper(ch)) {
//            ch = std::tolower(ch); // 大写转小写
//        }
//        else if (std::islower(ch)) {
//            ch = std::toupper(ch); // 小写转大写
//        }
//
//        // 我懂为什么最开始输出的结果是全部了，因为每个字符中并没有用换行符。
//        std::cout << "转换后的字符: " << ch << std::endl;
//        std::cout << ch << std::flush;
//    }
//}