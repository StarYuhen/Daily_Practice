package com.example.springbootlearn;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

// 定义控制器方法
@RestController
public class HelloController {

    // 使用注解反射配置文件内容
    @Value("${name}")
    private String name;

    // 配置请求路径，router
    @RequestMapping("/hello")
    public String Hello() {
        return "Hello SpringBoot " + name;
    }
}
