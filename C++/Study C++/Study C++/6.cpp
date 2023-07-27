// 作业6： 要求用户输入小时数和分钟数，将值传递给一个void函数，而后显示内容


#include <iostream>

void TimeNumberCout(int hours, int minutes) {
	std::cout << "Enter the number of hours: " << hours << std::endl;
	std::cout << "Enter the number of minutes: " << minutes << std::endl;
	// 本想使用printf 用于格式化字符串，但是想了想算了，这样安全些
	std::cout << "Time: " << hours << ":" << minutes << std::endl;
}


//int main() {
//	TimeNumberCout(9, 28);
//}