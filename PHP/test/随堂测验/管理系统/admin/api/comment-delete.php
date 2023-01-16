<?php 

// 根据客户端传过来的ID删除对应数据

require_once '../../functions.php';

if (empty($_GET['id'])) {
	exit('缺少必要参数');
}

$id = $_GET['id'];
// $id = (int)$_GET['id'];
// => '1 or 1 = 1'
// sql 注入

$rows = bx_execute('delete from comments where id in (' . $id .');');

// if ($rows > 0) {}
// header('Location:'. $_SERVER['HTTP_REFERER']);

header('Content-Type: application/json');

echo json_encode(($rows > 0));