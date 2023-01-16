<?php 

// 根据客户端传过来的ID删除对应数据

require_once '../functions.php';

if (empty($_GET['id'])) {
	exit('缺少必要参数');
}

$id = $_GET['id'];
// $id = (int)$_GET['id'];
// => '1 or 1 = 1'
// sql 注入

$rows = bx_execute('delete from users where id in (' . $id .');');

// if ($rows > 0) {}
header('Location: /admin/users.php');