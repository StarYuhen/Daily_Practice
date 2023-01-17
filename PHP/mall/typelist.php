<html>
<head>
    <title>商品列表表</title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8"/>
</head>
<body>
<table border="1" align="center" cellspacing="0" cellpadding="0">
    <tr>
        <th>类别号</th>
        <th>类别名</th>
        <th>简介</th>
    </tr>
    <?php
    include_once("Drive/mysql.php");

    use Driver\MysqlDriver;

    $query = "select * from type";
    $New = new MysqlDriver();
    $conn = $New->New();
    $result = mysqli_query($conn, $query);  //执行查询命令
    while ($type = mysqli_fetch_assoc($result)) //从查询结果中提取数据

    {
        echo "<tr>";
        echo "<td>" . $type["tid"] . "</td>";
        echo "<td>" . $type['tname'] . "</td>";
        echo "<td>" . $type['note'] . "</td>";
        echo "<td><a href='typeedit.php?tid=" . $type["tid"] . "
     &tname=" . $type['tname'] . "&note=" . $type['note'] . "'>编辑</a>
     <a href=typedelete.php?tid=" . $type["tid"] . ">删除</a>
     </td>";
        echo "</tr>";
    }
    mysqli_free_result($result); //关闭结果集
    $New->Close($conn);
    ?>
</table>
</body>
</html>