#include <iostream>
#include <sstream>


// 我们尝试简单魔改一下<<运算符,调用它的字符串会自动加上StarYuhen is： ，如果没有声明命名空间会告诉你编译器不知道调用哪个，我们可以写个命名空间
namespace YuhenStream {
//    template<typename T>
//    std::ostream &operator<<(std::ostream output, const T &value) {
//        output << 'is StarYuhen' << value;
//        return output;
//    }
    // class YuhenStreamBuf : public std::streambuf  这一句话申明YuhenStreamBuf继承std::streambuf，相当于Java的extends
    class YuhenStreamBuf : public std::streambuf {
        // public标识共有
    public:
        YuhenStreamBuf() {
            setp(buffer, buffer + bufferSize - 1);
        }
        // 保护
    protected:
        virtual int_type overflow(int_type ch = traits_type::eof()) {
            if (ch != traits_type::eof()) {
                *pptr() = ch;
                pbump(1);
            }

            return flushBuffer();
        }

        // 虚函数
        virtual int sync() {
            return flushBuffer();
        }
        // 私有
    private:
        int flushBuffer() {
            if (pbase() == pptr()) {
                return 0;
            }

            *pptr() = '\0';

            std::cout << "StarYuhen is:" << buffer;

            setp(buffer, buffer + bufferSize - 1);

            return 0;
        }

        // 静态常量
        static const int bufferSize = 64;
        char buffer[bufferSize];
    };

    // 模板函数
    template<typename T>
    YuhenStreamBuf &operator<<(YuhenStreamBuf &out, const T &value) {
        std::ostringstream ss;
        ss << value;
        std::string strValue = ss.str();
        out.sputn(strValue.c_str(), strValue.size());
        return out;
    }

    // 定义两个类型对象
    YuhenStreamBuf yuhenStreamBuf;

    std::ostream yuhenStream(&yuhenStreamBuf);
}


int main() {
    // 输出Hello,World!
    /*
     *  原理： 使用iostream库，同时导入也可以称之为iostream.h 的预处理 包含istream and ostream 两个基类，既输入/输出流
     *      std::cout << "Hello, World!" << std::endl; 这个语句
     *          是因为我们使用了(输出流)cout对象，他在iostream库中，被命名为extern ostream cout;并且设置为导出值
     *              而<<运算符是插入运算符，能将值插入左边的对象,这个语句将"Hello, World!"字符串插入std::cout的输出流对象
     *                  而std::endl则是标识一个换行符，模板语法 return flush(__os.put(__os.widen('\n')));,这是源代码,同时他采用了C+的模板语法
     *                    template<typename _CharT, typename _Traits>
                            inline basic_ostream<_CharT, _Traits>&
                            endl(basic_ostream<_CharT, _Traits>& __os)
                            { return flush(__os.put(__os.widen('\n'))); }

    扩展知识:
        <<我们第一反应是左移操作符，在C++中他使用了重载概念，这是操作符重载，和Java的重载不同的是，他能函数和操作符重载
    重载的<<源代码：
       // Partial specializations
  template<typename _Traits>
    inline basic_ostream<char, _Traits>&
    operator<<(basic_ostream<char, _Traits>& ___out, const char* __s)
    {
      if (!__s)
	___out.setstate(ios_base::badbit);
      else
	__ostream_insert(___out, __s,
			 static_cast<streamsize>(_Traits::length(__s)));
      return ___out;
    }

     其中operator<<是重载<<的运算符，同样可以operator+，则是重载+这个运算符
        这里详细讲解一下<<重载和源代码的解释
              if (!__s)
	___out.setstate(ios_base::badbit);
     这个语句告诉我们，判断传入的值是否为空，如果是空则返回baabit(标识流错误)，倘若正确的就可以开始执行：
           else
	__ostream_insert(___out, __s,
			 static_cast<streamsize>(_Traits::length()__s));
      return ___out;
        这个语句告诉我们，如果不是空则会插入值，__s(指的是传入的const char的指针)
            而后调用__ostream_insert，并调用静态成员static_cast，同时调用_Traits::length()__s)，返回__s的字符串长度
             同时使 static_cast<streamsize>将返回值改为streamsize类型
              而Traits类型是用于在编译时确定char类型的特征集（traits）的类型参数
        其中 template关键字是模板，typename定义一个类型名
           inline则是在函数调用时，将这个函数代码插入函数调用的位置，避免函数调用的开销
           而inline basic_ostream<char, _Traits>&则是标识运算符返回ostream的返回类型，同时basic_ostream<char, _Traits>&又是一个泛型，他用于输出char类型的输出流
     */

    std::cout << "Hello, World!" << std::endl;
    // 尝试使用自定义的<<重载苻
    // 引用自己自定义的命名空间
    using namespace YuhenStream;
    YuhenStream::yuhenStream << "Hello, world!" << std::endl;
    return 0;
}




