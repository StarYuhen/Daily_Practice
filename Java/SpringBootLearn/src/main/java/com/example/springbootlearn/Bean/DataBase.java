package com.example.springbootlearn.Bean;

import org.springframework.beans.factory.config.YamlPropertiesFactoryBean;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.PropertySource;
import org.springframework.stereotype.Component;

//TODO 自动解析配置文件
// 声明为组件
@Component
// 获取配置文件信息,由于是有前缀的形式，所以需要使用此函数
@PropertySource(value = "classpath:application.yaml")
@ConfigurationProperties(prefix = "database")
public class DataBase {
    private String user;
    private String password;

    // 返回database字段配置内容
    public String DataBaseConfig() {
        return user + password;
    }
}
