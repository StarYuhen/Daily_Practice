// 用于学习第三章内容的代码

// 今天发现只用略过其他函数的main方法就可以直接使用，意外。


// 任何数值或指针都可以被转隐式转化为bool，任何非零都是true，任何0值后事false

#include <iostream>
#include <limits>


using namespace std;

// 获取变量的大小
//int main() {
//	using namespace std;
//	// 读取变量的大小
//	int IntValue = 10;
//	short ShortValue = 100;
//	long LongValue = 1000;
//	cout << "int max :" << INT_MAX << endl;
//	cout <<"short max :" << SHRT_MAX << endl;
//	cout << "long max :" << LONG_MAX << endl;
//
//
//	// 读取变量的大小
//	cout << "int size :" << sizeof(IntValue)<< endl;
//	cout << "short size :" << sizeof(ShortValue) << endl;
//	cout << "long size :" << sizeof(LongValue) << endl;
//
//}



// 使用define 等预处理命令，定义常量，ps：#include 也是预处理
#define Yu 0

// 熟悉预处理命令和无符号类型
//int main() {
//	// 定义无符号类型 unsigned 是unsigned int 的缩写
//	// 无符号类型的存储和原本的int类型大小一致，只不过范围不能为负数
//	// 如： short范围：-32768~32767，无符号short范围：0~65535
//	short rovertax = SHRT_MAX;
//	unsigned short rovert = rovertax;
//	// 开始测试溢出的范围
//	cout << "变量溢出+1: " << rovertax+ 1 << endl;
//	cout << "常量溢出+1: " <<rovert+1 << endl;
//	// 将其设为0
//	rovert = Yu;
//	rovertax = Yu;
//	cout << "变量溢出-1: " << rovertax - 1 << endl;
//	cout << "常量溢出-1: " << rovert - 1 << endl;
//
//
//	// 当溢出时，无符号类型会从最小值开始，而有符号类型会从最大值开始
//	// 但是不一定，主要看自己机器，在溢出时机器不会报错终止程序，而是会继续运行，所以需要注意溢出问题
//}

// 使用转义序列
//int main() {
//	cout << " \a欢迎登录程序 \"StarYuhen\" " << endl;
//	cout << "请输入序列号:_____\b\b\b\b\b";
//	long code;
//	cin >> code;
//	cout << "检验验证码: " << code << endl;
//	cout << "请按下回车键继续 \n";
//}


