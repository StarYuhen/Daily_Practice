

#include <iostream>
#include <cstring>
#include <ctime>

using namespace std;

// 通过预处理器定义别名
// 将char的别名定义为chat(笑)
#define chat char

// 通过typedef创建别名
typedef  char byte;
// 也可以更改为指针
typedef char * pstring;


//// 用于测试C类字符比较，C类字符串通常视为地址
//// 关系操作符能用来比较字符，因为字符实际上是整形
//// Cstring字符串是通过结尾空值字符定义的，而不是所在数组长度
//int main() {
//	// 使用strcmp来比对字符串内容，知道两个字符串相同为止
//	char word[5] = "?hat";
//	// 在C++中,'a'表示字符，"a"表示字符串，
//	/*
//	'' 用于表示字符常量，只能包含一个字符。
//	"" 用于表示字符串常量，可以包含一个或多个字符，并以空字符 \0 结尾。
//	*/
//	for (char ch = 'a'; strcmp(word, "what"); ch++) {
//		cout << word << endl;
//		word[0] = ch;
//	}
//	cout << "After loop ends,word is " << word << endl;
//}


// 创建延迟循环
//int main() {
//	cout << "请输入延迟时间，单位为秒：";
//	float secs;
//	cin >> secs;
//	// CLOCKS_PER_SEC 表示每秒的时钟数，也叫时钟滴答数
//	clock_t delay = secs * CLOCKS_PER_SEC;
//	cout << "starting\a\n";
//	// 获取当前时钟滴答数
//	clock_t start = clock();
//	// 当前的数，减去开始的数，如果小鱼设定的延迟时间，就继续循环
//	while (clock() - start < delay);
//	cout << "done \a\n";
//}

