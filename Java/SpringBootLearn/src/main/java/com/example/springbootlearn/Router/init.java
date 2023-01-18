package com.example.springbootlearn.Router;

import com.example.springbootlearn.Bean.DataBase;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
// 注入日志
@Slf4j
public class init {
    // 自动注入字段
    @Autowired
    DataBase dataBase;

    @RequestMapping("/init")
    public String Init() {
//        System.out.println(dataBase.DataBaseConfig());
        log.info(dataBase.DataBaseConfig());
        return "init success";
    }


}
