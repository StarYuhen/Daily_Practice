<?php
if (!defined('__TYPECHO_ROOT_DIR__')) exit;

/**
 * FileCloudreven typecho 的附件上传插件
 *
 * @package FileCloudreven
 * @author StarYuhen(玉衡)
 * @version 1.1.0
 * @link https://www.yuhenm.com
 */
class FileCloudreven_Plugin implements Typecho_Plugin_Interface
{
    const UPLOAD_DIR  = '/usr/uploads'; //上传文件目录路径
    const PLUGIN_NAME = 'FileCloudreven'; //插
    const PLUGIN_NAME_IMG = 'FileCloudreven'; //插件名称

    public static function activate()
    {
        Typecho_Plugin::factory('Widget_Upload')->uploadHandle     = array(__CLASS__, 'uploadHandle');
        Typecho_Plugin::factory('Widget_Upload')->modifyHandle     = array(__CLASS__, 'modifyHandle');
        Typecho_Plugin::factory('Widget_Upload')->deleteHandle     = array(__CLASS__, 'deleteHandle');
        Typecho_Plugin::factory('Widget_Upload')->attachmentHandle = array(__CLASS__, 'attachmentHandle');
    }

    public static function deactivate()
    {

    }

    public static function config(Typecho_Widget_Helper_Form $form)
    {
        $html = <<<HTML
<p>使用魔改化的Cloudreven的接口实现的附件上传功能，可用于图床类公开储存</p>
HTML;
        $desc = new Typecho_Widget_Helper_Form_Element_Text('desc', NULL, '', '插件介绍：', $html);
        $form->addInput($desc);

        $api = new Typecho_Widget_Helper_Form_Element_Text('api', NULL, 'http://api.yuhenm.com', 'Api：', '只需填写域名包含 http 或 https 无需<code style="padding: 2px 4px; font-size: 90%; color: #c7254e; background-color: #f9f2f4; border-radius: 4px;"> / </code>结尾<br><code style="padding: 2px 4px; font-size: 90%; color: #c7254e; background-color: #f9f2f4; border-radius: 4px;">示例地址：http://api.yuhenm.com</code>');
        $form->addInput($api);
        $token = new Typecho_Widget_Helper_Form_Element_Text('token', NULL, '', '用户名：', '如果为空，则上传的所属用户为游客。<a target="_blank" href="https://cloud.yuhenm.com/">用户名在线获取</a><br>如有需要请按示例严格填写>');
        $form->addInput($token);

        $strategy_id = new Typecho_Widget_Helper_Form_Element_Text('strategy_id', NULL, '', '密码：', '');
        $form->addInput($strategy_id);




        echo '<script>window.onload = function(){document.getElementsByName("desc")[0].type = "hidden";}</script>';
    }

    public static function personalConfig(Typecho_Widget_Helper_Form $form)
    {
    }

    public static function uploadHandle($file)/*上传文件*/
    {
        if (empty($file['name'])) {

            return false;
        }

        //获取扩展名
        $ext = self::_getSafeName($file['name']);
        //判定是否是允许的文件类型
        if (!Widget_Upload::checkFileType($ext) || Typecho_Common::isAppEngine()) {

            return false;
        }
        // 判断是否是图片
        if (self::_isImage($ext)) {

            return self::_uploadImg($file, $ext);
        }

        return self::_uploadOtherFile($file, $ext);
    }

    public static function deleteHandle(array $content): bool/*删除文件*/
    {

        $ext = $content['attachment']->type;//读取文件扩展名

        if (self::_isImage($ext)) {//判断是否是图片

            return self::_deleteImg($content);
        }

        return unlink($content['attachment']->path);
    }

    public static function modifyHandle($content, $file)
    {
        if (empty($file['name'])) {

            return false;
        }
        $ext = self::_getSafeName($file['name']);
        if ($content['attachment']->type != $ext || Typecho_Common::isAppEngine()) {

            return false;
        }

        if (!self::_getUploadFile($file)) {

            return false;
        }

        if (self::_isImage($ext)) {//判断是否是图片
            self::_deleteImg($content);

            return self::_uploadImg($file, $ext);
        }

        return self::_uploadOtherFile($file, $ext);
    }

    public static function attachmentHandle(array $content): string
    {
        $arr = unserialize($content['text']);
        /*在text字段中截取“.”后的3个字符作为文件扩展名*/
        $text = strstr($content['text'],'.');
        $ext = substr($text,1,3);
        if (self::_isImage($ext)) {

            return $content['attachment']->path ?? '';
        }

        $ret = explode(self::UPLOAD_DIR, $arr['path']);
        return Typecho_Common::url(self::UPLOAD_DIR . @$ret[1], Helper::options()->siteUrl);/**/
    }

