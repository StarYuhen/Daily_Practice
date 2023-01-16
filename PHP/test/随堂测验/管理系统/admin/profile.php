<?php 

  require_once('../functions.php');
  $user = bx_get_current_user();
 
  function user_update (){
    global $user;

    $id = $user['id'];
    $avatar = empty($_POST['avatar']) ? $user['avatar'] : $_POST['avatar'];
    $user['avatar'] = $avatar;
    $slug = empty($_POST['slug']) ? $user['slug'] : $_POST['slug'];
    $user['slug'] = $slug;
    $nickname = empty($_POST['nickname']) ? $user['nickname'] : $_POST['nickname'];
    $user['nickname'] = $nickname;
    $bio = empty($_POST['bio']) ? $user['bio'] : $_POST['bio'];
    $user['bio'] = $bio;

    $sql = sprintf("update users set 
      avatar = '%s', slug = '%s', nickname = '%s', bio = '%s' 
      where id = %d;", $avatar, $slug, $nickname, $bio, $id);

    $row = bx_execute($sql);
    $GLOBALS['success'] = $row > 0;
    $GLOBALS['message'] = $row <= 0 ? '更新失败!' : '更新成功!';

    $_SESSION['current_login_user'] = $user;
  }

  if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    user_update();
  }
  
 ?>
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8">
  <title>Dashboard &laquo; Admin</title>
  <link rel="stylesheet" href="/static/assets/vendors/bootstrap/css/bootstrap.css">
  <link rel="stylesheet" href="/static/assets/vendors/font-awesome/css/font-awesome.css">
  <link rel="stylesheet" href="/static/assets/vendors/nprogress/nprogress.css">
  <link rel="stylesheet" href="/static/assets/css/admin.css">
  <script src="/static/assets/vendors/nprogress/nprogress.js"></script>
</head>
<body>
  <script>NProgress.start()</script>

  <div class="main">
    <?php include 'inc/navbar.php'; ?>
    <div class="container-fluid">
      <div class="page-title">
        <h1>我的个人资料</h1>
      </div>
      <!-- 有信息时展示 -->
      <?php if (isset($message)): ?>
        <?php if ($success): ?>
          <div class="alert alert-success">
            <strong>成功！</strong><?php echo $message; ?>
          </div>
        <?php else: ?>
          <div class="alert alert-danger">
            <strong>错误！</strong><?php echo $message; ?>
          </div>
        <?php endif ?>
      <?php endif ?>
      <form class="form-horizontal" action="<?php echo $_SERVER['PHP_SELF']; ?>" method="post">
        <div class="form-group">
          <label class="col-sm-3 control-label">头像</label>
          <div class="col-sm-6">
            <label class="form-image">
              <input id="avatar" type="file">
              <img src="<?php echo $user['avatar'] ?>">
              <input type="hidden" name="avatar">
              <i class="mask fa fa-upload"></i>
            </label>
          </div>
        </div>
        <div class="form-group">
          <label for="email" class="col-sm-3 control-label">邮箱</label>
          <div class="col-sm-6">
            <input id="email" class="form-control" name="email" type="type" value="<?php echo $user['email'] ?>" placeholder="邮箱" readonly>
            <p class="help-block">登录邮箱不允许修改</p>
          </div>
        </div>
        <div class="form-group">
          <label for="slug" class="col-sm-3 control-label">别名</label>
          <div class="col-sm-6">
            <input id="slug" class="form-control" name="slug" type="type" value="<?php echo $user['slug'] ?>" placeholder="slug">
            <p class="help-block">https://zce.me/author/<strong>zce</strong></p>
          </div>
        </div>
        <div class="form-group">
          <label for="nickname" class="col-sm-3 control-label">昵称</label>
          <div class="col-sm-6">
            <input id="nickname" class="form-control" name="nickname" type="type" value="<?php echo $user['nickname'] ?>" placeholder="昵称">
            <p class="help-block">限制在 2-16 个字符</p>
          </div>
        </div>
        <div class="form-group">
          <label for="bio" class="col-sm-3 control-label">简介</label>
          <div class="col-sm-6">
            <textarea id="bio" class="form-control"  name="bio" placeholder="Bio" cols="30" rows="6"><?php echo $user['bio'] ?></textarea>
          </div>
        </div>
        <div class="form-group">
          <div class="col-sm-offset-3 col-sm-6">
            <button type="submit" class="btn btn-primary">更新</button>
            <a class="btn btn-link" href="/admin/password-reset.php?id=<?php echo $user['id'] ?>">修改密码</a>
          </div>
        </div>
      </form>
    </div>
  </div>
  <?php $current_page = 'profile'; ?>
  <?php include 'inc/sidebar.php'; ?>

  <script src="/static/assets/vendors/jquery/jquery.js"></script>
  <script src="/static/assets/vendors/bootstrap/js/bootstrap.js"></script>
  <script>
    $(function ($) {

      $('#avatar').on('change', function () {
        // 当文件选择状态变化会执行这个时间处理函数
        var $this = $(this)
        var files = $this.prop('files')
        if (!files.length) return
        // 拿到我们要上传的文件
        var file = files[0]
        // FormData 是 HTML5 中新增的一个成员，专门配合 AJAX 操作 用于在客户端与服务端之间传递二进制数据
        var data = new FormData()
        data.append('avatar', file)

        var xhr = new XMLHttpRequest()
        xhr.open('POST', '/admin/api/upload.php')
        xhr.send(data) //借助于 form data 传递文件

        xhr.onload = function () {
          // console.log(this.responseText)
          $this.siblings('img').attr('src', this.responseText)
          $this.siblings('input').val(this.responseText)
        }
      })
    })
  </script>
  <script>NProgress.done()</script>
</body>
</html>
