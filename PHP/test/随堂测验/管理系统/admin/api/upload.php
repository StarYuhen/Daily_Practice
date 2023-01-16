<?php

// var_dump($_FILES['avatar']);
// 接收文件
// 保存文件
// 返回这个文件的访问URL

if (empty($_FILES['avatar'])) {
	exit('必须上传文件');
}

$avatar = $_FILES['avatar'];

if ($avatar['error'] !== UPLOAD_ERR_OK) {
	exit('上传失败');
}

// 校验 类型 大小

// 移动文件到网站范围
$ext = pathinfo($avatar['name'], PATHINFO_EXTENSION);
$target = '../../static/uploads/img-' . uniqid() . '.' . $ext;

if (!move_uploaded_file($avatar['tmp_name'], $target)) {
	exit('上传失败');
}

echo substr($target, 5);