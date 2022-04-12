#!/usr/bin/env python3
# -*- coding:utf-8 -*-
# Author:Ameng, jlx-love.com

import ipaddress
import optparse

from scapy.all import *
from scapy.layers.l2 import Ether, ARP


def scapy_arp_scan(network, ifname):
    net = ipaddress.ip_network(network)
    ip_addr = []
    for ip in net:
        ip = str(ip)
        ip_addr.append(ip)
    time.sleep(1)
    Packet = Ether(dst='FF:FF:FF:FF:FF:FF') / ARP(op=1, hwdst='00:00:00:00:00:00', pdst=ip_addr)
    arp = srp(Packet, iface=ifname, timeout=1, verbose=False)
    arp_list = arp[0].res
    IP_MAC_LIST = []
    for n in range(len(arp_list)):
        IP = arp_list[n][1][1].fields['psrc']
        MAC = arp_list[n][1][1].fields['hwsrc']
        IP_MAC = [IP, MAC]
        IP_MAC_LIST.append(IP_MAC)
    return IP_MAC_LIST


if __name__ == '__main__':
    # 启动命令 python3 scapy_arp_scan.py --network 192.168.0.1/24 --ifname eth0
    t1 = time.time()
    parser = optparse.OptionParser('用法：\n python3 scapy_arp_scan.py --network 扫描网段 --ifname 网卡名称')
    parser.add_option('--network', dest='network', type='string', help='扫描网段')
    parser.add_option('--ifname', dest='ifname', type='string', help='网卡名称')
    (options, args) = parser.parse_args()
    network = options.network
    ifname = options.ifname
    if network == None or ifname == None:
        print(parser.usage)
    else:
        active_ip_mac = scapy_arp_scan(network, ifname)
        print('存活的IP地址及对应MAC：')
        for ip, mac in active_ip_mac:
            print(ip, mac)
    t2 = time.time()
    print('所用时间为：{}'.format(int(t2 - t1)))
