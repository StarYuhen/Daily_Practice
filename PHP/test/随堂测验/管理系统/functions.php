 <?php

require_once 'config.php';
/**
 * 封装公用函数
 */

session_start();

/**
 * 获取当前登录用户信息，如果没有获取到则自动跳转到登录页面
 * @return [type] [description]
 */
function bx_get_current_user () {
  if (empty($_SESSION['current_login_user'])) {
    // 没有当前登录用户信息，意味着没有登录
    header('Location: /admin/login.php');
    exit(); // 没有必要再执行之后的代码
  }
  return $_SESSION['current_login_user'];
}

/**
 * 通过一个数据库查询获取多条数据
 * =>索引关联数组
 */
function bx_fetch_all ($sql) {
	$conn = mysqli_connect(BX_DB_HOST, BX_DB_USER, BX_DB_PASS, BX_DB_NAME);
	if (!$conn) {
		exit('连接失败');
	}

	$query = mysqli_query($conn, $sql);
	if (!$query) {
		return false;
	}

	$result = array();

	while ($row = mysqli_fetch_assoc($query)) {
		$result[] = $row;
	}

	mysqli_free_result($query);
	mysqli_close($conn);
	return $result;
}

/**
 * 获取单条数据
 * 
 */
function bx_fetch_once ($sql) {
	$res = bx_fetch_all($sql);
	return isset($res[0]) ? $res[0] : null;
}


/**
 * 处理单条增删改
 */
function bx_execute ($sql) {
	$conn = mysqli_connect(BX_DB_HOST, BX_DB_USER, BX_DB_PASS, BX_DB_NAME);
	if (!$conn) {
		exit('连接失败');
	}

	$query = mysqli_query($conn, $sql);
	if (!$query) {
		return false;
	}

	$affected_rows = mysqli_affected_rows($conn);

	mysqli_close($conn);

	return $affected_rows;
}
