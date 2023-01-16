<?php

namespace Driver;

// 判断是否登录，没有登录统一重定向到登录界面
use mysql_xdevapi\Exception;

class Cookie
{
    function CheckCookie()
    {
        $If = strstr(json_encode($_COOKIE), 'name');
        // 当未登录的时候则返回false
        if (!$If) {
           return "未登录";
        }
        return $_COOKIE['name'];

    }
}