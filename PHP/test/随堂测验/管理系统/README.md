# 			baixiu项目总结

### 准备阶段

- 多看看猪跑

  项目开始不会做并不一定是因为技术不到家，而常常是因为没见过。所以在做项目前多看看别人是怎么做的，别着急动手。因为自己做没有确定方向，最后可能会出现以下两种情况：

  - 没明确方向，进度慢.
  - 做出的产品不符合标准.

- 选定技术类型

  选定该项目用什么技术实现，如：代码语言、数据库等，选定之后中途一般不能更改。中途更改会导致许多东西要推到重做，成本太高。

  所以选定实现技术要求：

  - 一定要自己会的技术，项目是由时限的，不要用一个自己根本不会或者不熟悉的技术去做.
  - 尽量不要用市面上最新的技术，因为新技术虽然会有一些独到之处，但是也会有一些没发现的BUG。新技术还没经历过市场的考验，采用新技术如果出现一些不明所以错误，自己根本无从下手解决。而老的技术，常用的技术，经历过市场的考验，许多BUG都已经修复，即使出现错误都基本能找到解决的方法.

- 设计界面原型

  项目在开始写代码前一定要把界面原型设计好。界面原型能够清晰的反应项目的路线，该怎么做。无论是个人还是项目组，都有清晰的思路，项目组中各个模块能更好的接口。

  常用的界面原型设计工具：

  - axure

  - 墨刀

    该软件是网页在线版，主要针对移动端界面的设计。网页链接：https://modao.cc/

- 数据库设计

  数据库设计要在项目进入代码阶段前完成，尽量不要在中途添加和删减数据。因为数据库设计不但影响代码的实现，还可能与多个模块都有衔接，中途更改或导致代码要推到重来。

- 网站目录设计

  根据项目要求，合理分类文件所属目录，利于管理。

- 饭一口口吃(先设计静态界面)

  先把主要的界面设计出来，利用一些假数据做出实现效果，网页之间各个链接也要做好。完成之后分别测试效果，然后根据需求修改。

- 修改文件后缀

  一切准备就绪后可对文件后缀进行修改，例如：.html修改为.php.

  - 网站目录下的文件后缀修改.

    > 可通过命令行进行批量修改.
    >
    > 命令行转到所要修改的文件目录下.
    >
    > 某个路径> ren *.html *.php   -------------------------------->ren 为替换所有简写， *代表选中所有具有相同后缀名的文件， *.html为修改前的文件后缀， *.php为修改后的文件后缀。

  - 页面跳转链接或者导入文件批量修改.

    > 在Sublime Text3中批量修改操作：
    >
    > - Ctrl + F --------------->选中所要修改的内容
    > - 右键点击该文件----->Fine Advanced --------->In Parent Folders
    > - 在Replace框中输入替换后的内容

- 把公用的代码放在一个文件下

  - 把多个页面共用的代码块放在同一个文件里，节省内存，便于管理和维护。

  - 共用代码在各个页面有略微的差别，可在引用该文件前根据当前执行的文件设置特定的值，然后通过特定值判断该如何执行共用代码，也可以在共用代码开始前设置获取当前执行文件名或路径，根据路名或路径不同决定代码的执行。

    > 获取文件路名：
    >
    > $_SERVER['PHP_SELF']
    >
    > 获取文件路径：
    >
    >  __ FILE __  ------------>D:\XXXX\XX.PHP

  - 加载文件到调用模块相应的位置

    > include 'xxx.php';

- 

------



### 功能实现

- 画流程图

  流程图能够清晰显示功能逻辑。流程图反应了每一步该干什么事？出现什么情况？该怎么解决？解决了程序员走一步看一步的情况。有明确的逻辑思路不至于出错或代码东凑西凑。

