

#include <iostream>


using namespace std;



struct Box {
	char maker[40];
	float height;
	float width;
	float length;
	float volume;
};


void BoxLog(Box box);
void BoxVolume(Box* box);



//int main() {
//	Box box = {
//		"test",
//		10,
//		20,
//		30,
//		0
//	};
//	BoxLog(box);
//	BoxVolume(&box);
//
//}

void BoxLog(Box box) {
	cout << "传入结构体的值: " << box.maker << box.height << box.width << box.length << box.volume << endl;
}

void BoxVolume(Box* box) {
	box->volume = box->height * box->width * box->length;
}

