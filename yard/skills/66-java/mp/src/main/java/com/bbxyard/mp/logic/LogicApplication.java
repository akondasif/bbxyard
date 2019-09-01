package com.bbxyard.mp.logic;


import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("com.bbxyard.mp.logic.mapper")
public class LogicApplication {

    public static void main(String[] args) {
        SpringApplication.run(LogicApplication.class, args);
    }

}