    private static function _getUploadDir($ext = ''): string /*上传目录*/
    {
        if (self::_isImage($ext)) {
            $url = parse_url(Helper::options()->siteUrl);
            $DIR = str_replace('.', '_', $url['host']);
            return '/' . $DIR . self::UPLOAD_DIR;
        } elseif (defined('__TYPECHO_UPLOAD_DIR__')) {
            return __TYPECHO_UPLOAD_DIR__;
        } else {
            $path = Typecho_Common::url(self::UPLOAD_DIR, __TYPECHO_ROOT_DIR__);
            return $path;
        }
    }

    private static function _getUploadFile($file): string
    {
        return $file['tmp_name'] ?? ($file['bytes'] ?? ($file['bits'] ?? ''));
    }

    private static function _getSafeName(&$name): string
    {
        $name = str_replace(array('"', '<', '>'), '', $name);
        $name = str_replace('\\', '/', $name);
        $name = false === strpos($name, '/') ? ('a' . $name) : str_replace('/', '/a', $name);
        $info = pathinfo($name);
        $name = substr($info['basename'], 1);

        return isset($info['extension']) ? strtolower($info['extension']) : '';
    }

    private static function _makeUploadDir($path): bool
    {
        $path    = preg_replace("/\\\+/", '/', $path);
        $current = rtrim($path, '/');
        $last    = $current;

        while (!is_dir($current) && false !== strpos($path, '/')) {
            $last    = $current;
            $current = dirname($current);
        }

        if ($last == $current) {
            return true;
        }

        if (!@mkdir($last)) {
            return false;
        }

        $stat  = @stat($last);
        $perms = $stat['mode'] & 0007777;
        @chmod($last, $perms);

        return self::_makeUploadDir($path);
    }

    private static function _isImage($ext): bool
    {
//        $img_ext_arr = array('gif','jpg','jpeg','png','tiff','bmp','ico','psd','webp','JPG','BMP','GIF','PNG','JPEG','ICO','PSD','TIFF','WEBP','zip','rar','mp4'); //允许的图片扩展名
//        return in_array($ext, $img_ext_arr);
        return true;
    }

    private static function _uploadOtherFile($file, $ext)
    {
        $dir = self::_getUploadDir($ext) . '/' . date('Y') . '/' . date('m');
        if (!self::_makeUploadDir($dir)) {

            return false;
        }

        $path = sprintf('%s/%u.%s', $dir, crc32(uniqid()), $ext);
        if (!isset($file['tmp_name']) || !@move_uploaded_file($file['tmp_name'], $path)) {

            return false;
        }

        return [
            'name' => $file['name'],
            'path' => $path,
            'size' => $file['size'] ?? filesize($path),
            'type' => $ext,
            'mime' => @Typecho_Common::mimeContentType($path)
        ];
    }

    private static function _uploadImg($file)
    {


        $options = Helper::options()->plugin(self::PLUGIN_NAME_IMG);
        $api     = $options->api . '/api/file/upload';

        $tmp     = self::_getUploadFile($file);
        if (empty($tmp)) {

            return false;
        }

        $img = $file['name'];//保留图片原始名称到图床
        if (!rename($tmp, $img)) {

            return false;
        }
        $params = ['File' => new CURLFile($img)];

        $res = self::_curlPost($api, $params);

        unlink($img);

        if (!$res) {

            return false;
        }


        $json = json_decode($res, true);


        if ($json['status'] === false) {  // 上传失败
            file_put_contents('./usr/plugins/FileCloudreven/msg.log', json_encode($json, 256) . PHP_EOL, FILE_APPEND);
            return false;
        }
        // return $json['msg'];
        $data = $json['data'];/*图床json信息处理*/

        return [
            'img_key' => $data['key'],//图片唯一key
            'img_id' => $data['md5'],//图片md5
            'name'   => $data['originName'],//原始文件名
            'path'   => $data['links'],//图片url
            'size'   => $data['size'],//图片大小
            'type'   => $data['extension'],//图片后缀名
            'mime'   => $data['mimetype'],//图片类型
            'description'  => $data['mimetype'],//图片类型添加到描述
        ];
    }

    private static function _deleteImg(array $content): bool
    {

        $options = Helper::options()->plugin(self::PLUGIN_NAME);

        $api     = $options->api . '/api/file/upload';
        $token   = $options->token;


        $id = $content['attachment']->img_key;
        if (empty($id)) {
            return false;
        }

        $res  = self::_curlDelete($api . '/' . $id, ['key' => $id], $token);
        $json = json_decode($res, true);

        if (!is_array($json)) {

            return false;
        }

        return true;
    }

    private static function _curlDelete($api, $post, $token)
    {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $api);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
        curl_setopt($ch, CURLOPT_FOLLOWLOCATION, false);
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "DELETE");
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_HTTPHEADER, ['Authorization: '. $token]);
        curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36');
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $post);
        $res = curl_exec($ch);
        curl_close($ch);

        return $res;
    }

    private static function _curlPost($api, $post)
    {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $api);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
        curl_setopt($ch, CURLOPT_FOLLOWLOCATION, false);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36');
        curl_setopt($ch, CURLOPT_POST, 1);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $post);
        $res = curl_exec($ch);
        curl_close($ch);

        return $res;
    }
}