- 表单提交

  验证表单提交数据是否为空

  - GET表单提交

    ```
    function doGet () {
        //GET提交验证
        if (empty($_GET['name'])) {//判断是否提交了该数据
             $GLOBALS['message'] = '提示信息';
             return;
        }
    }
    ```

  - POST表单提交

    ```
    function doPost () {
        //POST提交验证
        if (empty($_POST['name'])) {//判断是否提交了该数据
             $GLOBALS['message'] = '提示信息';
             return;
        }
    }
    ```

    

  - upload文件上传验证

    ```
    function doFile () {
        //文件提交验证
        if(empty($_FILE['name])){ //判断是否提交了该数据
             $GLOBALS['message'] = '提示信息';
             return;
        }
        $file = $_FILE['name];
        if ($file['error'] !== UPLOAD_ERR_OK) {
    	exit('上传失败');
    	}
    	
    	//校验文件的类型
        if (strpos($file['type'],'type/') !== 0) {
          $GLOBALS['message'] = '文件格式错误';
          return;
        }
        
        //校验文件的类型大小
        if ($file['size'] > 2*1024*1024) {
          $GLOBALS['_message'] = '文件过大';
          return;
        }
        
        //获取文件扩展名
    	$ext = pathinfo($avatar['name'], PATHINFO_EXTENSION);
    	//目的文件路径
    	$target = '../../static/uploads/img-' . uniqid() . '.' . $ext;
    	// 移动文件到网站范围
    	if (!move_uploaded_file($avatar['tmp_name'], $target)) {
    		exit('上传失败');
    	}
    	
    	//提取相对路径
    	substr($target, 5);
    }
    ```

  - 多文件上传

  - ```
    <!--enctype="multipart/form-data"设置文件上传  autocomplete="off"关闭文本框默认提示-->
    <form action="<?php echo $_SERVER['PHP_SELF']; ?>" method="post" enctype="multipart/form-data" autocomplete="off">
          <div class="form-group">
            <label for="title"></label>
            <input type="text" class="form-control" id="title" name="title">
          </div>
          <div class="form-group">
            <label for="artist"></label>
            <input type="text" class="form-control" id="artist" name="artist">
          </div>
          <div class="form-group">
            <label for="images"></label>
            <!-- multiple 可以让一个文件域多选 -->
            <input type="file" class="form-control" id="images" name="images[]" accept="image/*" multiple>
          </div>
          <div class="form-group">
            <label for="source">音乐</label>
            <!-- accept可以限制文件域能够选择的文件种类， 阻止是 MIME Type -->
            <input type="file" class="form-control" id="source" name="source" accept="audio/*">
          </div>
          <button class="btn btn-primary btn-block">保存</button>
        </form>
    ```

    表单提交

    ```
    if ($_SERVER['REQUEST_METHOD'] === 'POST') {
        fun();
      }
    ```

- 数据库MySQL连接

  - 一般会把配置信息放在一个单独的文件中，用常量define定义.

    ```
    <?php 
    header("Content-type:text/html;charset=utf-8");
    /**
     * 数据库主机
     */
    define('DB_HOST', 'localhost');
    /**
     * 数据库用户名
     */
    define('DB_USER', 'root');
    /**
     * 数据库密码
     */
    define('DB_PASS', '123456');
    /**
     * 数据库名字
     */
    define('DB_NAME', 'baixiu');
    
    /**
     * F:\xx\xx\xx\xx.php
     * F:\xx\xx\xx\
     */
    define('DIR_ROOT', dirname(__FILE__));
    ```

  - 封装数据库函数

    ```
    <?php 
    //引用常量
    require_once 'config.php';
    
    //连接数据库
    function data_sql() {
        $conn = mysqli_connect(DB_HOST, DB_USER, DB_PASS, DB_NAME);
    	if (!$conn) {
    		exit('连接失败');
    	}
    	
    	$query = mysqli_query($conn, $sql);
    	if (!$query) {
    		return false;
    	}
    }
    
    //查询全部数据
    function data_exectueQueryAll ($sql) {
        data_sql();
        
    	$result = array();
    
    	while ($row = mysqli_fetch_assoc($query)) {
    		$result[] = $row;
    	}
    
    	mysqli_free_result($query);
    	mysqli_close($conn);
    	return $result;
    }
    
    //单条数据查询
    function data_exectueQueryOne ($sql) {
        data_sql();
        $res = data_exectueQueryAll($sql);
    	return isset($res[0]) ? $res[0] : null;
    }
    
    //数据处理(增、删、改)
    function data_executeData ($sql) {
        data_sql();
        
        $affected_rows = mysqli_affected_rows($conn);
    
    	mysqli_close($conn);
    
    	return $affected_rows;
    }
    ```

