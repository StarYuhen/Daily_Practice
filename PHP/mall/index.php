<html>

<head>
    <meta http-equiv="content-type" content="text/html;charset=utf-8"/>
    <link rel="stylesheet" href="layui/css/layui.css">
    <script src="layui/layui.js"></script>
    <title>商品后台管理界面</title>
</head>
<!-- <FRAMESET rows="10%,*">-->
<!--    <frame src="head.php" name="frame1"></frame>-->
<!--	<FRAMESET COLS="20%,*">-->
<!--	<frame src="left.php" name="frame2"></frame>-->
<!--	<frame src="goodslist.php" name="frame3"></frame>-->
<!--    </FRAMESET>-->
<!-- </FRAMESET>-->

<body>
<div>
    <ul class="layui-nav layui-nav-tree layui-nav-side layui-bg-cyan" lay-filter="test">
        <!-- 侧边导航: <ul class="layui-nav layui-nav-tree layui-nav-side"> -->
        <!--        <li class="layui-nav-item layui-nav-itemed">-->
        <!--            <a href="javascript:;">用户</a>-->
        <!--            <dl class="layui-nav-child">-->
        <!--                <dd><a href="login.php" target="aframe">用户登录</a></dd>-->
        <!--                <dd><a href="enter.php" target="aframe">用户注册</a></dd>-->
        <!--            </dl>-->
        <!--        </li>-->

        <li class="layui-nav-item layui-nav-itemed">
            <a href="javascript:;">商品</a>
            <dl class="layui-nav-child">
                <dd><a href="goodslist.php" target="aframe">商品列表</a></dd>
                <dd><a href="goodsadd.php" target="aframe">商品添加</a></dd>
            </dl>
        </li>

        <li class="layui-nav-item layui-nav-itemed">
            <a href="javascript:;">分类添加</a>
            <dl class="layui-nav-child">
                <dd><a href="typelist.php" target="aframe">分类列表</a></dd>
                <dd><a href="typeadd.php" target="aframe">分类添加</a></dd>
            </dl>
        </li>
    </ul>

</div>

<div class="layui-body">
    <ul class="layui-nav " lay-bar="disabled">
        <li class="layui-nav-item" lay-unselect="">
            <a href="javascript:;"><img src="avatar.jpg" class="layui-nav-img">
                <?php
                include_once "Drive/Cookie.php";

                use Driver\Cookie;

                $cookie = new Cookie();
                $Login = $cookie->CheckCookie();
                echo $Login;
                ?>
            </a>
            <dl class="layui-nav-child">
                <dd style="text-align: center;"><a id="cleareLogin" href="">退出登录</a></dd>
            </dl>
        </li>
    </ul>
    <iframe name="aframe" id="frameId" src="typelist.php" style="width:90%;height:90%;"></iframe>
</div>

</body>


<script>
    //注意：导航 依赖 element 模块，否则无法进行功能性操作
    layui.use(['element', 'jquery', 'layer'], function () {
        let element = layui.element;
        let $ = layui.$;
        let layer = layui.layer;
        $(document).ready(function () {
            $("#jump").bind('click', function () {
                $("#frameId").attr("src", "https://www.yuhenm.com");
            });
        });

        // 判断是否登录，未登录强制弹窗输入登录/注册
        let BoolLogin = `<?php
        // 尝试获取是否登录
        echo $Login;
        ?>`
        if (BoolLogin === "未登录") {
            // 跳转到login.php
            location.href = "login.php";
        }


        $('#cleareLogin').on("click", function () {
            document.cookie = "name=; expires=Thu, 01 Jan 1970 00:00:00 GMT";
            layer.msg('退出登录成功', function () {
                location.href = "login.php";
            });
        });


    });


</script>
</html>