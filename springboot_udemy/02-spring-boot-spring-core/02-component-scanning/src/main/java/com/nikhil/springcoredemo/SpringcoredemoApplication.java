package com.nikhil.springcoredemo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

//@SpringBootApplication(scanBasePackages = {"com.nikhil.springcoredemo","com.nikhil.util"})
@SpringBootApplication
public class SpringcoredemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(SpringcoredemoApplication.class, args);
	}

}
