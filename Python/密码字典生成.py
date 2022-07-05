# -*- coding: utf-8 -*-

import itertools


def readinformationList():
    try:
        # 读取信息
        informationFile = open('persion_info.txt', 'r')
        lines = informationFile.readlines()
        for line in lines:
            infolist.append(line.strip().split(':')[1])
        # print(infolist)
    except Exception as e:
        print(e + "\n")
        print("Read persion_info error!")


def CreatNumberList():
    # 数字元素
    words = "0123456789"
    # 产生不同数字排列，数字组合长度为3
    itertoolsNumberList = itertools.product(words, repeat=3)
    for number in itertoolsNumberList:
        # 写入数字列表备用
        numberList.append("".join(number))


def CreatSpecialList():
    # 创建特殊字符写入列表
    specialWords = "`~!@#$%^&*()?|/><,."
    for i in specialWords:
        specialList.append("".join(i))


def AddTopPwd():
    try:
        # 读取Pwd文件，存入password字典文件
        informationFile = open('pwd.txt', 'r')
        lines = informationFile.readline()
        for line in lines:
            dictionaryFile.write(line)
    except Exception as e:
        print(e + "\n")
        print("Read pwd error!")


# 字典生成算法主体
def Combination():
    for a in range(len(infolist)):
        # 把个人信息大于等于8位数的直接输出到字典
        if len(infolist[a]) >= 8:
            dictionaryFile.write(infolist[a] + "\n")
        # 小于8位数的个人信息，利用数字补全到8位数输出
        else:
            needWords = 8 - len(infolist[a])
            for b in itertools.permutations("1234567890", needWords):
                dictionaryFile.write(infolist[a] + ''.join(b) + '\n')
        # 将个人信息元素两两互相拼接，大于等于8位的输出到字典
        for c in range(0, len(infolist)):
            if len(infolist[a] + infolist[c]) >= 8:
                dictionaryFile.write(infolist[a] + infolist[c] + '\n')
                # 在两个人信息元素中加入特殊字符组合起来，大于等于8位就输出到字典
        for d in range(0, len(infolist)):
            for e in range(0, len(specialList)):
                if len(infolist[a] + specialList[e] + infolist[d]) >= 8:
                    # 特殊字符加在尾部
                    dictionaryFile.write(infolist[a] + infolist[d] + specialList[e] + '\n')
                    # 特殊字符加在中部
                    dictionaryFile.write(infolist[a] + specialList[e] + infolist[d] + '\n')
                    # 特殊字符加在头部
                    dictionaryFile.write(specialList[e] + infolist[a] + infolist[d] + '\n')

    dictionaryFile.close()


if __name__ == '__main__':
    # 创建文件对象
    global dictionaryFile
    # 创建字典文件
    dictionaryFile = open('passwords.txt', 'w')
    # 用户信息列表
    global infolist
    infolist = []
    # 数字列表
    global numberList
    numberList = []
    # 特殊字符列表
    global specialList
    specialList = []
    # 读取个人信息dictionaryFile
    readinformationList()
    # 创建数字列表
    CreatNumberList()
    # 创建特殊字符列表
    CreatSpecialList()
    # 把常见密码写入字典文件
    AddTopPwd()
    # 字典生成主体，将个人信息+数字列表+特殊字符列表进行组合加入字典
    Combination()
    print('\n' + u"字典生成成功！" + '\n' + '\n' + u"字典文件名：passwords")
