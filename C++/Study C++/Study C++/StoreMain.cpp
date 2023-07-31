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


//// 指针的尝试
//int main() {
//	int donuts = 6;
//	auto cups = 4.5;
//	// 可以用*储存指针，但是要注意，指针的类型必须和指向的类型一致，否则会报错
//	// *是指针的标识符，用于声明指针，也可以用于解除指针，&是取地址符
//	int* p_donuts = &donuts;
//	// p_cups是指针，而p_cups2只会是一个常规double，每个指针都需要有一个*作为声明指针的提取
//	// double* p_cups, p_cups2;
//	// *p_cups, p_cups2 = &cups,10.0;
//
//	cout << "donuts 的内存地址： " << &donuts << endl;
//	cout << "cups 的内存地址： " << &cups << endl;
//	cout << "p_donuts的内存地址" << *p_donuts << endl;
//
//	/*
//	donuts 的内存地址： 0000007A25EFF5C4
//	cups 的内存地址： 0000007A25EFF5E8
//	0000007A25EFF5E8是十六进制的内容，转化一下 0x3FCF18F924
//*/
//
//
//// 这里的p_cups和p_cups2的内存地址是一样的，但是p_cups是指针，而p_cups2只是一个常规的double
//// *double 无法转化为double
//// cout << "p_cups和p_cups2的内存地址" << *p_cups << p_cups2 << endl;
//
//
//
//// 再来一个常见的错误,C++指针只是用于分配储存地址的内存，不会分配储存指针指向数据的内存，为数据创建一个地址内存无法忽略
//	//long* num;
//	//*num = 100;
//	// 提示为初始化的局部变量，这就是错位，这个变量为初始化数据内存，只是申明了指针内存
//	// 解决方法使用new,为他分配内存地址。
//	// 单纯的指针只是赋予了他指针地址，但是并未赋予内存地址，所以需要用new为他分配内存地址
//	long* num = new long;
//	*num = 100;
//	//delete 删除指针所指向的内存地址，但是并不会删除指针本身
//	// 如果是指针数组，需要使用delete[]来删除
//	delete num;
//
//
//
//}


//// 尝试数组指针
//int main() {
//	int wages[2] = { 1,2 };
//	double stacks[2] = { 100.0,200.0 };
//
//	// 指针数组
//	int* wg = &wages[0];
//	double* sk = &stacks[0];
//	// 将其改为赋值成&wages[0]的指针后+1，依旧能访问其他下标，原因是指针数组的指针是连续的，且指针数组是由指针构成的数组
//
//	// 打印指针数组的内容
//	cout << "指针wages的内容" << *wg << endl;
//	cout << "指针stacks的内容" << *sk << endl;
//
//	// 指针+1呢
//	cout << "指针wages+1的内容" << *(wg + 1) << endl;
//	cout << "指针stacks+1的内容" << *(sk + 1) << endl;
//
//	// 很有意思
//	// 在数组指针中，加一相当于元素跳转到wg[1]，数组指针中原本默认为0，+1则为1这个元素下标去了
//
//	// 尝试char储存字符串，字符串通常被列为只读的数据，所以必须使用const
//	const char* name = "hello";
//	cout << "name的内容" << name << endl;
//}

// 声明函数的存在
//int testInt();



//// 尝试理解自动存储，静态存储，动态存储
//int main() {
//	// 局部变量又称自动储存，会在函数执行完自动销毁
//	int age = 10;
//	// 静态储存则是在程序执行期间一直存在，直到程序结束才会销毁，程序不是函数
//	// 可以在函数外面定义，也可以在函数里面定义这是一致的
//	static int age2 = 20;
//	cout << "age2的内容" << testInt() << endl;
//
//	// 动态储存，则是使用new键字，他会在程序执行期间一直存在，直到使用delete销毁
//	int* age3 = new int;
//	*age3 = 30;
//}



//int testInt() {
//	return  age2;
//}