- 登录页处理

  - 数据匹配和登录设置

    > 1. 验证数据是否提交.
    >
    > 2. 验证提交的数据是否与数据库匹配.
    >
    >    [^注]: 查找必须要有限制条件，在sql语句中加入 limit 1 限制数据库查询到就立即停止，减少性能消耗。密码是加密存储，要加密才能匹配。
    >
    > 3. 登录失败要配备提示信息，提示错误.
    >
    > 4. 登录成功，设置SESSION，跳转到首页.
    >
    >    ```
    >    $_SESSION['current_login_user'] = $user;
    >    
    >    header('Location: /admin/');
    >    ```
    >
    > 5. 登出要清除SESSION
    >
    >    ```
    >    // 方法一：登录前先清除登录标记
    >     unset($_SESSION['current_login_user']);
    >     
    >     // 方法二：通过退出键传来的特定值删除登录标记
    >      if ($_SERVER['REQUEST_METHOD'] === 'GET' && isset($_GET['action']) && 		$_GET['action'] === 'logout') {
    >        // 删除登录标识
    >       	unset($_SESSION['current_login_user']);
    >      }
    >    ```

  - AJAX异步获取用户头像

    > 1. 利用正则表达式验证输入的是否为邮箱格式.
    >
    >    ```
    >     var emailFormat = /^[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+$/
    >    ```
    >
    > 2. 当输入域失去焦点 (blur)  并且文本域不为空发送AJAX请求.
    >
    >    ```
    >    $('#email').on('blur', function () {
    >            var value = $(this).val()
    >            // 忽略掉文本框为空或者不是一个邮箱
    >            if (!value || !emailFormat.test(value)) return
    >    ```
    >
    > 3. 发送AJAX请求.
    >
    >    ```
    >    		$.get('/admin/api/avatar.php', {email: value}, function (res) {
    >              // res =>这个邮箱对应的头像地址
    >              if (!res) return
    >    ```
    >
    > 4. 将获取的结果res设置成头像img中src的值.
    >
    >    ```
    >    		$('.avatar').fadeOut(function() {
    >                // 等到 淡出完成
    >                $(this).on('load', function() {
    >                  // 等到图片完全加载成功后
    >                  $(this).fadeIn()
    >                }).attr('src', res) $('.avatar').fadeOut(function() {
    >                // 等到 淡出完成
    >                $(this).on('load', function() {
    >                  // 等到图片完全加载成功后
    >                  $(this).fadeIn()
    >                }).attr('src', res)
    >              })
    >            })
    >          })
    >    ```


