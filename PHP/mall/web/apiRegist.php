<?php
include_once("../Drive/mysql.php");

use Driver\MysqlDriver;

$no = $_POST["user"];
$name = $_POST["name"];
$pwd = $_POST["pwd"];
$New = new MysqlDriver();
$query = "insert into users(uid,uname,password) values('$no','$name','$pwd')";
$result = mysqli_query($New->New(), $query);  //执行插入命令
if ($result) {
    echo $name;
} else {
    echo "0";
}




