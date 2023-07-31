// 作业1 
/*
要求：让其输出指定信息，并且名字接受多个到处，且程序会下调一个成绩，他觉得自己成绩是B，现实则会是C
*/

#include <iostream>
#include <string>

//int main() {
//	using namespace std;
//	// 储存成绩
//	char grade;
//
//	cout << "What is your first name? ";
//	char FirstName[100], LastName[100];
//	cin.getline(FirstName, sizeof(FirstName));
//	cout << "What is your last name? ";
//	cin.getline(LastName, sizeof(LastName));
//
//	// 自动下移一位
//	cout << "What letter grade do you deserve? ";
//	cin >> grade;
//	// 转化为ASCII 值
//	int gradeASCII = static_cast<int>(grade);
//	// 转化为字符串
//	char Grade = static_cast<char>(gradeASCII + 1);
//
//	cout << "What is your age? ";
//	int Age;
//	cin >> Age;
//
//
//	// 输出内容
//	cout << "Name: " << FirstName <<", "<< LastName << endl;
//	cout << "Grade: " << Grade << endl;
//	cout << "Age: " << Age << endl;
//
//}