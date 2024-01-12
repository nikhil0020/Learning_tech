## Qualifiers

When we use @Autowired , then spring will scan for @Component and check if there is some implementation or not
But if there are multiple implementation then how will spring decide which one to choose.


```java
@RestController
public class DemoController {
    private Coach myCoach;

    @Autowired
    public DemoController( @Qualifier("tennisCoach") Coach theCoach){
        myCoach = theCoach;
    }

    @GetMapping("/dailyworkout")
    public String getDailyWorkout(){
        return myCoach.getDailyWorkout();
    }
}

// We have to provide the implementation Name but first letter is small.
// @Qualifier annotation is used
```