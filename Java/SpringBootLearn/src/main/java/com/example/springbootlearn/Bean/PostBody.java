package com.example.springbootlearn.Bean;

import lombok.Data;

// post请求获取传入的值模型
@Data
public class PostBody {
    private String name;
    private int age;
}
