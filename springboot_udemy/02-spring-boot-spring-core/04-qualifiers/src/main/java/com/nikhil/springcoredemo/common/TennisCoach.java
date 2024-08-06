package com.nikhil.springcoredemo.common;

import org.springframework.stereotype.Component;

@Component
public class TennisCoach implements Coach{

    @Override
    public String getDailyWorkout(){
        return "Play Tennis";
    }
}
