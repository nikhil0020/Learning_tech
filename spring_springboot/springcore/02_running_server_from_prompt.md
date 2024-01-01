## Running from the command line

Two options for running app

1. Use **java -jar**
    *  java -jar mycoopapp.jar
2. Use Spring Boot Maven plugin
    * mvnw spring-boot:run 


* use **./mvnw package** for creating jar file from root project folder.
    1. then run **java -jar <appName.jar file>** in target folder.
    2. we can run **./mvnw spring-boot:run** in root project folder

So these are the two ways for running spring-boot application from the command line.