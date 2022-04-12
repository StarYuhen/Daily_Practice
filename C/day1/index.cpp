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
//   // printf("Running V :%d",RoundMAXPT());
//
//
//
////    auto Code="Name String Char Bit \n";
////    printf(Code);
////    printf("Code Size %d \n",sizeof(Code));
//////    int input;
//////    // 测试读取 & 和Go的一样，使用指针
//////    printf("input int:");
//////    scanf("%d",&input);
//////    printf("input %d \n",input);
////    printf("#define :%d,const:%s \n",(MAX_SIZE/10,CodeString));
////    printf("%d",CodeFunction(100));
//    return 0;

    while(true){
        int time1, time2, hour1, hour2, minute1, minute2, a, b;//定义八个整数型变量 并赋值
        printf("请输入4位整数形式的火车出发时间和到达时间（24小时制），中间用空格隔开，如12点01分出发，15点30分到达，则输入：1201 1530。请输入：");//输入提示
        scanf("%d %d", &time1, &time2);//数据输入
        hour1 = time1 / 100;//计算数据
        hour2 = time2 / 100;//计算数据
        minute1 = time1 % 100;//计算数据
        minute2 = time2 % 100;//计算数据
        a = hour2 - hour1;//计算数据
        b = minute2 - minute1;//计算数据
        if (hour1>=00 && hour1<=23 && hour2>=00 && hour2<=23 && minute1>=00 && minute1<=59 && minute2>=00 && minute2<=59)//判断输入是否合法，若输入合法运行下列程序
        {
            if (b < 0)//if判断语句 若满足判断则进行下列运算
            {
                a = a - 1;//计算数据
                b = b + 60;//计算数据
            }
            if (a < 0)//if判断语句 若满足判断则进行下列运算
            {
                a = a + 24;//计算数据
            }
            printf("%02d:%02d", a, b);//输出数据
            break;//返回一个整数0
        }
        else//若输入不合法，运行下列程序
        {
            printf("输入格式错误，程序结束，请重新运行。\n");//数据不合法输出提示
        }
    }
    return 0;
}



