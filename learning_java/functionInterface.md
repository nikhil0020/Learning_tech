* A functtional interface is a interface that has only one method.

```java
@FunctionalInterface
interface Person{
    void show();
}
// if we declare two or more methods with the above tag (@FunctionalInterface) then this will throw a error.
// Except methods that are present in Object class , we can use them
```


### How to use functional interfaces

* We know that to declare a class we need some syntax but this syntax can be memorised by compiler

```java
// Interface Person has only one method

// How we use it without lambda expression

class Student implements Person {
    public void show(){
        System.out.println("Hello there");
    }
};

// Then we use this class Student to create a object and then use show method

Person nikhil = new Student();

nikhil.show();


// But we  can also do this using anonymous class

Person ram = new Person(){
    public void show(){
        System.out.println("Hello there from anonymous class");
    }
}

ram.show();

```

### Let's see how we can do this using lambda expression

```java
Person shyam = () -> {
    System.out.println("Hello there from lambda expression");
}
// This can be done in functional interface only where there is only one method
// We can also remove the curly brackets if there is one statement
// we can also pass the parameter if there are any
// -> this is called lambda expression

```