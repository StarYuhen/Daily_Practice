package com.example.springbootlearn.Router;

import com.example.springbootlearn.Bean.DataBase;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class init {

    @RequestMapping("/init")
    public String Init() {
        return "init success";
    }


}
