import ssl
import MySQLdb
import parse
from nonebot.plugin import on_command
from nonebot.adapters.onebot.v11 import Bot, MessageSegment, Event
from nonebot.log import logger
import requests

__zx_plugin_name__ = "博客动态"
__plugin_usage__ = """
usage：
    博客动态
    指令：
        输入：动态/博客动态 获取博客动态评论
""".strip()
__plugin_des__ = "博客动态"
__plugin_cmd__ = ["动态", "博客动态"]
__plugin_version__ = 1.01
__plugin_author__ = "StarYuhen"

__plugin_settings__ = {
    "level": 5,
    "default_status": True,
    "limit_superuser": False,
    "cmd": ["博客动态", "动态"],
}

StarYuhen = on_command('动态', aliases={'博客动态'}, priority=5)


def comment():
    db = MySQLdb.connect("127.0.0.1", "root", "19b053343bb98618", "www_yuhen_club", charset='utf8',
                         unix_socket='/tmp/mysql.sock')
    # 使用cursor()方法获取操作游标
    cursor = db.cursor()
    # SQL 查询语句
    sql = "SELECT text FROM typecho_comments WHERE mail='%s' order by rand() limit 1" % "3446623843@qq.com"
    try:
        # 执行SQL语句
        cursor.execute(sql)
        # 获取所有记录列表
        results = cursor.fetchall()
        for row in results:
            # 打印结果
            logger.info(f"查询博客动态成功{row[0]}")
            return True, row[0]
    except:
        return False, "读取博客动态失败", ""


# 提取消息和内容
def imgURL(content):
    # 截取掉消息和图片地址，消息只展示一个图片内容
    try:
        src = content.split("<")
        pattern = parse.parse('img src={src}>', src[1])
    except:
        return False, src
    return True, src[0], eval(pattern["src"])


@StarYuhen.handle()
async def _(bot: Bot, ev: Event):
    Comment = comment()
    # 为真则说明读取成功
    if Comment[0]:
        commentFunction = imgURL(Comment[1])
        # 为真就说明有图片
        if commentFunction[0]:
            logger.info(f"动态有图片,消息：{commentFunction[1]},图片地址：{commentFunction[2]}")
            img = requests.get(commentFunction[2])
            await StarYuhen.send(MessageSegment.text(commentFunction[1]) + MessageSegment.image(img.content),
                                 at_sender=True)
        # 没有图片只有消息体
        else:
            logger.info(f"动态没有图片,消息：{commentFunction[1]}")
            await StarYuhen.send(commentFunction[1], at_sender=True)
    else:
        logger.info(f"动态没有图片,消息：{Comment[1]}")
        await StarYuhen.send(Comment[1], at_sender=True)

# def main():
#     Comment = comment()
#     # 为真则说明读取成功
#     if Comment[0]:
#         commentFunction = imgURL(Comment[1])
#         # 为真就说明有图片
#         if commentFunction[0]:
#             print(commentFunction[1], commentFunction[2])
#         else:
#             print(commentFunction[1])
#     else:
#         print(Comment[1])
#
#
# main()
