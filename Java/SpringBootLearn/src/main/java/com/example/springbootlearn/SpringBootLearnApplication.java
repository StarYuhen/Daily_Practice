package com.example.springbootlearn;

import com.example.springbootlearn.Bean.DataBase;
import com.example.springbootlearn.Bean.Pet;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

// 扩展SprintBoot基础扫描包路径
/*
 同时也可以将@SpringBootApplication 注解改为在代码里的注释从而使用包扫描@ComponentScan
@SpringBootConfiguration
@EnableAutoConfiguration
@ComponentScan
 */
@SpringBootApplication(scanBasePackages = "com.example")
public class SpringBootLearnApplication {

    // 自学文档 https://www.yuque.com/atguigu/springboot
    public static void main(String[] args) {
//        ConfigurableApplicationContext run = SpringApplication.run(SpringBootLearnApplication.class, args);
        // 查看框架使用的组件
//        String[] GetSpringBootName = run.getBeanDefinitionNames();
//        for (String Name : GetSpringBootName) {
//            System.out.println(Name);
//        }
//        // 获取指定组件
//        Pet pet = run.getBean("PetConfig", Pet.class);
//        System.out.println(pet);
        // 获取容器中有没有指定组件
//        boolean Pet = run.containsBean("PetConfig");
//        System.out.println(Pet);
        SpringApplication.run(SpringBootLearnApplication.class, args);
    }


}
