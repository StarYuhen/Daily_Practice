package com.example.springbootlearn.Router;

import com.example.springbootlearn.Bean.PostBody;
import org.springframework.web.bind.annotation.*;

@RestController
public class RequestBodyJson {

    // 请求返回Json
    @ResponseBody // 响应body
    @GetMapping("/request/json")
    public PostBody Json() {
        PostBody body = new PostBody();
        body.setName("test");
        body.setAge(10);
        return body;
    }

    // 请求返回xml
}
