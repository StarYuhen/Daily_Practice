// 作业7 用char读取单词，直到用户输入done，而后支出输入了多少个单词（不包括done）


#include <iostream>
#include <cstring>

using namespace std;

//int main() {
//	using namespace std;
//	char* word = new char[1000];
//	int count = 0;
//	char complate[5];
//	char Done[] = "done";
//	cout << "请输入单词，输入done结束:  ";
//	cin.getline(word, 1000);
//	for (int j = 0; j < 1000; j++) {
//		if (word[j] == '\0') {
//			
//		}
//	}
//
//
//	cout << word[4] << endl;
//
//}



const int maxSize = 30; // 假设输入的字符串最大长度为 30

//int main() {
//    char input[maxSize];
//    char done[] = "done";
//    int wordCount = 0;
//
//    cout << "请输入多个单词（以 done 结束输入）：" << endl;
//
//    // 使用循环读取单词
//    // 草，才发现cin读取内容，自动通过空格和\0l来判断单词，不用自己写循环了
//    while (cin >> input) {
//        // 使用 strcmp 函数比较输入的单词是否为 "done"
//        if (strcmp(input, done) == 0) {
//            break; // 输入为 "done" 时结束循环
//        }
//        wordCount++;
//    }
//
//    cout << "输入了 " << wordCount << " 个单词。" << endl;
//
//    return 0;
//}

