<?php
namespace Driver;
class MysqlDriver{
    function New(): bool|\mysqli
    {
        $conn=mysqli_connect("localhost","root","abc123456");
        mysqli_select_db($conn,"market");//选择market数据库
        return $conn;
    }
    function Close ($conn){
        mysqli_close($conn); //关闭连接
    }
}