- 后台管理首页

  首页一般要注意的有两点：

  - 必须首先检查是否设置了SESSION，如不设置则跳转回登录页，其他后台页面也是。

  - 首页一般是一些数据的综合展示。为了更直观反应数据的变化，占比等，一般都不会直接用纯数据或表格来展示，而是会用些图表图形形式来展示。例如下面两个库：

    - chart.js    https://www.chartjs.org/  (英文版)
    - echart.js  http://echarts.baidu.com/  (中文版)

    > 以echart.js为例，常见的图形有：
    >
    > 1. pie
    >
    >    ![pie](F:\web前端\bx-images\pie.PNG)
    >
    >    HTML代码
    >
    >    ```
    >    <div id="canvas-holder" style="width:40%">
    >    	<canvas id="chart-area"></canvas>
    >    </div>
    >    ```
    >
    >    javaScript代码
    >
    >    ```
    >    <script>
    >        var ctx = document.getElementById('chart-area').getContext('2d')
    >        var myChart = new Chart(ctx, {
    >            type: 'pie',
    >            data: {
    >                datasets: [{
    >                    data: [
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                    ],
    >                    backgroundColor: [
    >                        window.chartColors.red,
    >                        window.chartColors.orange,
    >                        window.chartColors.yellow,
    >                        window.chartColors.green,
    >                        window.chartColors.blue,
    >                    ],
    >                    label: 'Dataset 1'
    >                }],
    >                labels: [
    >                    'Red',
    >                    'Orange',
    >                    'Yellow',
    >                    'Green',
    >                    'Blue'
    >                ]
    >            },
    >            options: {
    >                responsive: true
    >            }
    >        }
    >    </script>
    >    ```
    >
    >    
    >
    > 2. Line
    >
    >    ![Line](F:\web前端\bx-images\line.png)
    >
    >    HTML代码
    >
    >    ```
    >    div style="width:75%;">
    >        <canvas id="canvas"></canvas>
    >    </div>
    >    ```
    >
    >    javaScript代码
    >
    >    ```
    >    <script>
    >    	var config = {
    >            type: 'line',
    >            data: {
    >                labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
    >                datasets: [{
    >                    label: 'My First dataset',
    >                    backgroundColor: window.chartColors.red,
    >                    borderColor: window.chartColors.red,
    >                    data: [
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor()
    >                    ],
    >                    fill: false,
    >                }, {
    >                    label: 'My Second dataset',
    >                    fill: false,
    >                    backgroundColor: window.chartColors.blue,
    >                    borderColor: window.chartColors.blue,
    >                    data: [
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor(),
    >                        randomScalingFactor()
    >                    ],
    >                }]
    >            },
    >            options: {
    >                responsive: true,
    >                title: {
    >                    display: true,
    >                    text: 'Chart.js Line Chart'
    >                },
    >                tooltips: {
    >                    mode: 'index',
    >                    intersect: false,
    >                },
    >                hover: {
    >                    mode: 'nearest',
    >                    intersect: true
    >                },
    >                scales: {
    >                    xAxes: [{
    >                        display: true,
    >                        scaleLabel: {
    >                            display: true,
    >                            labelString: 'Month'
    >                        }
    >                    }],
    >                    yAxes: [{
    >                        display: true,
    >                        scaleLabel: {
    >                            display: true,
    >                            labelString: 'Value'
    >                        }
    >                    }]
    >                }
    >            }
    >        };
    >    
    >        window.onload = function() {
    >            var ctx = document.getElementById('canvas').getContext('2d');
    >            window.myLine = new Chart(ctx, config);
    >        };
    >    </script>
    >    ```
    >
    >    
    >
    > 3. 其中还有很多图表 ，可以参考：https://www.chartjs.org/samples/latest/
    >
    > 4. 参考手册：https://www.chartjs.org/docs/latest/

