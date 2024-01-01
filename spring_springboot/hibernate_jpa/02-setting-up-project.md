### Automatic Data Source Configuration

* In Spring Boot , Hibernate is the default implementation of JPA.
* EntityManager is main component for creating queries etc...
* EntityManager is from Jakarta Persistence API (JPA).
* Based on configs, Spring Boot will automatically create the beans:
    * DataSource , EntityManager, ...
* You can then inject these into your app, for example your DAO.


### Setting up Project with Spring Initializer;

* At Spring Initializer website, start.spring.io
* Add dependencies
    * MySQL Driver: mysql-connecter-j
    * Spring Data JPA: spring-boot-starter-data-jpa
* Spring boot will automatically read the database properties from the application.properties file.

#### application.properties
* spring.datasource.url=jdbc:mysql://localhost:3306/student_tracker
* spring.datasource.username=springstudent
* spring.datasource.password=springstudent