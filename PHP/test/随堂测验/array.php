<!DOCTYPE html>
<html>
<head>
    <title></title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8">
</head>
<body>


<form>
    学号<input name="no">
    <br>
    姓名<input name="name">
    <br>
    性别<input name="sex">
    <br>
    <input type="submit" value="下一个" name="next">
</form>
<?php
// StarYuhen https://www.yuhenm.com
$students=array(
    array("no"=>"2010893","name"=>"黄德源","sex"=>"男"),
    array("no"=>"2021032","name"=>"吴泽宇","sex"=>"男"),
    array("no"=>"20210078","name"=>"高于禾","sex"=>"男")
);

// 当点击下一步
if (isset($_GET["next"])) {
    $Index= intval($_COOKIE["ArrayIndex"]);
    $url='http://'.$_SERVER['HTTP_HOST'].$_SERVER['PHP_SELF'];
    // 有cookie则操作dom
    if($Index>2){
        setcookie("ArrayIndex", "0");
        echo "<script >
    // getElementsByName 对相应dom的值赋值计算结果
    window.location.href='$url'
</script>";
        exit;
        return;
    }
    $New=$Index+1;

    $stringIndex=strval($New);

    setcookie("ArrayIndex", $stringIndex);
}
// 如果cookie不存在则设置cookie索引为1
try {
    $Index = intval($_COOKIE["ArrayIndex"]);
    // 有cookie则操作dom
    $ArrayFor = $students[$Index];
    $stringJs = "";
    foreach ($ArrayFor as $value) {
        $key = array_keys($ArrayFor, $value);
        $stringJs .= "document.getElementsByName('$key[0]')[0].value='$value'; \n";
    }
    echo "<script >
    // getElementsByName 对相应dom的值赋值计算结果
    $stringJs
</script>";

} catch (Exception $e) {
    setcookie("ArrayIndex", "0");
}








?>



<!--<table>-->
<!--    <tr>-->
<!--        <th>学号</th>-->
<!--        <th>姓名</th>-->
<!--        <th>性别</th>-->
<!--    </tr>-->
<!--    --><?php
//    // $student1=array("no"=>"2021031546","name"=>"马群","sex"=>"女");
//    // $student2=array("no"=>"2021031632","name"=>"曾琴","sex"=>"女");
//    // $student3=array("no"=>"2021050078","name"=>"王龙彪","sex"=>"男");
//    foreach($students as $s ){
//        echo "<tr>";
//        foreach($s as $v){
//            echo "<td>";
//            echo $v;
//            echo "</td>";
//        }
//        echo "</tr>";
//
//    }
//
//
//
//    ?>
<!---->
<!--</table>-->
</body>
</html>