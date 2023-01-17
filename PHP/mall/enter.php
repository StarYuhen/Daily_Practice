<html>
<head>
    <title>用户添加</title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8"/>
</head>
<body>
<form>
    <table>
        <tr>
            <td>账号：</td>
            <td><input type="text" name="user"></td>
        </tr>
        <tr>
            <td>姓名：</td>
            <td><input type="text" name="name"></td>
        </tr>
        <tr>
            <td>密码：</td>
            <td><input type="text" name="sex"></td>
        </tr>
        <tr>
            <td colspan="2" align="center">
                <input type="reset" value="重置">
                <input type="submit" name="s" value="注册">
            </td>
        </tr>
    </table>
</form>
<?php
include_once("Drive/mysql.php");

use Driver\MysqlDriver;

if (isset($_GET["s"])) {
    $no = $_GET["user"];
    $name = $_GET["name"];
    $sex = $_GET["sex"];
    $New = new MysqlDriver();
    $query = "insert into users(uid,uname,password) values('$no','$name','$sex')";
    $result = mysqli_query($New->New(), $query);  //执行插入命令
    if ($result) {
        echo "<script>
         location.href='login.php'
         alert('注册成功')		
		</script>";
    } else {
        echo "<script>
             alert('注册失败')
            </script>";
    }
}
?>
</body>
</html>