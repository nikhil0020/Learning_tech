package com.springdemo.mycoolapp.rest;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class MyRestContoller {

    @Value("${profile.name}")
    private String name;

    @GetMapping("/")
    public String sayHello(){
        return "Hello World";
    }

    @GetMapping("/profile")
    public String getProfile(){
        return "My name is " + name;
    }
}
