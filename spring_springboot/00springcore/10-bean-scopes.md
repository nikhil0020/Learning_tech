### Bean scopes

* Scope refers to the lifecycle of a bean
* How long does the bean live?
* How many instances are created?
* How is the bean shared?

**Default score** : it is singleton

What is a Singleton?

* Spring Container creates only one instance of the bean, by default
* It is cached in memory
* All dependency injections for the bean
    * will reference the SAME bean

### Explicitly Specify Bean Scope

```java
import org.springframework.beans.factory.config.ConfigurableBeanFactory;
import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Component;

@Component
@Scope(ConfigurableBeanFactory.SCOPE_SINGLETON)
public class CricketCoach implements Coach {

}
```

![Scopes](./assets/bean_scopes.png)


## Bean Lifecycle

* Container Start -> 
* Bean initiated -> 
* Dependency injected -> 
* Internal Spring Processing -> 
* Your custom init method -> 
* Bean ready for use , container is shutdown -> 
* Custom Destoy method -> 
* STOP 

### Bean lifecycle methods/ hooks

* You can add custom code during bean initialization.
    * Calling custom business logic methods.
    * Setting up handles to resources (db , sockets, file etc).

* You can add custom code during bean initialization
    * Calling custom business logic method.
    * Clean up handles to resources (db, sockets, file etc).


Step 1 : Define your method for init and destroy
Step 2 : Add annotations: @PostConstruct and @PreDestroy


```java
@Component
public class CricketCoach implements Coach{
    
    public CricketCoach(){
        System.out.println("In Constructor : ", getClass.getSimpleNamwe());
    }

    // define our init method
    @PostConstruct
    public void doMyStartupStuff(){
        System.out.println("In doMyStartupStuff : ", getClass.getSimpleNamwe());
    }

    // define our destroy method
    @PreDestroy
    public void doMyCleanupStuff(){
        System.out.println("In doMyCleanupStuff : ", getClass.getSimpleNamwe());
    }


    @Override
    public String getDailyWorkout(){
        return "Practice fast bowling for 15 minutes";
    }
}
```


```java
/*
Prototype Beans and Destroy Lifecycle
There is a subtle point you need to be aware of with "prototype" scoped beans.

For "prototype" scoped beans, Spring does not call the destroy method. Gasp!

---

In contrast to the other scopes, Spring does not manage the complete lifecycle of a prototype bean: the container instantiates, configures, and otherwise assembles a prototype object, and hands it to the client, with no further record of that prototype instance.

Thus, although initialization lifecycle callback methods are called on all objects regardless of scope, in the case of prototypes, configured destruction lifecycle callbacks are not called. The client code must clean up prototype-scoped objects and release expensive resources that the prototype bean(s) are holding.



Prototype Beans and Lazy Initialization
Prototype beans are lazy by default. There is no need to use the @Lazy annotation for prototype scopes beans.
*/
```