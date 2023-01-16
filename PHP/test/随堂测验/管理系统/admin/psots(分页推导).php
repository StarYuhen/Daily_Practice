<?php 

  require_once('../functions.php');

  bx_get_current_user();

  // 处理分页参数
  $page = empty($_GET['page']) ? 1 : (int)$_GET['page'];
  $size = 20;

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
  order by posts.created desc
  limit {$offset}, {$size};");


  // 处理分页页码
  // ==============================================
  
  // 求出最大页码
  $total_count = (int)bx_fetch_once('select 
    count(1) as num
  from posts
  inner join categories on posts.category_id = categories.id
  inner join users on posts.user_id = users.id;')['num'];
  $total_pages = (int)ceil($total_count / $size);

  // 计算页码
  $visiables = 5;
  $region = (int)floor($visiables/2); //左右区间
  $begin = $page - $region; //开始页码
  $end = $begin + $visiables; //结束页码 + 1
  // 可能出现 $begin 和 $end 的不合理情况(越界)
  // $begin 必须 > 0
  // 确保 $begin 最小为1
  
  // 也可用三元表达式
  if ($begin < 1) {
    $begin = 1;
    // $begin 修改意味着必须修改 $end
    $end = $begin + $visiables;
  }
  // $end 必须 <= 最大页数
  if ($end > $total_pages + 1) {
    // end 超出范围
    $end = $total_pages + 1;
    // $end 修改意味着必须修改 $begin
    $begin = $end - $visiables;
    if ($begin < 1) {
      $begin = 1;
    }
  }


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

  // function get_author ($user_id) {
  //   return bx_fetch_once('select nickname from users where id= ' . $user_id .';')['nickname'];
  // }

  // function get_category ($category_id) {
  //   return bx_fetch_once('select name from categories where id= ' . $category_id .';')['name'];
  // }
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
        <a href="post-add.html" class="btn btn-primary btn-xs">写文章</a>
      </div>
      <!-- 有错误信息时展示 -->
      <!-- <div class="alert alert-danger">
        <strong>错误！</strong>发生XXX错误
      </div> -->
      <div class="page-action">
        <!-- show when multiple checked -->
        <a class="btn btn-danger btn-sm" href="javascript:;" style="display: none">批量删除</a>
        <form class="form-inline">
          <select name="" class="form-control input-sm">
            <option value="">所有分类</option>
            <option value="">未分类</option>
          </select>
          <select name="" class="form-control input-sm">
            <option value="">所有状态</option>
            <option value="">草稿</option>
            <option value="">已发布</option>
            <option value="">回收站</option>
          </select>
          <button class="btn btn-default btn-sm">筛选</button>
        </form>
        <ul class="pagination pagination-sm pull-right">
          <li><a href="#">上一页</a></li>
          <?php for ($i = $begin; $i < $end; $i++): ?>
            <li <?php echo $i === $page ? ' class="active"' : ''; ?>>
              <a href="?page=<?php echo $i; ?>"><?php echo $i; ?></a>
            </li>
          <?php endfor ?>
          <li><a href="#">下一页</a></li>
        </ul>
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
              <td class="text-center"><input type="checkbox"></td>
              <td><?php echo $item['title']; ?></td>
              <td><?php echo $item['user_nickname']; ?></td>
              <td><?php echo $item['categoy_name']; ?></td>
              <td class="text-center"><?php echo convert_date($item['created']); ?></td>
              <td class="text-center"><?php echo convert_status($item['status']); ?></td>
              <td class="text-center">
                <a href="javascript:;" class="btn btn-default btn-xs">编辑</a>
                <a href="javascript:;" class="btn btn-danger btn-xs">删除</a>
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
  <script>NProgress.done()</script>
</body>
</html>