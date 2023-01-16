<?php

// 接收客户端的AJAX 请求 返回评论数据

require_once '../../functions.php';

// 取得客户端传递过来的分页页码
$page = empty($_GET['page']) ? 1 : intval($_GET['page']);


$length = 20;
// 根据页码计算有过多少条
$offset = ($page - 1) * $length;

$sql = sprintf('select 
	c.id, 
	c.author,
	c.content,  
	c.created, 
	c.status, 
	p.title
from comments as c
inner join posts as p on c.post_id = p.id
order by c.created desc
limit %d, %d;', $offset, $length);
// 查询所有评论数据
$comments = bx_fetch_all($sql);

// 查询所有数据的数量
$total_count = bx_fetch_once('select count(1) as num
from comments as c
inner join posts as p on c.post_id = p.id;')['num'];

$total_Pages = ceil($total_count / $length); //虽然返回的数据类型是float 但数字一定是一个整数

// 因为网络之间传输的是字符串
// 因此现将数据转换成字符创(序列化)

$json = json_encode(array(
	'total_Pages' => $total_Pages,
	'comments' => $comments
));

// 设置响应的响应体类型为JSON
header('Content-Type: application/json');

// 响应给客户端
echo $json;