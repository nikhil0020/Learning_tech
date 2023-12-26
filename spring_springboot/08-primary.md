### Alternative solution to multiple implementation 

* Instead of specifying a coach by name using @Qualifier

* I simply need a coach ... I dont care which coach
    * If there are multiple coaches.
    * Then you coaches figure it out.. and tell me who's the **primary** coach


```java
package com.nikhil.springcoredemo.common;

import org.springframework.stereotype.Component;

@Component
@Primary
public class TennisCoach implements Coach{

    @Override
    public String getDailyWorkout(){
        return "Play Tennis";
    }
}

// here we mention TennisCoach as primary Coach
```

* @Primary - Only one
    * When using @Primary, can have only one for multiple implementations.
    * If you mark multiple classes with @Primary ... then wew have a problem.

### Mixing @Primary and @Qualifier

* @Qualifier has higher priority
* @Qualifier is generally recommended to use