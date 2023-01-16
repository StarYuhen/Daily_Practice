<?php 

  require_once('../functions.php');
  bx_get_current_user();

 ?>
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8">
  <title>Comments &laquo; Admin</title>
  <link rel="stylesheet" href="/static/assets/vendors/bootstrap/css/bootstrap.css">
  <link rel="stylesheet" href="/static/assets/vendors/font-awesome/css/font-awesome.css">
  <link rel="stylesheet" href="/static/assets/vendors/nprogress/nprogress.css">
  <link rel="stylesheet" href="/static/assets/css/admin.css">
  <script src="/static/assets/vendors/nprogress/nprogress.js"></script>
  <!-- <style>
    table{
      width：100%;
      table-layout:fixed;/* 只有定义了表格的布局算法为fixed，下面td的定义才能起作用。 */
    }
    td{
      word-break:keep-all;/* 不换行 */
      /*white-space:nowrap;/* 不换行 */*/
      overflow:hidden;/* 内容超出宽度时隐藏超出部分的内容 */
      text-overflow:ellipsis;/* 当对象内文本溢出时显示省略标记(...) ；需与overflow:hidden;一起使用*/
    }
  </style> -->
</head>
<body>
  <script>NProgress.start()</script>

  <div class="main">
    <?php include 'inc/navbar.php'; ?>
    <div class="container-fluid">
      <div class="page-title">
        <h1>所有评论</h1>
      </div>
      <!-- 有错误信息时展示 -->
      <!-- <div class="alert alert-danger">
        <strong>错误！</strong>发生XXX错误
      </div> -->
      <div class="page-action">
        <!-- show when multiple checked -->
        <div class="btn-batch" style="display: none">
          <button class="btn btn-info btn-sm">批量批准</button>
          <button class="btn btn-warning btn-sm">批量拒绝</button>
          <button class="btn btn-danger btn-sm">批量删除</button>
        </div>
        <ul class="pagination pagination-sm pull-right"></ul>
      </div>
      <table class="table table-striped table-bordered table-hover">
        <thead>
          <tr>
            <th class="text-center" width="40"><input type="checkbox"></th>
            <th class="text-center" width="100">作者</th>
            <th class="text-center">评论</th>
            <th class="text-center" width="150">评论在</th>
            <th class="text-center" width="100">提交于</th>
            <th class="text-center" width="100">状态</th>
            <th class="text-center" width="150">操作</th>
          </tr>
        </thead>
        <tbody>
          <script id="comments_tmpl" type="text/x-template">
            {{each comments}}
              <tr {{if $value.status == 'held'}} class="warning"{{else if $value.status == 'rejected'}} class="danger"{{/if}}  data-id="{{$value.id}}">
                <td class="text-center"><input type="checkbox"></td>
                <td class="text-center">{{$value.author}}</td>
                <td>{{$value.content}}</td>
                <td class="text-center">《{{$value.title}}》</td>
                <td class="text-center">{{$value.created}}</td>
                <td class="text-center">{{if $value.status == 'held'}}待审{{else if $value.status == 'rejected'}}拒绝{{else if $value.status == 'approved'}}批准{{/if}}</td>
                <td class="text-center">
                  {{if $value.status == 'held'}}
                    <a href="javascript:;" class="btn btn-default btn-xs">批准</a>
                    <a href="javascript:;" class="btn btn-warning btn-xs">拒绝</a>
                  {{/if}}
                  <a href="javascript:;" class="btn btn-danger btn-xs btn-delete">删除</a>
                </td>
              </tr>
            {{/each}}
          </script>
        </tbody>
      </table>
    </div>
  </div>
  <?php $current_page = 'comments'; ?>
  <?php include 'inc/sidebar.php'; ?>
  <script src="/static/assets/vendors/jquery/jquery.js"></script>
  <script src="/static/assets/vendors/bootstrap/js/bootstrap.js"></script>
  <script src="/static/assets/vendors/template-web/template-web.js"></script>
  <script src="/static/assets/vendors/twbs-pagination/jquery.twbsPagination.js"></script>
  <script>
    $(function ($) {
      $(document)
       .ajaxStart(function () {
         NProgress.start()
       })
       .ajaxStop(function () {
         NProgress.done()
       })

      // 1.发送AJAX请求获取列表所需数据
      // 2.将数据渲染到页面上
      
      var currentPage = 1

      function loadPageDate (page) {
        $.getJSON('/admin/api/comments.php', {page: page}, function (data) {
          $('.pagination').twbsPagination({
            first:'&laquo;',
            last: '&raquo;',
            totalPages: data.total_Pages,
            visiablePages: 5,
            initiateStartPageClick: false,
            onPageClick: function (e, page) {
             // 第一次初始化就会触发一次
             loadPageDate(page)
            }
          })
          var html = template('comments_tmpl', {comments: data.comments })
          $('tbody').html(html)
          currentPage = page
        })
      }
      
      loadPageDate(currentPage)
      
      // 删除功能
      // 
      // 由于删除按钮是动态添加的，而且执行动态添加代码是在此之后执行，过早注册不上
      // $('.btn-delete').on('click', function () {
      //   console.log(11)
      // })

      // 通过事件委托
      // 方法一：重新载入当前页的数据
      $('tbody').on('click', '.btn-delete', function () {
        // 删除条数据的时机
        // 1.拿到需要删除数据的 ID
        var tr = $(this).parent().parent()
        var id = tr.data('id')
        // 2.发送一个AJAX请求 告诉服务端要删除哪一条具体数据
        $.get('/admin/api/comment-delete.php', {id: id}, function (res) {
          if (!res) return
            // 3.重新载入当前页的数据
            loadPageDate(currentPage)
        })
      })

      // 方法二：在界面上移除这个元素
      /*$('tbody').on('click', '.btn-delete', function () {
        // 删除条数据的时机
        // 1.拿到需要删除数据的 ID
        var $tr = $(this).parent().parent()
        var id = $tr.data('id')
        // 2.发送一个AJAX请求 告诉服务端要删除哪一条具体数据
        $.get('/admin/api/comment-delete.php', {id: id}, function (res) {
          if (!res) return
            // 3.根据服务端返回的删除是否成功决定是否在界面上移除这个元素
            $tr.remove()
        })
      })*/
    })
  </script>
  <script>NProgress.done()</script>
</body>
</html>