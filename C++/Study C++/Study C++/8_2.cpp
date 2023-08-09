
#include <iostream>

using namespace std;

struct CandyBar {
	char brand[20];
	double weight;
	int calories;
};

void CandyBarFunction(const CandyBar candyBar = { "Millennium Munch",2.85,350 });

//int main() {
//	CandyBar Candybar = {
//	"StarYuhen",
//	2.85,
//	350
//	};
//	CandyBarFunction(Candybar);
//	CandyBarFunction();
//}


void CandyBarFunction(const CandyBar candyBar) {
	cout << "CandyBar 品牌名称: " << candyBar.brand << endl;
	cout << "CandyBar 重量: " << candyBar.weight << endl;
	cout << "CandyBar 卡路里: " << candyBar.calories << endl;
}