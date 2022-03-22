//
// Created by 玉衡 on 2022/3/21.
//

#include <cstdio>

// 使用预处理常量
#define MAX_SIZE 100
// 设置正常常量
const auto CodeString="C path E: end";

int CodeFunction(int stc){
    return stc/10;
}



int main(){
    auto Code="Name String Char Bit \n";
    printf(Code);
    printf("Code Size %d \n",sizeof(Code));
//    int input;
//    // 测试读取 & 和Go的一样，使用指针
//    printf("input int:");
//    scanf("%d",&input);
//    printf("input %d \n",input);
    printf("#define :%d,const:%s \n",(MAX_SIZE/10,CodeString));
    printf("%d",CodeFunction(100));
    return 0;
}



