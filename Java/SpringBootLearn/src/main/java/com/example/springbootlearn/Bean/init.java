package com.example.springbootlearn.Bean;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

// 新建一个配置组件类，供自定义测试
@Configuration
public class init {
    // 标志为Init的组件名
    @Bean("init")
    public init Init() {
        init Init = new init();
        return Init;
    }
}
