package com.example.springbootlearn.Bean;

import ch.qos.logback.classic.AsyncAppender;
import org.springframework.boot.actuate.autoconfigure.influx.InfluxDbHealthContributorAutoConfiguration;
import org.springframework.boot.autoconfigure.condition.ConditionalOnBean;
import org.springframework.boot.autoconfigure.security.SecurityProperties;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.stereotype.Controller;

import java.util.function.ToDoubleBiFunction;

// 自定义配置组件 配置类==配置组件
/*
    1. 默认使用的Bean，方法是单例方法
    2. 配置类本身也是组件
    3. proxyBeanMethods 是否代理Bean方法，使用他主要是为了保证实例不会重复注册，单实例对象，false则是可以重复声明对象
       用于解决组件依赖之间的问题,同时false 更为轻量，因为不会检查
       这个方法可以理解为单例模式：
        单例类只能有一个实例；单例类必须自己创建自己的唯一实例；单例类必须给所有其他对象提供这一实例。
       当为true时，容器内的组件只会是一个对象的内容，当为false，每创建一个组件对象，都有不同的组件对象内容
    4. @Import({Pet.class, InfluxDbHealthContributorAutoConfiguration.class}) 容器中创建这两个类型的组件
        可以用于任意配置类和组件,默认的组件名为全类名
 */


@Import({AsyncAppender.class})
@Configuration(proxyBeanMethods = false)
// 设置当容器中有该组件，此配置类代码才生效,进行组件注入,，谈过没有指定组件，注释下面的代码全部无法使用
// 如：此配置类没有PetConfig的组件，则MyConfig全部无法使用，若在方法前面加上该注释，没有满足条件则会该方法无法使用
@ConditionalOnBean(name = "PetConfig")
// 增加注册组件配置 开启DataBase的属性配置功能
// @EnableConfigurationProperties(DataBase.class)
public class MyConfig {
    // 设置容器组件，方法名作为组件ID，返回类型是组件类型
    @Bean("PetConfig")
    // 随便用一个包测试类
    public Pet PetConfig() {
        return new Pet();
    }

    @Bean("UserConfig")
    public User UserConfig() {
        return new User();
    }
}
