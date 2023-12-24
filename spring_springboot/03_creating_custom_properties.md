## We can create custom properties in applicaton.properties file

we can create custom properties in properties file.

```java
// we can use them with @Value("${properties name}");

@Value("${profile.name}")
public String profileName;

// This will bind profile.name to profileName;
// profile.name is a property in application.properties file.
```

**Properties are roughly grouped into following categories**

1. Core
2. Web 
3. Security
4. Data
5. Actuator
6. Integration
7. DevTools
8. Testing

```java
// Log levels security mapping
// Setting log levels based on package names
logging.level.org.springframework=DEBUG
logging.level.org.hibernate=TRACE
logging.level.com.mycoolapp=INFO

// TRACE, DEBUG , INFO , WARN , ERROR, FATAL , OFF

// log into a file

logging.file.name=my-logging-file.log
logging.file.path=/myapps/demo

// HTTP server port

server.port=7070

// Context path of the application (by default is '/')
server.servlet.context-path=/hello
// Ex - http://localhost:7070/hello/

// session time out (Default is 30min)
server.servlet.session.timeout=15m // set timeout to 15 min

//Setting base path for actuator
// management.endpoints.web.base-path=/actutor (By default)

management.endpoints.web.base-path=/myNewPath

// Security properites
// Default user name
spring.security.user.name=admin

// Password for default user
spring.security.user.password=topsecret

```

