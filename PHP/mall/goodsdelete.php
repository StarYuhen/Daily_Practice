<?php
$gid = $_GET["gid"];
include_once("Drive/mysql.php");

use Driver\MysqlDriver;

$New = new MysqlDriver();
$conn = $New->New();
$query = "delete from goods where gid='$gid' and gid='$gid'";//删除商品表中指定类别的商品
$result = mysqli_query($conn, $query);  //执行查询命令
if ($result) {
    echo "<script>
         location.href='goodslist.php'		
		</script>";
} else {
    echo "<script>
             alert('删除失败')
            </script>";
    $New->Close($conn);

}
?>