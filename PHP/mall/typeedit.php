<html>
<head>
    <title>类别修改</title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8"/>
</head>
<body>
<form>
    <table>
        <tr>
            <td>类别号：</td>
            <td><input type="text" name="user" value=<?php if (isset($_GET['tid'])) echo $_GET["tid"] ?>></td>
        </tr>
        <tr>
            <td>类别名：</td>
            <td><input type="text" name="pwd" value=<?php if (isset($_GET['tname'])) echo $_GET["tname"] ?>></td>
        </tr>
        <tr>
            <td>简介：</td>
            <td><input type="text" name="sex" value=<?php if (isset($_GET['note'])) echo $_GET["note"] ?>></td>
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
    $tid = $_GET["user"];
    $tname = $_GET["pwd"];
    $note = $_GET["sex"];
    $New = new MysqlDriver();

    $query = "update type set tid='$tid',tname='$tname' where note='$note'";
    $result = mysqli_query($New->New(), $query);  //执行查询命令
    if ($result) {
        echo "<script>
         location.href='typelist.php'		
		</script>";
    } else {
        echo "<script>
             alert('修改失败')
            </script>";
    }
}
?>
</body>
</html>
