package com.example.springbootlearn.mapper;

import com.baomidou.mybatisplus.generator.FastAutoGenerator;
import com.baomidou.mybatisplus.generator.config.OutputFile;
import com.baomidou.mybatisplus.generator.engine.FreemarkerTemplateEngine;
import com.example.springbootlearn.databases.User;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.Collections;
import java.util.List;

// 使用测试
@SpringBootTest
class UserMapperTest {
    @Autowired
    private UserMapper userMapper;

    @Test
    public void testSelect() {
        System.out.println(("----- selectAll method test ------"));
        // 查询全部
        /*
        ----- selectAll method test ------
User(id=1, name=Jone, age=18, email=test1@baomidou.com)
User(id=2, name=Jack, age=20, email=test2@baomidou.com)
User(id=3, name=Tom, age=28, email=test3@baomidou.com)
User(id=4, name=Sandy, age=21, email=test4@baomidou.com)
User(id=5, name=Billie, age=24, email=test5@baomidou.com)
         */
        List<User> userList = userMapper.selectList(null);
        Assertions.assertEquals(6, userList.size());
        userList.forEach(System.out::println);
        // 插入一个信息 id=6, name=Yuhen, age=20, email=test6@baomidou.com
//        User user = new User();
//        long id = 6;
//        user.setId(id);
//        user.setAge(20);
//        user.setName("Yuhen");
//        user.setEmail("test6@baomidou.com");
//        userMapper.insert(user);
//        System.out.println("insert success");

        // 更新,删除，查询文档都有就不测试了，差不多
//        userMapper.selectPage()

    }

    // 代码生成器
    @Test
    public void testGroup() {
        FastAutoGenerator.create("jdbc:mysql://localhost:3306/eim", "root", "abc123456")
                .globalConfig(builder -> {
                    builder.author("StarYuhen") // 设置作者
                            .enableSwagger() // 开启 swagger 模式
                            .fileOverride() // 覆盖已生成文件
                            .outputDir("E://"); // 指定输出目录
                })
                .packageConfig(builder -> {
                    builder.parent("com.baomidou.mybatisplus.samples.generator") // 设置父包名
                            .moduleName("system") // 设置父包模块名
                            .pathInfo(Collections.singletonMap(OutputFile.xml, "E://")); // 设置mapperXml生成路径
                })
                .strategyConfig(builder -> {
                    builder.addInclude("user") // 设置需要生成的表名
                            .addTablePrefix("t_", "c_"); // 设置过滤表前缀
                })
                .templateEngine(new FreemarkerTemplateEngine()) // 使用Freemarker引擎模板，默认的是Velocity引擎模板
                .execute();
    }

}