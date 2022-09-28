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
// 获取 dom然后赋值
document.getElementsByName('circle')[0].value=$circle
document.getElementsByName('area')[0].value=$area
</script>";
//    echo "长方形的周长是：$circle";
//    echo  "长方形的面积是：$area";
}


//    echo "<script >
//alert('长方形的周长是：'$circle \n '长方形的面积是：'$area)
//</script>";



?>
</body>
</html>