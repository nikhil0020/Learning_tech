## Configuring bean with java code

### Development process

1. Create @Configuration class
2. Define @Bean method to configure the bean
3. Inject the bean into our controller


```java
public class SwimCoach implements Coach{
    @Override
    public String getDailyWorkout() {
        return "Swim 1000 meters as a warm up";
    }
}

// Here we don't use @Component class
```

Step 1: Create a Java class and annotate as @Configuration

```java
package com.nikhil.springdemo.config;

import org.springframework.context.annotation.Configuration;

@Configuration
public class SportConfig{

}
```

Step 2: Define @Bean method to configure the bean

```java
package com.nikhil.springdemo.config;

import org.springframework.context.annotation.Configuration;
import com.nikhil.springcoredemo.common.Coach;
import com.nikhil.springcoredemo.common.SwimCoach;
import org.springframework.context.annotation.Beans;

@Configuration
public class SportConfig{

    @Bean
    public Coach swimCoach(){
        return new SwimCoach();
    }
    // The bean id defaults to the method name
}
```

Step 3: Inject the into our Controller

```java
@RestController
public class DemoController {
    private Coach myCoach;

    @Autowired
    public DemoController( @Qualifier("swimCoach") Coach theCoach){
        myCoach = theCoach;
    }

    // Inject the bean using the bean id
}
```

### Use case for @Bean

* Make an existing third-party class available to Spring framework;
* You may not have access to the source code of third-party class.
* However , you would like to use the third-party class as a Spring bean.


### Custom bean id

```java
@Bean("aquatic")
public Coach swimCoach(){
    return new SwimCoach();
} 

// Here bean id is aquatic not swimCoach;
```