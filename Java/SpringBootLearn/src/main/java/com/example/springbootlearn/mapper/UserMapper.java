package com.example.springbootlearn.mapper;

import com.example.springbootlearn.databases.User;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Import;
import org.springframework.stereotype.Component;

@Component
public interface UserMapper extends BaseMapper<User> {

}
