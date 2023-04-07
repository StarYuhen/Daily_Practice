//
// Created by 玉衡 on 2023/4/7.
//

#include <ostream>
#include <string>
#include <regex>
#include <iostream>


int FlagsDefault() {
    return 10 + 10;
}

// 默认值参数
int flagsInt(int Int = FlagsDefault()) {
    return Int+1;
}

int main() {
    /*
     *  分析原理：
     *
     *
     */
    // 尝试使用regex替换值
    std::string cheese = "His X ,How are you?"; // 用于生成一个我们需要替换的字符串，其中X是我们需要替换的值
    std::regex reg("X"); // 用于申明一个正则对象
    std::string result = std::regex_replace(cheese, reg, "StarYuhen"); // 用于进行对象替换
    std::cout << result << std::endl; // 最后输出


    // 输出默认值的函数
    std::cout << flagsInt(10) << std::endl;
}




