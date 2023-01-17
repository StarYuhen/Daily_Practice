 <html>
 <head>
   <meta http-equiv="content-type" content="text/html;charset=utf-8"/>
   <form method="post">
     <table>
     <tr><td colspan="2" align="center">
       <input type="submit" name='s' value="退出登录">


     </td></tr>
   </table>
   </form>
 </head>
   欢迎您！
  <?php
   if (isset($_POST["s"])){
  setcookie("name",'',time() - 3600);
  echo "<script>
  alert('请重新登录')
  </script>";
  return;
}

try {
    echo $_COOKIE["name"];
}
catch (Exception $e) {
    echo "<script>
  alert('请重新登录')
  </script>";
    }
  ?>
</html>