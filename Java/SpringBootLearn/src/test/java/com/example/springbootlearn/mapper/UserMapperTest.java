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


/*
测试常用注解
● @Test :表示方法是测试方法。但是与JUnit4的@Test不同，他的职责非常单一不能声明任何属性，拓展的测试将会由Jupiter提供额外测试
● @ParameterizedTest :表示方法是参数化测试，下方会有详细介绍
● @RepeatedTest :表示方法可重复执行，下方会有详细介绍
● @DisplayName :为测试类或者测试方法设置展示名称
● @BeforeEach :表示在每个单元测试之前执行
● @AfterEach :表示在每个单元测试之后执行
● @BeforeAll :表示在所有单元测试之前执行
● @AfterAll :表示在所有单元测试之后执行
● @Tag :表示单元测试类别，类似于JUnit4中的@Categories
● @Disabled :表示测试类或测试方法不执行，类似于JUnit4中的@Ignore
● @Timeout :表示测试方法运行如果超过了指定时间将会返回错误
● @ExtendWith :为测试类或测试方法提供扩展类引用
 */
// 可通过assertions类进行断言
/*
 往测试里面注入参数
@ValueSource: 为参数化测试指定入参来源，支持八大基础类以及String类型,Class类型
@NullSource: 表示为参数化测试提供一个null的入参
@EnumSource: 表示为参数化测试提供一个枚举入参
@CsvFileSource：表示读取指定CSV文件内容作为参数化测试入参
@MethodSource：表示读取指定方法的返回值作为参数化测试入参(注意方法返回需要是一个流)
 */
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