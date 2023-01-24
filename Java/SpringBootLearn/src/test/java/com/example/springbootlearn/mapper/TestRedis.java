package com.example.springbootlearn.mapper;

import jakarta.annotation.Resource;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.StringRedisTemplate;

@SpringBootTest
public class TestRedis {
    @Resource
    private RedisTemplate redisTemplate;

    @Resource
    private StringRedisTemplate stringRedisTemplate;

    @Test
    public void TestRedis() {
        // 添加
        redisTemplate.opsForValue().set("engine", "java");
        // 查询
        System.out.println(redisTemplate.opsForValue().get("engine"));
        // 删除
        redisTemplate.delete("engine");
        // 更新
        redisTemplate.opsForValue().set("engine", "Golang");
        System.out.println(redisTemplate.opsForValue().get("engine"));
        // 剩下的StringrRedisTemplate 和上面一样
        // redisTemplate.opsForList() 这个则是列表，多看语法提示
//        redisTemplate.opsForList().set();
    }
}
