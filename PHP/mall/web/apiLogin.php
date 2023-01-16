<?php
include_once("../Drive/mysql.php");

use Driver\MysqlDriver;

$uname = $_POST["uname"];
$pwd = $_POST["pwd"];
$New = new MysqlDriver();
$query = "select * from users where uname='$uname' and  password='$pwd' ";
$result = mysqli_query($New->New(), $query);  //执行查询命令
$arr = mysqli_fetch_assoc($result);
mysqli_free_result($result); //关闭结果集
if ($arr) {
    setcookie("name", $arr["uname"]);
    echo $arr["uname"];
} else {
    echo "0";
}
