<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>登录</title>
    <link rel="stylesheet" href="layui/css/layui.css">
    <link rel="stylesheet" href="web/login.css">
</head>
<body>
<!-- 粒子背景 -->
<div id="particles-js">
    <canvas class="particles-js-canvas-el" width="1920" height="1080" style="width: 100%; height: 100%;"></canvas>
</div>
<!-- 登录表单 -->
<form class="layui-form login-form layui-panel" autocomplete="off">
    <h1 class="login-title"><img style="height: 30px" src="web/logo.png" alt="logo"> <span>用户登录/<a
                    href="regist.php">用户注册</a></span></h1>
    <!-- 用户名 -->
    <div class="layui-form-item"><i class="layui-icon layui-icon-username"></i>
        <input type="text" name="username" id="username" placeholder="请输入用户名" class="layui-input"
               lay-verify="required"
               lay-reqtext="请输入用户名">
    </div>
    <!-- 密码 -->
    <div class="layui-form-item"><i class="layui-icon layui-icon-password"></i>
        <input type="password" name="password" id="password" placeholder="请输入密码" class="layui-input"
               lay-verify="required"
               lay-reqtext="请输入密码">
    </div>
    <!-- 验证码 -->
    <div class="layui-form-item"><i class="layui-icon layui-icon-vercode"></i>
        <input type="text" name="captcha" placeholder="请输入验证码" class="layui-input" lay-verify="required"
               lay-reqtext="请输入验证码">
        <img id="captcha" src="web/captcha.jpg" alt="验证码生成失败！"></div>
    <div class="layui-form-item">
        <button class="layui-btn layui-btn-normal layui-btn-fluid" id="submit" type="button" lay-submit=""
                lay-filter="login"> 登录
        </button>
    </div>
</form>
<script src="layui/layui.js"></script>
<script src="web/particles.js"></script>
<script>
    layui.use(["form", 'jquery', 'layer'], function () {
        let $ = layui.$;
        const form = layui.form;
        //监听提交
        form.on("submit(login)", function (data) {
            //调用登录接口
            $.ajax({
                url: 'web/apiLogin.php',
                type: 'post',
                dataType: 'text',
                data: {
                    uname: $('#username').val(),
                    pwd: $('#password').val(),
                },

                success: function (data) {
                    console.log("data", data)
                    if (data !== "0") {
                        layer.msg('登录成功');
                        document.cookie = "name=" + data;
                        location.href = "index.php";
                        return
                    } else {
                        layer.msg('登录失败');
                    }
                }
            })
            //防止页面跳转
            return false;
        });

        //按下回车登录
        $("body").on("keyup", (e) => {
            if (e.which === 13) {
                $("#submit").trigger("click");
            }
        });

        //点击刷新验证码
        $("#captcha").on("click", (e) => {
            e.target.src += "?" + Math.random();
        });

    });


</script>
</body>
</html>

<!---->
<!--<html>-->
<!--<head>-->
<!--    <title>用户登录</title>-->
<!--    <meta http-equiv="content-type" content="text/html;charset=utf-8"/>-->
<!--</head>-->
<!--<body>-->
<!---->
<!--<form method="post">-->
<!--    <table>-->
<!--        <tr>-->
<!--            <td>账号：</td>-->
<!--            <td><input type="text" name="user"></td>-->
<!--        </tr>-->
<!--        <tr>-->
<!--            <td>姓名：</td>-->
<!--            <td><input type="text" name="uname"></td>-->
<!--        </tr>-->
<!--        <tr>-->
<!--            <td>密码：</td>-->
<!--            <td><input type="text" name="pwd"></td>-->
<!--        </tr>-->
<!--        <tr>-->
<!--            <td colspan="2" align="center">-->
<!--                <input type="reset" value="重置">-->
<!--                <input type="submit" name="s" value="登录">-->
<!--            </td>-->
<!--        </tr>-->
<!--    </table>-->
<!--</form>-->
<?php
//include_once("Drive/mysql.php");
//
//use Driver\MysqlDriver;
//
//if (isset($_POST["s"])) {
//    $user = $_POST["user"];
//    $uname = $_POST["uname"];
//    $pwd = $_POST["pwd"];
//    $New = new MysqlDriver();
//    $query = "select * from users where uid='$user' and uname='$uname' and  password='$pwd' ";
////      mysqli_query($conn,"set names utf8");//数据库中文乱码处理
//    $result = mysqli_query($New->New(), $query);  //执行查询命令
//    $arr = mysqli_fetch_assoc($result);
//    mysqli_free_result($result); //关闭结果集
////      mysqli_close($conn); //关闭连接
//    if ($arr) {
//        setcookie("name", $arr["uname"]);
//
//        echo "<script>
//      location.href='typelist.php';
//      alert('登录成功')
//      </script>";
//        var_dump("已登录成功");
//    } else {
//        echo "<script>
//      alert('登录失败')
//      </script>";
//
//    }
//
//}
//?>
<!--</body>-->
<!--</html>-->