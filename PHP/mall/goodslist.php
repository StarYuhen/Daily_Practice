<html>
<head>
    <title>商品列表</title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8"/>
</head>
<body>
<table border="1" align="center" cellspacing="0" cellpadding="0">
    <tr>
        <th>商品号</th>
        <th>商品名</th>
        <th>商品单价</th>
        <th>商品库存量</th>
        <th>是否上架</th>
        <th>是否上首页推荐</th>
        <th>所属类别</th>
    </tr>
    <?php
    include_once("Drive/mysql.php");

    use Driver\MysqlDriver;

    $New = new MysqlDriver();
    // $conn=mysqli_connect("localhost","root","abc123456");
    $conn = $New->New();
    // mysqli_select_db($conn,"market");//选择market数据库
    $query = "select * from goods";
    $result = mysqli_query($conn, $query);  //执行查询命令
    while ($goods = mysqli_fetch_assoc($result)) //从查询结果中提取数据

    {
        echo "<tr>";
        echo "<td>" . $goods["gid"] . "</td>";
        echo "<td>" . $goods['gname'] . "</td>";
        echo "<td>" . $goods['price'] . "</td>";
        echo "<td>" . $goods['count'] . "</td>";
        echo "<td>" . $goods['on1'] . "</td>";
        echo "<td>" . $goods['recommend'] . "</td>";
        echo "<td>" . $goods['type'] . "</td>";

        echo "<td><a href='goodsedit.php?gid=
		 " . $goods["gid"] . "
		 &gname=
		 " . $goods['gname'] . "
		 &price=
		 " . $goods['price'] . "
		 &count=
		 " . $goods["count"] . "
		 &on1=
		 " . $goods["on1"] . "
		 &recommend=
		 " . $goods["recommend"] . "
		 &type=
		 " . $goods["type"] . "
		 '>编辑</a>
		 
		 
		 
		 
		 
		 <a href='goodsdelete.php?gid=" . $goods["gid"] . "'>删除</a>
		 </td>";
        echo "</tr>";
    }
    mysqli_free_result($result); //关闭结果集
    $New->Close($conn);
    ?>
</table>
</body>
</html>