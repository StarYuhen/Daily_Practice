<?php 

  // 载入配置文件
  require_once('../config.php');

  // 给用户找一个箱子 （如果你之前有就用之前的，没有就给新的）
  session_start();

  // 方法一：登录前先清除登录标记
  // unset($_SESSION['current_login_user']);

  function login () {
    // 1.接收并校验
    // 2.持久化
    // 3.响应
    
    if (empty($_POST['email'])) {
     $GLOBALS['message'] = '请填写邮箱';
     return;
    }

    if (empty($_POST['password'])) {
     $GLOBALS['message'] = '请填写密码';
     return;
    }

    $email = $_POST['email'];
    $password = $_POST['password'];


    $conn = mysqli_connect(BX_DB_HOST, BX_DB_USER, BX_DB_PASS, BX_DB_NAME);

    if (!$conn) {
      exit('<center><h1>连接数据库失败</h1></center>');
    }

    $query = mysqli_query($conn, "select * from users where email = '{$email}' limit 1;");

    if (!$query) {
      $GLOBALS['message'] = '登录失败，请重试！';
      return;
    }

    $user = mysqli_fetch_assoc($query);

    if (!$user) {
      // 用户名不存在
      $GLOBALS['message'] = '邮箱与密码不匹配！';
      return;
    }
   
   // 一般密码是加密存储的
    if ($user['password'] !== md5($password)) {
      // 密码不正确
      $GLOBALS['message'] = '邮箱与密码不匹配';
      return;
    }

    // 存一个登录标识
    // $_SESSION['is_logged_in'] = true;
    // 为了后学可以直接获取当前登录用户的信息，这里直接将用户信息放到 session 中
    $_SESSION['current_login_user'] = $user;

    header('Location: /admin/');
  }

  if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    login();
  }

  // 方法二：通过退出键传来的特定值删除登录标记
  if ($_SERVER['REQUEST_METHOD'] === 'GET' && isset($_GET['action']) && $_GET['action'] === 'logout') {
    // 删除登录标识
    unset($_SESSION['current_login_user']);
  }

 ?>

<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8">
  <title>Sign in &laquo; Admin</title>
  <link rel="stylesheet" href="/static/assets/vendors/bootstrap/css/bootstrap.css">
  <link rel="stylesheet" href="/static/assets/vendors/animate/animate.css">
  <link rel="stylesheet" href="/static/assets/css/admin.css">
</head>
<body>
  <div class="login">
    <!-- 可以通过在 form 表单上添加 navalidate 取消浏览器自带的校验功能 -->
    <!-- autocomplete="off" 关闭客户端的自动完成功能 -->
    <form class="login-wrap <?php echo isset($message) ? ' shake animated' : ''; ?>" action="<?php echo $_SERVER['PHP_SELF']; ?>" method="post" novalidate autocomplete="off">
      <img class="avatar" src="/static/assets/img/default.png">
      <!-- 有错误信息时展示 -->
     <?php if (isset($message)): ?>
       <div class="alert alert-danger">
        <strong>错误！</strong><?php echo $message; ?>
      </div>
     <?php endif ?>
      <div class="form-group">
        <label for="email" class="sr-only">邮箱</label>
        <input id="email" name="email" type="email" class="form-control" placeholder="邮箱" autofocus value="<?php echo empty($_POST['email']) ? '' : $_POST['email']; ?>">
      </div>
      <div class="form-group">
        <label for="password" class="sr-only">密码</label>
        <input id="password" name="password" type="password" class="form-control" placeholder="密码">
      </div>
      <button class="btn btn-primary btn-block">登 录</button>
    </form>
  </div>
  <script src="/static/assets/vendors/jquery/jquery.js"></script>
  <script>
    $(function ($) {
      // 1.单独作用域
      // 2.确保页面加载过后之子那个
       
      // 目标：在用户输入自己的邮箱过后，页面上展示这个邮箱对应的头像
      // 实现：
      // -时机：邮箱文本款失去焦点
      // -时间：获取这个文本框中填写的邮箱的头像对应的地址，展示到上面的img标签上
      
      var emailFormat = /^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+$/
      $('#email').on('blur', function () {
        var value = $(this).val()
        // 忽略掉文本框为空或者不是一个邮箱
        if (!value || !emailFormat.test(value)) return

        // 用户输入一个合理的邮箱地址
        // 获取这个邮箱对应的头像地址，展示到上面的 img 元素上
        // 因为客户端的JS无法直接操作数据库，应该通过JS发送AJAX 请求 高数服务端的某个接口，让改接口帮助客户端获取相应的数据
        
        $.get('/admin/api/avatar.php', {email: value}, function (res) {
          // 希望 res =>这个邮箱对应的头像地址
          if (!res) return
          // 展示到上面的img元素大上
          // $('.avatar').attr('src', res)
          $('.avatar').fadeOut(function() {
            // 等到 淡出完成
            $(this).on('load', function() {
              // 等到图片完全加载成功后
              $(this).fadeIn()
            }).attr('src', res)
          })
        })
      })
    })
  </script>
</body>
</html>
