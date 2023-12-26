* By default, when your application starts, all beans are initialized.
    * @Component, etc ...

* Spring will create an instance of each and make them available.


### Lazy Initialization

* Instead of creating all beans up front, we can specify lazy initialization
* A bean will only be initialized in the following cases:
    * It is needed for dependency injection.
    * Or it is explicitly requested
* Add the @Lazy annotation to a given class.
* To add @Lazy to all the classes turns into tedious work for a large number of classes
* we could set a global configuration property

```java
package com.nikhil.springcoredemo.common;

import org.springframework.stereotype.Component;

@Component
@Lazy
public class TennisCoach implements Coach{

    @Override
    public String getDailyWorkout(){
        return "Play Tennis";
    }
}

// here we mention TennisCoach as primary Coach
```

### Lazy Initialization - Global configuration

```java
// File: application.properties
spring.main.lazy-initialization=true

// All beans are lazy , no beans are created until needed.
// Inclding our DemoController
```


Advantages
* Only create object as sneeded.
* May help with faster startup time if you have large number of components.

Disadvantages
* If you have web related components like @RestController, not created until requested
* May not discover configuration issues until too late
* Need to make sure you have enough memory for all beans once created.


Lazy initialization feature is disabled by default