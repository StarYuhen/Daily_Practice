<?php 

  require_once('../functions.php');
  bx_get_current_user();

  function add_user () {
    if (empty($_POST['email']) || empty($_POST['slug']) || empty($_POST['nickname']) || empty($_POST['password'])) {
      $GLOBALS['message'] = '请完整填写表单';
      $GLOBALS['success'] = false;
      return;
    }

    $email = $_POST['email'];
    $slug = $_POST['slug'];
    $nickname = $_POST['nickname'];
    $password = md5($_POST['password']);

    $row = bx_execute("insert into users values (null, '{$slug}', '{$email}', '{$password}', '{$nickname}', null, null, 'inactive');");

    $GLOBALS['success'] = $row > 0;
    $GLOBALS['message'] = $row <= 0 ? '添加失败!' : '添加成功!';
  }

  function edit_user () {
    global $current_edit_user;

    $id = $current_edit_user['id'];
    $email = empty($_POST['email']) ? $current_edit_user['email'] : $_POST['email'];
    $current_edit_user['email'] = $email; 
    $slug = empty($_POST['slug']) ? $current_edit_user['slug'] : $_POST['slug'];
    $current_edit_user['slug'] = $slug; 
    $nickname = empty($_POST['nickname']) ? $current_edit_user['nickname'] : $_POST['nickname'];
    $current_edit_user['nickname'] = $nickname; 
    $password = empty($_POST['password']) ? $current_edit_user['password'] : md5($_POST['password']);
    $current_edit_user['password'] = $password; 

    $row = bx_execute("update users set slug = '{$slug}', email = '{$email}', password = '{$password}', nickname = '{$nickname}' where id = '{$id}';");

    $GLOBALS['success'] = $row > 0;
    $GLOBALS['message'] = $row <= 0 ? '更新失败!' : '更新成功!';
  }

  // 判断是否为编辑状态
  if (empty($_GET['id'])) {
    if ($_SERVER['REQUEST_METHOD'] === 'POST') {
      // 添加
      add_user();
    }
  } else {
    // 编辑
    $current_edit_user = bx_fetch_once('select * from users where id = ' . $_GET['id'] . ';');
    if ($_SERVER['REQUEST_METHOD'] === 'POST') {
      edit_user();
    }
  }

  // 查询全部分类数据
  $users = bx_fetch_all('select * from users;');
 ?>
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8">
  <title>Users &laquo; Admin</title>
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
        <h1>用户</h1>
      </div>
      <?php if (isset($success)): ?>
        <?php if ($success): ?>
          <div class="alert alert-success">
            <strong>成功！</strong><?php echo $message; ?>
          </div>
        <?php else: ?>
          <!-- 有错误信息时展示 -->
          <div class="alert alert-danger">
            <strong>错误！</strong><?php echo $message; ?>
          </div>
        <?php endif ?>
      <?php endif ?>
      <div class="row">
        <div class="col-md-4">
          <?php if (isset($current_edit_user)): ?>
             <form action="<?php echo $_SERVER['PHP_SELF']; ?>?id=<?php echo $current_edit_user['id']; ?>" method="post" autocomplete="off">
              <h2>编辑《<?php echo $current_edit_user['nickname']; ?>》</h2>
              <div class="form-group">
                <label for="email">邮箱</label>
                <input id="email" class="form-control" name="email" type="email" placeholder="邮箱" value="<?php echo $current_edit_user['email']; ?>">
              </div>
              <div class="form-group">
                <label for="slug">别名</label>
                <input id="slug" class="form-control" name="slug" type="text" placeholder="slug" value="<?php echo $current_edit_user['slug']; ?>">
                <p class="help-block">https://zce.me/author/<strong>slug</strong></p>
              </div>
              <div class="form-group">
                <label for="nickname">昵称</label>
                <input id="nickname" class="form-control" name="nickname" type="text" placeholder="昵称" value="<?php echo $current_edit_user['nickname']; ?>">
              </div>
              <div class="form-group">
                <label for="password">密码</label>
                <input id="password" class="form-control" name="password" type="password" placeholder="密码" value="<?php echo $current_edit_user['password']; ?>">
              </div>
              <div class="form-group">
                <button class="btn btn-primary" type="submit">编辑</button>
                <a href="/admin/users.php" class="btn btn-outline-secondary">取消</a>
              </div>
            </form>
          <?php else: ?>
            <form action="<?php echo $_SERVER['PHP_SELF']; ?>" method="post" autocomplete="off">
              <h2>添加新用户</h2>
              <div class="form-group">
                <label for="email">邮箱</label>
                <input id="email" class="form-control" name="email" type="email" placeholder="邮箱">
              </div>
              <div class="form-group">
                <label for="slug">别名</label>
                <input id="slug" class="form-control" name="slug" type="text" placeholder="slug">
                <p class="help-block">https://zce.me/author/<strong>slug</strong></p>
              </div>
              <div class="form-group">
                <label for="nickname">昵称</label>
                <input id="nickname" class="form-control" name="nickname" type="text" placeholder="昵称">
              </div>
              <div class="form-group">
                <label for="password">密码</label>
                <input id="password" class="form-control" name="password" type="password" placeholder="密码">
              </div>
              <div class="form-group">
                <button class="btn btn-primary" type="submit">添加</button>
              </div>
            </form>
          <?php endif ?>
        </div>
        <div class="col-md-8">
          <div class="page-action">
            <!-- show when multiple checked -->
            <a id="btn_delete" class="btn btn-danger btn-sm" href="/admin/user-delete.php" style="display: none">批量删除</a>
          </div>
          <table class="table table-striped table-bordered table-hover">
            <thead>
               <tr>
                <th class="text-center" width="40"><input type="checkbox"></th>
                <th class="text-center" width="80">头像</th>
                <th>邮箱</th>
                <th>别名</th>
                <th>昵称</th>
                <th>状态</th>
                <th class="text-center" width="100">操作</th>
              </tr>
            </thead>
            <tbody>
              <?php foreach ($users as $item): ?>
                <tr>
                  <td class="text-center"><input type="checkbox" data-id="<?php echo $item['id']; ?>"></td>
                  <td class="text-center"><img class="avatar" src="<?php echo $item['avatar']; ?>"></td>
                  <td><?php echo $item['email']; ?></td>
                  <td><?php echo $item['slug']; ?></td>
                  <td><?php echo $item['nickname']; ?></td>
                  <td><?php echo $item['status'] === 'activated' ? '已激活' : '待激活'; ?></td>
                  <td class="text-center">
                    <a href="/admin/users.php?id=<?php echo $item['id'] ?>" class="btn btn-default btn-xs">编辑</a>
                    <a href="/admin/user-delete.php?id=<?php echo $item['id'] ?>" class="btn btn-danger btn-xs">删除</a>
                  </td>
                </tr>
              <?php endforeach ?>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
  <?php $current_page = 'users'; ?>
  <?php include 'inc/sidebar.php'; ?>

  <script src="/static/assets/vendors/jquery/jquery.js"></script>
  <script src="/static/assets/vendors/bootstrap/js/bootstrap.js"></script>
  <script>
    // 1.不要重复使用无意义的选择操作，应该采用变量去本地化
    $(function ($) {
      // 在表格中任意一个 checkbox 选中状态变化时
      var tbodyCheckBoxs = $('tbody input')
      var btnDelete = $('#btn_delete')

      // 定义一个数组记录被选中
      var allCheckeds = []
      tbodyCheckBoxs.on('change', function () {
        var id = $(this).data('id')

        // 根据有没有选中当前这个 checkbox 决定是添加还是删除
        if ($(this).prop('checked')) {
          allCheckeds.includes(id) || allCheckeds.push(id)
        } else {
          allCheckeds.splice(allCheckeds.indexOf(id), 1)
        }

        // 根据 allCheckeds 是否为空决定 是否像是批量删除按钮
        allCheckeds.length ? btnDelete.fadeIn() : btnDelete.fadeOut()
        btnDelete.prop('search', 'id=' + allCheckeds)
      })

      // 全选和全不选
      $('thead input').on('change', function () {
        // 1.获取当前选中状态
        var checked = $(this).prop('checked')
        // 2.设置给标体中的每一个
        tbodyCheckBoxs.prop('checked', checked).change()
      })
    })
  </script>
  <script>NProgress.done()</script>
</body>
</html>
