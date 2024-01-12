## Setter Injection

Inject dependencies by calling setter method(s) on your class

### Development process

1. Create setter method(s) in your class for injections.
2. Configure the dependency injection with @Autowired Annotation.

The Spring framework will perform operations behind the scenes for you

```java
// File: Coach.java
package com.nikhil.springcoredemo;

public interface Coach{
    String getDailyWorkout();
}
```
```java
// File: CricketCoach.java

package com.nikhil.springcoredemo;

import org.springframework.stereotype.Component;

@Component
public class CricketCoach implements Coach{
    @Override
    public String getDailyWorkout(){
        return "Practice fast bowling for 15 minutes";
    }
}

```java
// File: DemoController.java

package com.nikhil.stringcoredemo;

import org.springframework.web.bind.annotation.RestController;

@RestController
public class DemoController{
    private Coach myCoach;

    @Autowired
    public void setCoach(Coach theCoach){
        myCoach = theCoach;
    }
}
```

### How Spring will process the application
```java
Coach theCoach = new CricketCoach();
DemoController dm = new DemoController();
dm.setCoach(theCoach);
```

Inject dependencies by calling ANY method on your class, Simply give : @Autowired


### Spring Injection Types

* Recommended by the spring.io development team.
    * Constructor Injection: required dependencies
    * Setter Injection: optional dependencies

* Not recommended by the spring.io development team
    * Field Injection (Because it makes the code harder to unit test)


### Field Injection Example

```java
@RestController
public class DemoController{
    @Autowired
    private Coach myCoach;

    @GetMapping("/dailyworkout")
    public String getDailyWorkout(){
        return myCoach.getDailyWorkout();
    }
}

// no need of constructor and setters
// we dont use it because it make unit testing harder
```