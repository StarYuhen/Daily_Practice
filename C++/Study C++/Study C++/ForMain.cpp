

#include <iostream>
#include <cstring>
#include <ctime>

using namespace std;

// ͨ��Ԥ�������������
// ��char�ı�������Ϊchat(Ц)
#define chat char

// ͨ��typedef��������
typedef  char byte;
// Ҳ���Ը���Ϊָ��
typedef char * pstring;


//// ���ڲ���C���ַ��Ƚϣ�C���ַ���ͨ����Ϊ��ַ
//// ��ϵ�������������Ƚ��ַ�����Ϊ�ַ�ʵ����������
//// Cstring�ַ�����ͨ����β��ֵ�ַ�����ģ��������������鳤��
//int main() {
//	// ʹ��strcmp���ȶ��ַ������ݣ�֪�������ַ�����ͬΪֹ
//	char word[5] = "?hat";
//	// ��C++��,'a'��ʾ�ַ���"a"��ʾ�ַ�����
//	/*
//	'' ���ڱ�ʾ�ַ�������ֻ�ܰ���һ���ַ���
//	"" ���ڱ�ʾ�ַ������������԰���һ�������ַ������Կ��ַ� \0 ��β��
//	*/
//	for (char ch = 'a'; strcmp(word, "what"); ch++) {
//		cout << word << endl;
//		word[0] = ch;
//	}
//	cout << "After loop ends,word is " << word << endl;
//}


// �����ӳ�ѭ��
//int main() {
//	cout << "�������ӳ�ʱ�䣬��λΪ�룺";
//	float secs;
//	cin >> secs;
//	// CLOCKS_PER_SEC ��ʾÿ���ʱ������Ҳ��ʱ�ӵδ���
//	clock_t delay = secs * CLOCKS_PER_SEC;
//	cout << "starting\a\n";
//	// ��ȡ��ǰʱ�ӵδ���
//	clock_t start = clock();
//	// ��ǰ��������ȥ��ʼ���������С���趨���ӳ�ʱ�䣬�ͼ���ѭ��
//	while (clock() - start < delay);
//	cout << "done \a\n";
//}

