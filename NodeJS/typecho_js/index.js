const request = require('request');
const http = require('http');
const axios = require('axios');
axios.defaults.baseURL = 'https://www.yuhenm.com'; //设置请求url
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'; //请求类型


//集成百度，必应提交接口
http.createServer(function(request, response) {
    response.writeHead(200, { 'Content-Type': 'text/plain' });
    response.end('Hello World\n');
    let url_post = "https://www.yuhenm.com" + request.url;
    console.log("请求收录地址", url_post);
    BaiDuURL_KK(url_post, url_post);
}).listen(200, "0.0.0.0");


function BaiDuURL_KK(url, url_post) {
    //检测百度是否收录
    axios.get('/baiduPHP.php?domain=' + url, {}).then(function(response) {
        console.log(response.data);
        if (response.data.msg !== "该URL已被百度收录！") {
            console.log("开始提交百度收录");
            postBaiDu_url(url_post);
            postBin_url(url_post);
        }
    })
}


function postBin_url(url) {
    //必应提交api
    console.log('开始提交必应API');
    let myJson = {
        "siteUrl": "https://www.yuhenm.com",
        "urlList": [url]
    };
    request({
        url: 'https://ssl.bing.com/webmaster/api.svc/json/SubmitUrlbatch?apikey=' + '',
        /* xxx需替换为你的key */
        method: "POST",
        json: true, // <--Very important!!!  允许跨域
        body: myJson
    }, function(error, response, body) {
        console.log('提交必应结果', body);
    });
}


function postBaiDu_url(url) {
    console.log('开始提交百度API');
    const req = http.request({
        host: 'data.zz.baidu.com',
        path: '/urls?site=https://www.yuhenm.com&token=',
        method: 'post',
        headers: {
            'User-Agent': 'curl/7.12.1',
            "Content-Type": "text/plain",
            "Content-Length": url.length
        }
    }, function(res) {
        res.setEncoding('utf8');
        res.on('data', function(data) {
            console.log("提交百度API返回结果：", res.statusCode, ' - ', data);
        });
    });

    req.on('error', (e) => {
        console.log('提交百度API请求失败：', e);
    });
    req.write(url)
    req.end();
}