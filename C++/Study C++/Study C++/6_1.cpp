// ��ҵ1����ȡ�������룬֪������@���ţ����������룬������д�ַ�ת��ΪСд��Сд�ַ�ת��Ϊ��д

#include <iostream>
#include <string>
#include <cctype>

using namespace std;

// ���Ƿ��û�ж������
//int main() {
//	char name;
//	char CharName[10000];
//	cout << "����������(����@���Զ�ֹͣ): ";
//	while (cin >> name) {
//		// 64����@��ASCII��
//		if (name == 64) {
//			// ��д��ĸתСд
//			switch (islower(name)) {
//			case 0:
//				// �ж�ΪСд��ĸ���Զ�ת��д
//				name = toupper(name);
//				strcat(CharName, name);
//				break;
//			case 1:
//				// �ж�Ϊ��д��ĸ���Զ�תСд
//				name = tolower(name);
//				cout << name;
//				break;
//			}
//			break;
//		}
//		// ��������֣�ֱ������
//		if (isdigit(name)) {
//			continue;
//		}
//	}
//}

//int main() {
//	char ch;
//
//	std::cout << "�������ַ�������@����ֹͣ���룺" << std::endl;
//	// ԭ������cin��¼�����ݴ���ch��Ȼ���ж�ch�Ƿ�Ϊ@���������@�������ѭ��
//	// �����ж�ch�Ƿ�Ϊ��д������Ǵ�д����ת��ΪСд�������Сд����ת��Ϊ��д
//	// ����@��ֹ���󣬻�ֱ������������Ȼ������һ��whileѭ����std::cout << ch;������ݵ���䣬���ǲ�����Ӱ����
//	// ��Ϊ������@��ֹ�󣬲Ż�ִ��
//
//
//    while (std::cin >> ch && ch != '@') {
//        std::cout << "������ַ�: " << ch << std::endl;
//
//        if (std::isupper(ch)) {
//            ch = std::tolower(ch); // ��дתСд
//        }
//        else if (std::islower(ch)) {
//            ch = std::toupper(ch); // Сдת��д
//        }
//
//        // �Ҷ�Ϊʲô�ʼ����Ľ����ȫ���ˣ���Ϊÿ���ַ��в�û���û��з���
//        std::cout << "ת������ַ�: " << ch << std::endl;
//        std::cout << ch << std::flush;
//    }
//}