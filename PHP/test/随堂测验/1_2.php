<html>
<head>
<title>第一个php页面</title>
</head>
<body>
<p>
    今天是2022年9月5日
</p>

<?php
    date_default_timezone_set("PRC");
    echo date("Y/m/d")
?>
</body>
</html>
