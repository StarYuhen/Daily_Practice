<html>
<head>
    <meta content="text/html" charset="utf-8">
</head>

<body>

<from action="">
    请输入商品原价:<input type="text" name="price"/>
    折扣:<select name="discount">
        <option>九折</option>
        <option>八五折</option>
        <option>八折</option>
        <option>七五折</option>
        <option>七折</option>
    </select>
    <input type="submit" name="size" value="计算"/><br>
    促销价:<input type="text" name="dprice" >
</from>


<?php
if (isset($_GET["size"])){
    $Original_price=$_GET["price"];
    $discount=$_GET["discount"];

    switch($discount){
        case "九折":
            $Current_price=$Original_price*90/100;
            break;
        case "八五折":
            $Current_price=$Original_price*85/100;
            break;
        case "八折":
            $Current_price=$Original_price*80/100;
            break;
        case "七五折":
            $Current_price=$Original_price*75/100;
            break;
        default:
            $Current_price=$Original_price*70/100;
    }
//    echo"<script>
//	    document.getElementsByName('dprice')[0].value=$Current_price;
//	</script>";
    echo "<script>
       document.getElementsByName('dprice')[0].value=$Current_price;
   </script>";
}
?>

</body>
</html>