- 数据展示

  - 图表形式直观反映数据的各项对比.

  - 表格形式展示，便于对数据的各项操作.

    - 同步方式表格数据展示.

      > 直接把数据库查询到的数据遍历到表格中.
      >
      > ```
      > <?php foreach ($data as $value): ?>
      > 	<tr>
      > 		<td><?php echo $value['title']; ?></td>
      > 		<td><?php echo $value['content']; ?></td>
      > 	</tr>
      > <?php endforeach ?>
      > ```
      >
      > 

    - 异步加载数据

      > 异步加载展现数据一般有两个步骤：
      >
      > 1. 通过AJAX发送请求获取数据.
      >
      >    ```
      >     $.getJSON('/admin/api/comments.php', {page: page}, function (data) {
      >     //操作数据data
      >    }
      >    ```
      >
      > 2. 通过模板引擎渲染到页面上.
      >
      >    ```
      >    //定义模板
      >    <tbody>
      >      <script id="comments_tmpl" type="text/x-template">//x代表自定义，template为所使用的模板引擎
      >        {{each comments}}
      >          <tr data-id="{{$value.id}}">
      >            <td class="text-center"><input type="checkbox"></td>
      >            <td class="text-center">{{$value.author}}</td>
      >            <td>{{$value.content}}</td>
      >            <td class="text-center">《{{$value.title}}》</td>
      >            <td class="text-center">{{$value.created}}</td>
      >            <td class="text-center"></td>
      >            <td class="text-center">
      >              <a href="javascript:;" class="btn btn-default btn-xs">批准</a>
      >              <a href="javascript:;" class="btn btn-warning btn-xs">拒绝</a>
      >              <a href="javascript:;" class="btn btn-danger btn-xs btn-delete">删除</a>
      >            </td>
      >          </tr>
      >        {{/each}}
      >      </script>
      >    </tbody>
      >    //渲染
      >    <script>
      >     $.getJSON('/admin/api/comments.php', {page: page}, function (data) {
      >     	var html = template('comments_tmpl', {comments: data.comments })
      >        $('tbody').html(html)
      >      }
      >     </script>
      >    ```
      >
      >    上述用的模板引擎art-template: https://aui.github.io/art-template/
      >
      >    参考手册 https://aui.github.io/art-template/docs/

- 分页查询

  - 同步加载分页查询.

    > 同步加载时直接查询出数据有多少页，然后通过遍历输出到页面中，传递参数方式是URL ? 传参，每次刷新页面数据要刷新整个页面.
    >
    > ```
    > <?php
    > // 计算页码
    >   $visiables = 5;
    >   $region = (int)floor($visiables / 2);
    >   $begin = $page - $region; //开始页码
    >   $end = $begin + $visiables; //结束页码 + 1
    >   
    >   // 重点考虑合理性的问题(类似越界)
    >   // begin >0 end <=total_pages
    >   $begin = $begin < 1 ? 1 : $begin; // $begin > 1
    >   $end = $begin + $visiables; // $end 和 $begin 两者同步关系
    >   $end = $end > $total_pages + 1 ? $total_pages + 1 : $end; // $end 不大于 total_pages (最大页码)
    >   $begin = $end - $visiables; // $end 和 $begin 两者同步关系
    >   $begin = $begin < 1 ? 1 : $begin; // $begin > 1
    >   ?>
    >   
    > <ul class="pagination pagination-sm pull-right">
    >   <?php if ($page > 1): ?>
    >     <li><a href="?page=<?php echo $page - 1 . $search; ?>"><<</a></li>
    >   <?php endif ?>
    >   <?php if ($begin > 1): ?>
    >     <li class="disabled">
    >       <span>···</span>
    >     </li>
    >   <?php endif ?>
    >   <?php for ($i = $begin; $i < $end; $i++): ?>
    >     <li <?php echo $i === $page ? ' class="active"' : ''; ?>>
    >       <a href="?page=<?php echo $i . $search; ?>"><?php echo $i; ?></a>
    >     </li>
    >   <?php endfor ?>
    >   <?php if ($end <= $total_pages): ?>
    >     <li >
    >       <span>···</span>
    >     </li>
    >   <?php endif ?>
    >   <?php if ($page < $total_pages): ?>
    >      <li><a href="?page=<?php echo $page + 1 . $search; ?>">>></a></li>
    >   <?php endif ?> 
    > </ul>
    > ```
    >
    > 

  - 异步加载分页查询.

    > 异步加载可通过AJAX获取数据页数，可通过利用别人已有的库渲染分页导航。异步加载只刷新数据，而不不刷新网页。下面利用twbs-pagination分页库渲染：
    >
    > ```
    >  <script src="/static/assets/vendors/twbs-pagination/jquery.twbsPagination.js"></script>
    >  
    >  <script>
    >     $(function ($) {
    >       // 1.发送AJAX请求获取列表所需数据
    >       // 2.将数据渲染到页面上
    >       
    >       var currentPage = 1
    > 
    >       function loadPageDate (page) {
    >         $.getJSON('/admin/api/comments.php', {page: page}, function (data) {
    >           if (page > data.total_Pages) {
    >             loadPageDate(data.total_Pages)
    >             return false
    >           }
    >           
    >           $('.pagination').twbsPagination('destroy')
    >           $('.pagination').twbsPagination({
    >             first:'&laquo;',
    >             last: '&raquo;',
    >             startPage: page,
    >             totalPages: data.total_Pages,
    >             visiablePages: 5,
    >             initiateStartPageClick: false,
    >             onPageClick: function (e, page) {
    >              loadPageDate(page)
    >             }
    >           })
    >           var html = template('comments_tmpl', {comments: data.comments })
    >           $('tbody').html(html)
    >           currentPage = page
    >         })
    >       }
    >       
    >      loadPageDate(currentPage)
    >   </script>
    > ```
    >
    > twbs-paginatin官网: http://josecebe.github.io/twbs-pagination/

  - 同步和异步的区别.

    > 同步加载需要刷新整个网页，消耗性能和流量大。
    >
    > 异步加载只需刷新数据，性能消耗小，但不利于SEO。

