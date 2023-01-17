<html>
<head>
    <title>分类添加</title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8"/>
</head>
<body>
<form>
    <table>
        <tr>
            <td>类别号：</td>
            <td><input type="text" name="user"></td>
        </tr>
        <tr>
            <td>类别名：</td>
            <td><input type="text" name="pwd"></td>
        </tr>
        <tr>
            <td>简介：</td>
            <td><input type="text" name="sex"></td>
        </tr>
        <tr>
            <td colspan="2" align="center">
                <input type="reset" value="重置">
                <input type="submit" name="s" value="添加">
            </td>
        </tr>
    </table>
</form>
<?php
include_once("Drive/mysql.php");

use Driver\MysqlDriver;

if (isset($_GET["s"])) {
    $no = $_GET["user"];
    $name = $_GET["pwd"];
    $note = $_GET["sex"];
    $New = new MysqlDriver();
    $conn = $New->New();
    $query = "insert into type(tid,tname,note) values('$no','$name','$note')";
    // var_dump($query);
    // mysqli_query(,"set names utf8");//数据库中文乱码处理
    $result = mysqli_query($conn, $query);  //执行查询命令
    if ($result) {
        echo "<script>
      location.href='typelist.php';
      alert('添加成功')
      </script>";
    } else {
        echo "<script>
      alert('添加失败')
      </script>";

    }

    $New->Close($conn);

}
?>
</body>
</html>
</html>