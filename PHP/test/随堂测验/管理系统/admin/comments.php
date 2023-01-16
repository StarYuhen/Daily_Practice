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
          if (page > data.total_Pages) {
            loadPageDate(data.total_Pages)
            return false
          }
          $('.pagination').twbsPagination('destroy')
          $('.pagination').twbsPagination({
            first:'&laquo;',
            last: '&raquo;',
            startPage: page,
            totalPages: data.total_Pages,
            visiablePages: 5,
            initiateStartPageClick: false,
            onPageClick: function (e, page) {
             loadPageDate(page)
            }
          })
          var html = template('comments_tmpl', {comments: data.comments })
          $('tbody').html(html)
          currentPage = page
        })
      }
      
      loadPageDate(currentPage)
      

      // 通过事件委托
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
    })
  </script>
  <script>NProgress.done()</script>
</body>
</html>