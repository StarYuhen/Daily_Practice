<html>
<head>
    <title>商品添加</title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8"/>
</head>
<body>
<form>
    <table>
        <tr>
            <td>商品号：</td>
            <td><input type="text" name="gid"></td>
        </tr>
        <tr>
            <td>商品名：</td>
            <td><input type="text" name="gname"></td>
        </tr>
        <tr>
            <td>商品单价：</td>
            <td><input type="text" name="price"></td>
        </tr>
        <tr>
            <td>商品库存量：</td>
            <td><input type="text" name="count"></td>
        </tr>
        <tr>
            <td>是否上架：</td>
            <td><input type="text" name="on1"></td>
        </tr>
        <tr>
            <td>是否上首页推荐：</td>
            <td><input type="text" name="recommend"></td>
        </tr>
        <tr>
            <td>所属类别：</td>
            <td><input type="text" name="type"></td>
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

    $gid = $_GET["gid"];
    $gname = $_GET["gname"];
    $price = $_GET["price"];
    $count = $_GET["count"];
    $on1 = $_GET["on1"];
    $recommend = $_GET["recommend"];
    $type = $_GET["type"];
    $New = new MysqlDriver();
    $conn = $New->New();
    $query = "insert into goods(gid,gname,price,count,on1,recommend,type) values('$gid','$gname','$price','$count','$on1','$recommend','$type')";
    $result = mysqli_query($conn, $query);  //执行插入命令
    if ($result) {
        echo "<script>
         location.href='goodslist.php'		
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