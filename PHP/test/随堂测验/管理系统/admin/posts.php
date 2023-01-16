<?php 

  require_once('../functions.php');

  bx_get_current_user();

  // 处理分页参数
  $page = empty($_GET['page']) ? 1 : (int)$_GET['page'];
  $size = 20;
  
  $where = '1 = 1';
  $search = '';

  // 分类筛选
  if(isset($_GET['category']) && $_GET['category'] !== '-1'){
    $where .= ' and posts.category_id = ' . $_GET['category'];
    $search .= '&category=' .$_GET['category'];
  }

  if(isset($_GET['status']) && $_GET['status'] !== '-1'){
    $where .= " and posts.status= '{$_GET['status']}'";
    $search .= '&status=' .$_GET['status'];
  }

  if ($page < 1) {
    header('Location: /admin/posts.php?page=1' . $search);
  }

  // $where => "1 = 1 and posts.category_id = 1 and posts.status = ''"
  // "$search = &category=1&status=''"
  // 求出最大页码
  $total_count = (int)bx_fetch_once("select 
    count(1) as num
  from posts
  inner join categories on posts.category_id = categories.id
  inner join users on posts.user_id = users.id
  where {$where};")['num'];
  $total_pages = (int)ceil($total_count / $size);

  if ($page > $total_pages) {
    header('Location: /admin/posts.php?page=' . $total_pages . $search);
  }

  // 计算出要越过多少条取数据
  $offset = ($page - 1) * $size; 

  // 获取全部数据
  $posts = bx_fetch_all("select 
    posts.id,
    posts.title, 
    users.nickname as user_nickname, 
    categories.name as categoy_name, 
    posts.created, 
    posts.status
  from posts
  inner join categories on posts.category_id = categories.id
  inner join users on posts.user_id = users.id
  where {$where}
  order by posts.created desc
  limit {$offset}, {$size};");

  // 查询全部分类数据
  $categories = bx_fetch_all('select * from categories;');

  // 查询全部状态数据
  $post_status = bx_fetch_all('select status from posts group by status;');

  /*处理分页页码==============================================*/
 
  // 计算页码
  $visiables = 5;
  $region = (int)floor($visiables / 2);
  $begin = $page - $region; //开始页码
  $end = $begin + $visiables; //结束页码 + 1
  
  // 重点考虑合理性的问题(类似越界)
  // begin >0 end <=total_pages
  $begin = $begin < 1 ? 1 : $begin; // $begin > 1
  $end = $begin + $visiables; // $end 和 $begin 两者同步关系
  $end = $end > $total_pages + 1 ? $total_pages + 1 : $end; // $end 不大于 total_pages (最大页码)
  $begin = $end - $visiables; // $end 和 $begin 两者同步关系
  $begin = $begin < 1 ? 1 : $begin; // $begin > 1

  // 处理格式转换
  
  /**
   * 转换状态显示
   * @param  string $status 英文状态
   * @return string         中文状态
   */
  function convert_status ($status) {
    $dict = array(
      'published' => '已发布',
      'drafted' => '草稿',
      'trashed' => '回收站'
    );
    return isset($dict[$status]) ? $dict[$status] : '未知';
  }
 
  function convert_date ($created) {
    // => '2018-10-11 08:10:00'
    $timestamp = strtotime($created);
    return date('Y年m月d日<b\r>H:i:s', $timestamp);
  }
 ?>
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8">
  <title>Posts &laquo; Admin</title>
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
        <h1>所有文章</h1>
        <a href="/admin/post-add.php" class="btn btn-primary btn-xs">写文章</a>
      </div>
      <!-- 有错误信息时展示 -->
      <!-- <div class="alert alert-danger">
        <strong>错误！</strong>发生XXX错误
      </div> -->
      <div class="page-action">
        <!-- show when multiple checked -->
        <a class="btn btn-danger btn-sm" href="javascript:;" style="display: none">批量删除</a>
        <form class="form-inline" action="<?php echo $_SERVER['PHP_SELF']; ?>">
          <select name="category" class="form-control input-sm">
            <option value="-1">所有分类</option>
            <?php foreach ($categories as $item): ?>
              <option value="<?php echo $item['id']; ?>" <?php echo isset($_GET['category']) && $_GET['category'] == $item['id'] ? ' selected' : '' ?>><?php echo $item['name']; ?></option>
            <?php endforeach ?>
          </select>
          <select name="status" class="form-control input-sm">
            <option value="-1">所有状态</option>
            <?php foreach ($post_status as $item): ?>
              <option value="<?php echo $item['status']; ?>"<?php echo isset($_GET['status']) && $_GET['status'] == $item['status'] ? ' selected' : '' ?>><?php echo convert_status($item['status']) ?></option>
            <?php endforeach ?>
          </select>
          <button class="btn btn-default btn-sm">筛选</button>
        </form>
        <ul class="pagination pagination-sm pull-right">
          <?php if ($page > 1): ?>
            <li><a href="?page=<?php echo $page - 1 . $search; ?>"><<</a></li>
          <?php endif ?>
          <?php if ($begin > 1): ?>
            <li class="disabled">
              <span>···</span>
            </li>
          <?php endif ?>
          <?php for ($i = $begin; $i < $end; $i++): ?>
            <li <?php echo $i === $page ? ' class="active"' : ''; ?>>
              <a href="?page=<?php echo $i . $search; ?>"><?php echo $i; ?></a>
            </li>
          <?php endfor ?>
          <?php if ($end <= $total_pages): ?>
            <li >
              <span>···</span>
            </li>
          <?php endif ?>
          <?php if ($page < $total_pages): ?>
             <li><a href="?page=<?php echo $page + 1 . $search; ?>">>></a></li>
          <?php endif ?> 
        </ul>
      </div>
      <div class="page-action">
        <!-- show when multiple checked -->
        <a id="btn_delete" class="btn btn-danger btn-sm" href="/admin/post-delete.php" style="display: none">批量删除</a>
      </div>
      <table class="table table-striped table-bordered table-hover">
        <thead>
          <tr>
            <th class="text-center" width="40"><input type="checkbox"></th>
            <th>标题</th>
            <th>作者</th>
            <th>分类</th>
            <th class="text-center">发表时间</th>
            <th class="text-center">状态</th>
            <th class="text-center" width="100">操作</th>
          </tr>
        </thead>
        <tbody>
          <?php foreach ($posts as $item): ?>
            <tr>
              <td class="text-center"><input type="checkbox" data-id="<?php echo $item['id']; ?>"></td>
              <td><?php echo $item['title']; ?></td>
              <td><?php echo $item['user_nickname']; ?></td>
              <td><?php echo $item['categoy_name']; ?></td>
              <td class="text-center"><?php echo convert_date($item['created']); ?></td>
              <td class="text-center"><?php echo convert_status($item['status']); ?></td>
              <td class="text-center">
                <a href="javascript:;" class="btn btn-default btn-xs">编辑</a>
                <a href="/admin/post-delete.php?id=<?php echo $item['id']; ?>" class="btn btn-danger btn-xs">删除</a>
              </td>
            </tr>
          <?php endforeach ?>
        </tbody>
      </table>
    </div>
  </div>
  <?php $current_page = 'posts'; ?>
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