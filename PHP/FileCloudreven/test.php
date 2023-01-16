<?php
 $str='
猫猫和修勾,理发店里拍的猫猫和修勾,https://api.yuhenm.com/api/file/source?name=1da8261f-c514-4a6e-8541-ee3541c6ecf4.jpg|
学校里遇见的猫猫,猫猫真可爱,https://api.yuhenm.com/api/file/source?name=0e490e72-8e6f-435c-a3b3-49e7b0124030.jpg|
学校的蓝天,真的很漂亮,https://api.yuhenm.com/api/file/source?name=95ab02d4-7ef9-42a2-8416-c48e20f77b45.jpg|
那一束光,学校里面拍的,https://api.yuhenm.com/api/file/source?name=efb4c047-eaac-4283-9355-5c6349cd3437.jpg|
白云,很有意境,https://api.yuhenm.com/api/file/source?name=f51b3f45-218d-4700-beb7-697e8fafb10d.jpg|
蒙蒙的天空,有个太阳,https://api.yuhenm.com/api/file/source?name=31563385-e853-4da7-845d-2cc0e6daa62a.jpg|
白云遮住的太阳,灰蒙蒙的诶,https://api.yuhenm.com/api/file/source?name=2d0dc96b-0d76-4291-9caa-a84d9cda11ac.jpg|
鳞片状的白云,很好看,https://api.yuhenm.com/api/file/source?name=7c12157c-e247-4314-9aec-e3d22c209923.jpg|
爬赵公山的图片,老累了,https://api.yuhenm.com/api/file/source?name=a919c152-9927-4e69-bc60-3a62f30f6826.jpg|
爬赵公山的图片,老累了,https://api.yuhenm.com/api/file/source?name=bf8781ae-5eb3-4781-b751-fb841c3e248a.jpg|
爬赵公山的图片,老累了,https://api.yuhenm.com/api/file/source?name=2ebb4a7e-3f46-43c4-888a-beab7801424d.jpg|
';

$data=explode( '|', $str );
foreach ($data as $Array){
    $ArrayURL=explode( ',', $Array);
    var_dump($ArrayURL);
}