- 异步加载后台数据

  > 多个数据区可以利用数组键值对方式区分，然后通过json方式返回值。
  >
  > ```
  > $json = json_encode(array(
  > 	'data1' => $data1,
  > 	'data2' => $data2
  > ));
  > 
  > // 设置响应的响应体类型为JSON
  > header('Content-Type: application/json');
  > 
  > // 响应给客户端
  > echo $json;
  > ```
  >
  > 客户端可以通过data.data1方式获取所需数据

- 批量操作

  > 对数据进行批量操作，如批量删除.
  >
  > 通过监听某条或某部分数据前面的多选框的转态变化，决定是否显示批量操作按钮。定义一个数组存放被选中数据的id，根据数据的多选框checked的值为ture则把该数据的id加入数组，在此点击变成false则移除。
  >
  > ```
  > <div class="page-action">
  >     <!-- show when multiple checked -->
  >     <a id="btn_delete" class="btn btn-danger btn-sm" href="/admin/category-delete.php" style="display: none">批量删除</a>
  > </div>
  > <tr>
  >   <td class="text-center"><input type="checkbox" data-id="<?php echo $item['id']; ?>"></td>
  > </tr>
  > 
  > <script>
  > // 1.不要重复使用无意义的选择操作，应该采用变量去本地化
  > $(function ($) {
  >   // 在表格中任意一个 checkbox 选中状态变化时
  >   var tbodyCheckBoxs = $('tbody input')
  >   var btnDelete = $('#btn_delete')
  > 
  >   // 定义一个数组记录被选中
  >   var allCheckeds = []
  >   tbodyCheckBoxs.on('change', function () {
  >     var id = $(this).data('id')
  > 
  >     // 根据有没有选中当前这个 checkbox 决定是添加还是删除
  >     if ($(this).prop('checked')) {
  >       // allCheckeds.indexOf(id) === -1 || allCheckeds.push(id)
  >       allCheckeds.includes(id) || allCheckeds.push(id)
  >     } else {
  >       allCheckeds.splice(allCheckeds.indexOf(id), 1)
  >     }
  > 
  >     // 根据 allCheckeds 是否为空决定 是否像是批量删除按钮
  >     allCheckeds.length ? btnDelete.fadeIn() : btnDelete.fadeOut()
  >     btnDelete.prop('search', 'id=' + allCheckeds)
  >   })
  > 
  >   // 全选和全不选
  >   $('thead input').on('change', function () {
  >     // 1.获取当前选中状态
  >     var checked = $(this).prop('checked')
  >     // 2.设置给标体中的每一个
  >     tbodyCheckBoxs.prop('checked', checked).change()
  >   })
  > })
  > </script>
  > ```
  >
  > 服务端处理数据
  >
  > ```
  > <?php 
  > 
  > // 根据客户端传过来的ID删除对应数据
  > 
  > require_once '../functions.php';
  > 
  > if (empty($_GET['id'])) {
  > 	exit('缺少必要参数');
  > }
  > 
  > $id = $_GET['id'];
  > // $id = (int)$_GET['id'];
  > // => '1 or 1 = 1'
  > // sql 注入
  > 
  > $rows = bx_execute('delete from categories where id in (' . $id .');');
  > 
  > // if ($rows > 0) {}
  > header('Location: /admin/categories.php');
  > ```
  >
  > 注意：当 id= ‘1 = 1’ 表示什么参数都没传

