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

$rows = bx_execute('delete from posts where id in (' . $id .');');

// if ($rows > 0) {}
// http 中的 referer 用来标识当前页面请求的来源
header('Location:' . $_SERVER['HTTP_REFERER']);