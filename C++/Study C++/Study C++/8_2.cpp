
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
	cout << "CandyBar Ʒ������: " << candyBar.brand << endl;
	cout << "CandyBar ����: " << candyBar.weight << endl;
	cout << "CandyBar ��·��: " << candyBar.calories << endl;
}