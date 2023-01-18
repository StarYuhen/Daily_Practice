package com.example.springbootlearn.Bean;

import lombok.Data;
import lombok.Getter;
import lombok.Setter;
import org.springframework.beans.factory.config.YamlPropertiesFactoryBean;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.PropertySource;
import org.springframework.stereotype.Component;

//TODO 自动解析配置文件
// 声明为组件
@Component
// 获取配置文件信息,由于是有前缀的形式，所以需要使用此函数
@PropertySource(value = "classpath:application.yaml")
@ConfigurationProperties(prefix = "database")
// 使用注解自行注入setter和getter 方法，虽然可以自动生成，但是还是可以用lombok方法
@Data
public class DataBase {
    private String user;
    private String password;

    public DataBase() {
    }


    public String DataBaseConfig() {
        return user + password;
    }
// 返回database字段配置内容 使用lombok 简易化
//
//    public String getUser() {
//        return user;
//    }
//
//    public void setUser(String user) {
//        this.user = user;
//    }
//
//    public String getPassword() {
//        return password;
//    }

//    public void setPassword(String password) {
//        this.password = password;
//    }

}
