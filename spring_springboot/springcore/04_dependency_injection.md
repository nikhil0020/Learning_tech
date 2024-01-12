### Dependency injection

The dependency inversion principle.

The client delegates to another object the responsibility of providing its dependencies.

### Spring Container
* Primary functions
    * Create and manage objects (Inversion of Control)
    * Inject object's dependencies (Dependency Injection)

### Injection Types

* There are multiple types of injection with Spring.

* We will cover the two recommended types of injection
    * Constructor Injection
    * Setter Injection

* Which one to use?
    * Constructor Injection
        * Use this when you have required dependencies
        * Generally recommended by the spring.io development team as first choice.

    * Setter Injection
        * Use this when you have optional dependencies
        * If dependency is not provided, your app can provide reasonable default logic.

### What is Spring AutoWiring

* For dependency injection, Spring can use autowiring.

* Spring will look for a class that matches
    * matches by type: class or interface

* Spring will inject it automatically ... hence it is autowired.

### AutoWiring Example

* Injecting a Coach implementation
* Spring will scan for @Components
* Any one implements the Coach interface??
* If so, let's inject them. For Ex - CricketCoach


## Development process for Constructor injection

* Define the dependency interface and class.
* Create demo REST Controller.
* Create a constructor in your class for injections.
* Add @GetMapping for /dailyworkout

### Step 1: Defining the dependency interface and class

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

// @Component annotation marks the class as a Spring Bean and makes it a candidate for dependency injection.
```

### @Component annotation

* @Component marks the class as a Spring Bean.
    * A Spring Bean is just a regular Java class that is managed by Spring.
* @Component also makes the bean available for dependency injection.

### Step 2: Create Demo REST Controller

```java
// File: DemoController.java

package com.nikhil.springcoredemo;

import org.springframework.web.bind.annotation.RestController;

@RestController
public class DemoController{

}
```


### Step 3: Create a constructor in your class for injections.

```java
// File: DemoController.java

package com.nikhil.springcoredemo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DemoController{
    private Coach myCoach;

    @Autowired
    public DemoController(Coach theCoach){
        myCoach = theCoach;
    }
}

// @Autowired annotation tells Spring to inject a dependency
// If you only have one constructor then @Autowired on constructor is optional.
```

### Step 4: Add @GetMapping for /dailyworkout

```java
// File: DemoController.java

package com.nikhil.springcoredemo;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class DemoController{
    private Coach myCoach;

    @Autowired
    public DemoController(Coach theCoach){
        myCoach = theCoach;
    }

    @GetMapping("/dailyworkout")
    public String getDailyWorkout(){
        return myCoach.getDailyWorkout();
    }
}
```
