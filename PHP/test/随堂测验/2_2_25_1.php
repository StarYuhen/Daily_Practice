<html>
<head>
    <meta content="text/html" charset="utf-8">
</head>

<body>
<form action="">
    请输入长方形的长：<input type="text" name="length"><br>
    请输入长方形的宽：<input type="text" name="width"><br>
    <input type="submit" value="计算" name="size"><br>
    长方形的周长是：<input type="text" name="circle"><br>
    长方形的面积是：<input type="text" name="area"><br>
</form>
<?php

if (isset($_GET["size"])){
    $length=$_GET["length"];
    $width=$_GET["width"];
    $circle=($length+$width)*2;
    $area=$length*$width;
    echo "<script >
    // getElementsByName 对相应dom的值赋值计算结果
document.getElementsByName('circle')[0].value=$circle
document.getElementsByName('area')[0].value=$area
</script>";

}




?>
</body>
</html>