// ��һ���Ѷȿ�ʼ������

#include <iostream>
#include <string>


using namespace std;

// �����cin��ʹ��
// 
//int main() {
//	const int size = 15;
//	char name[size];
//	char dessert[size];
//
//	// cin.getline ���ڶ�ȡһ�е����ݣ���Ϊ���ʹ��cin��ȡ���룬�����ո�ͻ�ֹͣ�������������������ݣ�����ǰ�����벢δ��ȡ
// // ��ͨ����ȡ���з����ж��Ƿ��ȡ��ϣ����ÿ��ַ�������з�
// // ��cin.get ���ǿ��Դ����з�����������get�����������з�
//
//
//	cout << "�������������: \n";
//	cin.getline(name, size);
//	cout << "��������ϲ����ˮ��: \n";
//	cin.getline(dessert, size);
//	cout << "���������: " << name;
//	cout << ", ��ϲ����ˮ����: " << dessert << endl;
//}



// ʹ�ýṹ��,��Go�Ľṹ�����
// �ýṹ�帳��һ����������ʱ������ֱ��ʹ�ýṹ������������Ҫʹ��struct���� struct inflatable test ; ����ֱ��д�� inflatable test;
struct inflatable {
	char name[20];
	float volume;
	double price;
	// ����λ�ֶΣ���ͨ��λ������ṹ���Ա�Ĵ����С
	int test : 4;
};

// ������
// ��������һ�����ݸ�ʽ�����ܹ����治ͬ���������ͣ���ֻ��ͬʱ�������е�һ������
// ���仰��˵�������彫����ṹ���int,long,double��Ϊͬһ���ڴ棬��ʹ��һ�����ʹ���ʱ���������͵�ֵ���Ḳ�ǡ�
union Test {
	int intValue;
	long longValue;
	double doubleValue;
};

// ����ʹ�ù�����
//int main() {
//	// ���������ҵ��������൱����һ������ӵ�ж�����͵�ѡ�����Ƿ��ͣ����ܴ��������ͣ���ȻҲ��ֻ����һ��һ��
//	Test test;
//	test.intValue = 10;
//	cout << "�����������:" << test.intValue << endl;
//	// �ٸ�ֵlongValue,��Ӧ����intValue��ֵ,��Ӧ����100
//	test.longValue = 100;
//	cout << "�����������:" << test.intValue << endl;
//}


// ö��
// ö����һ�ִ������ų����ķ�ʽ������һ�����ͳ���������ͨ��ö����������
// ö�ٵĵ�һ����Ա��Ĭ��ֵΪ0��������Ա��ֵ��ǰһ����Ա�Ļ����ϼ�1
enum spectrum {
	red, orange, yellow, green, blue, violet, indigo, ultraviolet
};



// ���ﶨ����ö�ٵ�ȡֵ��Χ�����ǲ���ö��ֵҲ�ܸ�ֵ��ö�ٱ���
enum bits {
	one = 1, two = 2, four = 4, eight = 8
};


