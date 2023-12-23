## We can create custom properties in applicaton.properties file

we can create custom properties in properties file.

```java
// we can use them with @Value("${properties name}");

@Value("${profile.name}")
public String profileName;

// This will bind profile.name to profileName;
// profile.name is a property in application.properties file.
```