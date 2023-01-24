package com.example.springbootlearn.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.example.springbootlearn.databases.User;
import org.springframework.stereotype.Component;


@Component
public interface UserMapper extends BaseMapper<User> {

}

