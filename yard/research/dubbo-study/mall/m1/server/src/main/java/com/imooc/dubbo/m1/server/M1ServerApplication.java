package com.imooc.dubbo.m1.server;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ImportResource;


@SpringBootApplication
@ImportResource(value = {"classpath:spring/spring-jdbc.xml","classpath:spring/spring-dubbo.xml"})
public class M1ServerApplication {
    public static void main(String[] args) {
        SpringApplication.run(M1ServerApplication.class, args);
    }
}
