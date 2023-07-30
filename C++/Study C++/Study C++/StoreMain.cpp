// 这一章难度开始升级了

#include <iostream>
#include <string>


using namespace std;

// 数组和cin的使用
// 
//int main() {
//	const int size = 15;
//	char name[size];
//	char dessert[size];
//
//	// cin.getline 用于读取一行的内容，因为如果使用cin来取输入，遇到空格就会停止，而后继续输出后面内容，但当前的输入并未读取
// // 他通过读取换行符来判断是否读取完毕，并用空字符替代换行符
// // 而cin.get 则是可以处理换行符，调用两次get即可消除换行符
//
//
//	cout << "请输入你的名字: \n";
//	cin.getline(name, size);
//	cout << "请输入你喜欢的水果: \n";
//	cin.getline(dessert, size);
//	cout << "你的名字是: " << name;
//	cout << ", 你喜欢的水果是: " << dessert << endl;
//}



// 使用结构体,和Go的结构体很像
// 用结构体赋予一个变量类型时，可以直接使用结构体名，而不需要使用struct，如 struct inflatable test ; 可以直接写成 inflatable test;
struct inflatable {
	char name[20];
	float volume;
	double price;
	// 增加位字段，能通过位来定义结构体成员的储存大小
	int test : 4;
};

// 共用体
// 共用体是一种数据格式，它能够储存不同的数据类型，但只能同时储存其中的一种类型
// 换句话来说：共用体将下面结构体的int,long,double设为同一块内存，当使用一种类型储存时，其他类型的值都会覆盖。
union Test {
	int intValue;
	long longValue;
	double doubleValue;
};

// 测试使用共用体
//int main() {
//	// 共用体在我的理解里，他相当于让一个变量拥有多个类型的选择，像是泛型，都能处理多个类型，当然也就只有这一点一样
//	Test test;
//	test.intValue = 10;
//	cout << "共用体的内容:" << test.intValue << endl;
//	// 再赋值longValue,理应覆盖intValue的值,他应当是100
//	test.longValue = 100;
//	cout << "共用体的内容:" << test.intValue << endl;
//}


// 枚举
// 枚举是一种创建符号常量的方式，它是一种整型常量，可以通过枚举名来访问
// 枚举的第一个成员的默认值为0，后续成员的值在前一个成员的基础上加1
enum spectrum {
	red, orange, yellow, green, blue, violet, indigo, ultraviolet
};



// 这里定义了枚举的取值范围，就是不是枚举值也能赋值给枚举变量
enum bits {
	one = 1, two = 2, four = 4, eight = 8
};