- 异步文件上传

  > 根据文本域转态变化决定是否执行异步上传处理
  >
  > ```
  > <div class="form-group">
  >   <label class="col-sm-3 control-label">头像</label>
  >   <div class="col-sm-6">
  >     <label class="form-image">
  >       <input id="avatar" type="file">
  >       <img src="<?php echo $user['avatar'] ?>">
  >       <input type="hidden" name="avatar">
  >       <i class="mask fa fa-upload"></i>
  >     </label>
  >   </div>
  > </div>
  > ```
  >
  > ```
  > <script>
  >     $(function ($) {
  >       $('#avatar').on('change', function () {
  >         // 当文件选择状态变化会执行这个时间处理函数
  >         var $this = $(this)
  >         var files = $this.prop('files')
  >         if (!files.length) return
  >         // 拿到我们要上传的文件
  >         var file = files[0]
  >         // FormData 是 HTML5 中新增的一个成员，专门配合 AJAX 操作 用于在客户端与服务端之间传递二进制数据
  >         var data = new FormData()
  >         data.append('avatar', file)
  > 
  >         var xhr = new XMLHttpRequest()
  >         xhr.open('POST', '/admin/api/upload.php')
  >         xhr.send(data) //借助于 form data 传递文件
  > 
  >         xhr.onload = function () {
  >           // console.log(this.responseText)
  >           $this.siblings('img').attr('src', this.responseText)
  >           $this.siblings('input').val(this.responseText)
  >         }
  >       })
  >     })
  >   </script>
  > ```
  >
  > 服务端处理数据只需把客户端传来的数据上传即可
  >
  > ```
  > <?php
  > 
  > // var_dump($_FILES['avatar']);
  > // 接收文件
  > // 保存文件
  > // 返回这个文件的访问URL
  > 
  > if (empty($_FILES['avatar'])) {
  > 	exit('必须上传文件');
  > }
  > 
  > $avatar = $_FILES['avatar'];
  > 
  > if ($avatar['error'] !== UPLOAD_ERR_OK) {
  > 	exit('上传失败');
  > }
  > 
  > // 校验 类型 大小
  > 
  > // 移动文件到网站范围
  > $ext = pathinfo($avatar['name'], PATHINFO_EXTENSION);
  > $target = '../../static/uploads/img-' . uniqid() . '.' . $ext;
  > 
  > if (!move_uploaded_file($avatar['tmp_name'], $target)) {
  > 	exit('上传失败');
  > }
  > 
  > echo substr($target, 5);
  > ```
  >
  > 

### 扩展

- 图表库
  - Chart.js  https://www.chartjs.org/
  - echart.js http://echarts.baidu.com/
- 模板引擎
  - art-template http://aui.github.io/art-template/
  - jsRender https://www.jsviews.com/
- 加载库
  - NProgress.js http://ricostacruz.com/nprogress/
  - Loading  http://loading.awesomes.cn/
- 分页库
  - twbs-pagination http://josecebe.github.io/twbs-pagination/

### 总结

- 不要想着把所有点想到再写代码，有些问题会在写代码过程中自然而然的显现，解决方法也会跟着产生。
- 在合适的时机干合适的事。