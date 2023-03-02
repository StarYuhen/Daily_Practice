# encoding:utf-8
import re
from random import choice
import asyncio
from aiohttp import ClientSession, ClientTimeout

path = 'ip.txt'  # 文件保存地址


# 随机返回请求头
async def getheaders():
    headers_list = [
        "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1",
        "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1",
        "Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11",
        "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6",
        "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6",
        "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/19.77.34.5 Safari/537.1",
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.9 Safari/536.5",
        "Mozilla/5.0 (Windows NT 6.0) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.36 Safari/536.5",
        "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
        "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
        "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
        "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
        "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
        "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
        "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
        "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
        "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
        "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24",
        "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24",
        "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36", ]
    headers = {'User-Agent': choice(headers_list)}
    return headers


# 清空文档
def clear_file():
    with open(path, 'w', encoding='utf-8') as f:
        f.truncate()


# 写入文档
async def record(text):
    with open(path, 'a', encoding='utf-8') as f:
        f.write(f'{text}\n')


# 实例化请求对象
async def create_aiohttp():
    # 实例化对象
    async with ClientSession() as session:
        task = [
            asyncio.create_task(get_page('http://www.kxdaili.com/dailiip/2/1.html', session=session)),
            get_page('https://proxy.seofangfa.com/', session=session),
            asyncio.create_task(get_page('https://www.kuaidaili.com/free/inha/1/', mod=2, session=session)),
            asyncio.create_task(get_page('https://www.kuaidaili.com/free/intr/2/', mod=2, session=session)),
            asyncio.create_task(get_page('http://www.66ip.cn/areaindex_1/1.html', session=session)),
            asyncio.create_task(get_page('http://www.66ip.cn/areaindex_5/1.html', session=session)),
            asyncio.create_task(get_page('http://www.66ip.cn/areaindex_14/1.html', session=session)),
            asyncio.create_task(
                get_page('https://www.proxy-list.download/api/v1/get?type=http', mod=3, session=session)),
        ]

        for i in range(1, 4):
            task.append(
                asyncio.create_task(get_page(f'http://www.nimadaili.com/http/{i}/', mod=4, session=session)))
            task.append(
                asyncio.create_task(get_page(f'https://www.89ip.cn/index_{i}.html', session=session)))
            task.append(asyncio.create_task(
                get_page(f'http://http.taiyangruanjian.com/free/page{i}/', mod=1, session=session)))
            task.append(
                asyncio.create_task(get_page(f'http://www.kxdaili.com/dailiip/1/{i}.html', session=session)))
            task.append(
                asyncio.create_task(get_page(f'http://www.ip3366.net/free/?stype=1&page={i}', session=session)))
            task.append(asyncio.create_task(get_page(f'http://www.66ip.cn/areaindex_1{i}/1.html', session=session)))
            task.append(asyncio.create_task(
                get_page(f'https://www.dieniao.com/FreeProxy/{i}.html', mod=5, session=session)))
        await asyncio.wait(task)


# 访问网页
async def get_page(url, session, mod=0):
    hd = await getheaders()  # 请求头
    tout = ClientTimeout(total=30)  # 超时时间
    try:
        async with await session.get(url=url, headers=hd, timeout=tout) as response:
            page_source = await response.text()  # 返回字符串形式的相应数据
            await soup_page(page_source, mod=mod)
    except Exception as e:
        print(f"['{url}']抓取失败:", e)


# 清洗页面 提取IP
# 生成代理链接格式: http://ip:port
async def soup_page(source, mod):
    if mod == 0:
        # 通用
        ips = re.findall(r'<td>[\s]*?(\d+\.\d+\.\d+\.\d+)[\s]*?</td>', source)
        posts = re.findall(r'<td>[\s]*?(\d{1,5})[\s]*?</td>', source)
        for i in range(len(ips)):
            await record(f"http://{ips[i]}:{posts[i]}")

    elif mod == 1:
        # 太阳
        ips = re.findall(r'<div.*?">(\d+\.\d+\.\d+\.\d+)</div>', source)
        posts = re.findall(r'<div.*?">(\d{1,5})</div>', source)
        for i in range(len(ips)):
            await record(f"http://{ips[i]}:{posts[i]}")

    elif mod == 2:
        # 快代理
        ips = re.findall(r'<td\s.*?="IP">(\d+\.\d+\.\d+\.\d+)</td>', source)
        posts = re.findall(r'<td\s.*?="PORT">(\d{1,5})</td>', source)
        for i in range(len(ips)):
            await record(f"http://{ips[i]}:{posts[i]}")

    elif mod == 3:
        # www.proxy-list.download/api/v1/get?type=http
        ip_list = source.split('\r\n')[:-1]
        for i in ip_list:
            await record(f"http://{i}")

    elif mod == 4:
        # 泥马代理
        ip_post = re.findall(r'<td>(.*?:\d+)</td>', source)
        for i in ip_post:
            await record(f"http://{i}")

    elif mod == 5:
        # 蝶鸟
        ip_list = re.findall(r"<span\sclass='f-address'>(.*?)</span>", source)[1:]
        port_list = re.findall(r"<span class='f-port'>(\d+)</span>", source)
        for i in range(len(ip_list)):
            await record(f'http://{ip_list[i]}:{port_list[i]}')

    elif mod == 6:
        # 站大爷
        pass


def ip_main():
    clear_file()  # 清空存放代理文件
    print('正在抓取代理ip。。。')
    asyncio.run(create_aiohttp())
    print("代理抓取完成!!!")

# ip_main()
