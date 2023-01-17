package com.example.springbootlearn;

import com.example.springbootlearn.Bean.DataBase;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class SpringBootLearnApplicationTests {
    DataBase dataBase;

    @Test
    void contextLoads() {
        System.out.println(dataBase.DataBaseConfig());
    }

}
