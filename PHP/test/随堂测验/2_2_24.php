<!DOCTYPE html>
<html>
<head>
    <title></title>
    <meta http-equiv="content-type" content="text/html;charset=utf-8">
</head>
<body>
<?php
$students=array(
    array("no"=>"2010893","name"=>"黄德源","sex"=>"男"),
    array("no"=>"2021032","name"=>"吴泽宇","sex"=>"男"),
    array("no"=>"20210078","name"=>"高于禾","sex"=>"男")
);

?>
<table border="1" cellpadding="0" cellpadding="0">
    <tr>
        <th>学号</th>
        <th>姓名</th>
        <th>性别</th>
    </tr>
    <?php
    // $student1=array("no"=>"2021031546","name"=>"马群","sex"=>"女");
    // $student2=array("no"=>"2021031632","name"=>"曾琴","sex"=>"女");
    // $student3=array("no"=>"2021050078","name"=>"王龙彪","sex"=>"男");
    foreach($students as $s ){
        echo "<tr>";
        foreach($s as $v){
            echo "<td>";
            echo $v;
            echo "</td>";
        }
        echo "</tr>";

    }



    ?>

</table>
</body>
</html>