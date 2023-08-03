// 作业3:10%的利率投资了100美元，利息是当前存款的5%.....


#include <iostream>
using namespace std;

//int main() {
//	// 单利，复利
//	double simple=100,compound = 100;
//	int count = 1;
//	while (true) {
//		// Daphne 的单利投资 10% 的利率投资了100美元
//		simple += 100 * 0.1;
//		// Cleo 的复利投资 5% 的利率投资了100美元，每年利息会包括本金加往年的利息
//		compound += compound * 0.05;
//		count++;
//		if (compound > simple) {
//			cout << count << "年后，Cleo 的复利投资超过了 Daphne 的单利投资" << endl;
//			cout << "Daphne 的单利投资：" << simple << endl;
//			cout << "Cleo 的复利投资：" << compound << endl;
//			break;
//		}
//	}
//}