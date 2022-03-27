//
// Created by 玉衡 on 2022/3/22.
//

#include <cstdio>
#include <climits>

// 题目1：
//编写一个计算球体体积的程序，其中球体半径为10 m，参考公式
//。注意，分数4/3应写为4.0f/3.0f 。（如果分数写成
//4/3 会产生什么结果？
const float PT=3.14157;
// 定义不需要申明储存空间的常量
extern const float Age=100.0f;



float RoundMAXPT(){
    float r=10;
    float src=4.0f/3.0f*PT*(r*r);
    printf("%f",sizeof(float));
    return src;
}
