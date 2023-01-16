<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>注册</title>
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
    <h1 class="login-title"><img style="height: 30px" src="web/logo.png" alt="logo"> <span>用户注册</span></h1>
    <div class="layui-form-item"><i class="layui-icon layui-icon-release"></i>
        <input type="text" name="email" id="email" placeholder="请输入邮箱地址" class="layui-input"
               lay-verify="required"
               lay-reqtext="请输入邮箱">
    </div>
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
                lay-filter="regist"> 注册
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
        form.on("submit(regist)", function (data) {
            //调用登录接口
            $.ajax({
                url: 'web/apiRegist.php',
                type: 'post',
                dataType: 'text',
                data: {
                    user: $('#email').val(),
                    name: $('#username').val(),
                    pwd: $('#password').val(),
                },

                success: function (data) {
                    console.log("data", data)
                    if (data !== "0") {
                        layer.msg('注册成功');
                        document.cookie = "name=" + data;
                        location.href = "index.php";
                        return
                    } else {
                        layer.msg('注册失败');
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
