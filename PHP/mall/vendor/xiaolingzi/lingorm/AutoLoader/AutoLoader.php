<?php
class AutoLoader
{
	public static function classLoader($className)
	{
	    $dir=dirname(__DIR__);
	    $classNameArr=explode("\\", $className);
	    $count=count($classNameArr);
	    for($i=1;$i<$count;$i++)
	    {
	    	$dir.="/".$classNameArr[$i];
	    	
	    }
		$filename=$dir.".php";
		if(is_file($filename))
		{
		  require_once $filename;
		}
	}
}

spl_autoload_register(array('AutoLoader', 'classLoader'));

