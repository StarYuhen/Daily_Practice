package com.example.springbootlearn.Router;

import com.example.springbootlearn.Bean.DataBase;
import com.example.springbootlearn.Bean.PostBody;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.web.server.Cookie;
import org.springframework.web.bind.annotation.*;

import javax.xml.namespace.QName;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;

@RestController
// 注入日志
@Slf4j

public class init {
    // 自动注入字段
    @Autowired
    DataBase dataBase;

    // 路径权重优先级 controller>resources 目录的资源路径
    @RequestMapping("/init")
    public String Init() {
//        System.out.println(dataBase.DataBaseConfig());
        log.info(dataBase.DataBaseConfig());
        log.info("正常配置成功");
        return "init success";
    }

    // 定义一个get模式的
    @RequestMapping(value = "Name", method = RequestMethod.GET)
    public String GETName() {
        log.info("get name success");
        return "test";
    }

    // 定义一个post模式的
    @RequestMapping(value = "Name", method = RequestMethod.POST)
    public String POSTName() {
        log.info("post name success");
        return "test";
    }

    // 还有PUT和DELETE模式，并不常用

    // 使用@PathVariable
    /*
    关于这些方法的功能可以看注释
    @PathVariable 获取请求路径对应的值
        还可以使用Map接收传来的说有值，如Map<String,String>
    @RequestHeader 获取请求的头部参数，User-Agent 表示获取请求的浏览器标识
        还可以使用Map接收传来的说有值，如Map<String,String>
    @RequestParam 获取请求的Param值，如Get的/get/test?age=10&name="Test"
    @CookieValue 获取请求的Cookie值
    @RequestAttribute 获取域属性
    @RequestBody 获取post请求的Json值 会自动序列化

     */
    @GetMapping("/get/{id}/test/{name}")
    public Map<String, Object> GetPathController(
            @PathVariable("id") Integer id,
            @PathVariable Map<String, String> PathMap,
            @RequestHeader("User-Agent") String Agent,
            @RequestHeader Map<String, String> HeaderMap,
            @RequestParam("age") Integer age,
            @RequestParam Map<String, String> ParamMap,
            @CookieValue("test") String test
    ) {
        // 创建Map的泛型接收参数 object 对应Go的 interface和any
        Map<String, Object> map = new HashMap<>();
        map.put("id", id);
        map.put("User-Agent", Agent);
        map.put("age", age);
        map.put("test", test);

        // 打印获取的集合对象
        log.info("PathMap:{},HeaderMap:{},ParamMap:{}", PathMap, HeaderMap, ParamMap);
        // PathMap:{name=yuhen, id=3},HeaderMap:{cookie=test=yuhen;, user-agent=PostmanRuntime/7.30.0, accept=*/*, postman-token=9b3f33ca-68ad-451e-8874-0172fd87ace7, host=127.0.0.2:8888, accept-encoding=gzip, deflate, br, connection=keep-alive},ParamMap:{age=10}
        return map;
    }

    // 使用post方法获取body值
    @PostMapping("/bodyTest")
    public String PostTest(
            @RequestBody String body
    ) {
        return body;
    }


}
