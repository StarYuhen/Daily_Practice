<html>
<head>
    <meta http-equiv="content-type" content="charset=utf-8"/>
    </title>四则运算</title>
</head>
<body>
<form action="">
    <input type="text" name="sz1">
    <select name="ysf">
        <option>+</option>
        <option>-</option>
        <option>*</option>
        <option>/</option>
        <lect>
            <input type="text" name="sz2">
            <input type="submit" value="="name="s">
            <input type="text" name="jg">
</form>
<?php
//判断是否提交
if(!empty($_GET['s'])){
    $sz1=$_GET['sz1'];
    $sz2=$_GET['sz2'];
    $ysf=$_GET['ysf'];
    switch($ysf)
    {
        case "+":
            $jg=$sz1+$sz2;
            break;
        case "-":
            $jg=$sz1-$sz2;
            break;
        case "*":
            $jg=$sz1*$sz2;
            break;
        case "/":
            $jg=$sz1/$sz2;
            break;
    }
    echo "<script type=\"text/javascript\">
					document.getElementsByName('jg')[0].value=$jg;
					document.getElementsByName('sz1')[0].value=$sz1;
					document.getElementsByName('sz2')[0].value=$sz2;
					</script>
				";
}
?>
</body>
<html>
