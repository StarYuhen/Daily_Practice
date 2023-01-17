<html>
<head>
    <title>类别修改</title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8"/>
</head>
<body>
<form>
    <table>
        <tr>
            <td>商品号：</td>
            <td><input type="text" name="gid" value=<?php if (isset($_GET['gid'])) echo $_GET["gid"] ?>></td>
        </tr>
        <tr>
            <td>商品名：</td>
            <td><input type="text" name="gname" value=<?php if (isset($_GET['gname'])) echo $_GET["gname"] ?>></td>
        </tr>
        <tr>
            <td>商品单价：</td>
            <td><input type="text" name="price" value=<?php if (isset($_GET['price'])) echo $_GET["price"] ?>></td>
        </tr>
        <tr>
            <td>商品库存量：</td>
            <td><input type="text" name="count" value=<?php if (isset($_GET['count'])) echo $_GET["count"] ?>></td>
        </tr>
        <tr>
            <td>是否上架：</td>
            <td><input type="text" name="on1" value=<?php if (isset($_GET['on1'])) echo $_GET["on1"] ?>></td>
        </tr>
        <tr>
            <td>是否上首页推荐：</td>
            <td><input type="text" name="recommend"
                       value=<?php if (isset($_GET['recommend'])) echo $_GET["recommend"] ?>></td>
        </tr>
        <tr>
            <td>所属类别：</td>
            <td><input type="text" name="type" value=<?php if (isset($_GET['type'])) echo $_GET["type"] ?>></td>
        </tr>
        <tr>
            <td colspan="2" align="center">

                <input type="submit" name="s" value="修改">
            </td>
        </tr>
    </table>
</form>
<?php
include_once("Drive/mysql.php");

use Driver\MysqlDriver;

if (isset($_GET['s'])) {
    $gid = $_GET["gid"];
    $gname = $_GET["gname"];
    $price = $_GET["price"];
    $count = $_GET["count"];
    $on1 = $_GET["on1"];
    $recommend = $_GET["recommend"];
    $type = $_GET["type"];
    $New = new MysqlDriver();
    $conn = $New->New();
    $query = "update goods set gname='$gname',price='$price',count='$count',on1='$on1',recommend='$recommend' where type='$type' and gid='$gid' ";
    // mysqli_query($New->New(),"set names utf8");//数据库中文乱码处理
    $result = mysqli_query($conn, $query);  //执行查询命令
    if ($result) {
        echo "<script>
         location.href='goodslist.php'		
		</script>";
    } else {
        echo "<script>
             alert('修改失败')
            </script>";
    }
    $New->Close($conn);
}
?>
</body>
</html>