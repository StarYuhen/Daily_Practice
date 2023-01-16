<?php 

/**
 * 根据用户邮箱获取用户头像
 * email => image
 */

require_once('../../config.php');

// 1.接收传递过来 的邮箱
if (empty($_GET['email'])) {
	exit('缺少必要参数');
}

$email = $_GET['email'];

// 2.c查询对应的头像地址
$conn = mysqli_connect(BX_DB_HOST, BX_DB_USER,BX_DB_PASS, BX_DB_NAME);
if (!$conn) {
	exit('链接数据库失败');
}

$res = mysqli_query($conn, "select avatar from users where email = '{$email}' limit 1;");
if (!$res) {
	exit('查询失败');
}

$row = mysqli_fetch_assoc($res);

// 3.echo

echo $row['avatar'];