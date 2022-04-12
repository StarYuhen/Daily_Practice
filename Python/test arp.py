from scapy.all import *
import sys, getopt, socket

from scapy.layers.l2 import Ether, ARP


def get_local_net():
    # 获取网段。如：192.168.50
    try:
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        s.connect(('8.8.8.8', 80))
        # 获取本机ip。如：192.168.50.110
        ip = s.getsockname()[0]
        ippre_list = ip.split(r".")
        ippre_list.pop()
        # 获取网段字串。如：192.168.50
        ipnet = '.'.join(ippre_list)
    except Exception:
        ipnet = False
    finally:
        s.close()

    return ipnet


def get_vlan_ip_and_mac(locnet, start_num=1, end_num=255):
    # 通过arp协议扫描，发现本网段存活ip和mac
    result = []
    localnet = locnet
    scansum = int(end_num) - int(start_num) + 1

    print("%s.%s - %s.%s 共计 %s 个被扫描ip" % (localnet, start_num, localnet, end_num, scansum))
    print()
    counter = 1

    # 如果无法识别本网段，则退出扫描
    if not localnet:
        print("扫描终止：无法识别本网段。")
        return result

    for ipFix in range(start_num, end_num + 1):
        # 构造本网段的ip。如：192.168.50.20
        ip = localnet + "." + str(ipFix)

        # 组合协议包
        # 通过 '/' 可叠加多个协议层(左底层到右上层)，如Ether()/IP()/UDP()/DNS()
        arpPkt = Ether(dst="ff:ff:ff:ff:ff:ff") / ARP(pdst=ip)
        # 发送arp请求，并获取响应结果。设置1s超时。
        res = srp1(arpPkt, timeout=1, verbose=0)

        # 如果ip存活
        if res:
            print("%3d --> %s  %s" % (counter, ip, res.hwsrc))
            result.append({"localIP": res.psrc, "mac": res.hwsrc})
            counter += 1
        # 如果ip不存活
        else:
            print("%3d --> %s" % (counter, ip))
            counter += 1

    return result


if __name__ == '__main__':
    # windows 需要下载WinPcap否则报错 https://www.winpcap.org/install/default.htm
    locnet = get_local_net()
    print("一、开始扫描本网段（%s.xx）活动的ip" % locnet)
    # 扫描ip起始和终止范围
    start_num = 126
    end_num = 135
    # 开始扫描
    result = get_vlan_ip_and_mac(locnet, start_num, end_num)

    print()
    print("二、Mac表汇总清单（活动ip共计 %s个）：" % len(result))
    for dic in result:
        print(dic)
