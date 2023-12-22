## How to use actuator 

* by default only /health route is visible , other routes are hidden.
* To see different routes , change the application.properties files

#### Insert the following in application.properties

```java
management.endpoints.web.exposure.include=health,info //(this will make health , info routes visible)
// management.endpoints.web.exposure.include=* //(this will make all the routes visible)
management.info.env.enabled=true
```

```java
// After exposing the info end point we will get access to /info route but there will be no information

// To get some information we have make some changes to our application.properties file

info.app.name = Multiverse
info.app.description = Welcome to the multiverse
info.app.version = 1.0.0
```


