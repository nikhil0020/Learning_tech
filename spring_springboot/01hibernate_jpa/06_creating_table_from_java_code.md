### It is not recommended to use this in production.

### Best way to handle creation of table in database is through sql scripts by DBAs.


1. JPA/Hibernate provides an option to automatically create database tables.
2. Creates tables based on Java code with JPA/Hibernate annotations.
3. Useful for development and testing.


### Configuration

* In Spring Boot Configuration file: application.properties
  spring.jpa.hibernate.ddl-auto=create
* When you run your app, JPA / Hibernate will drop table then create them.
* Based on the JPA / Hibernate annotations in your Java code.

| Properties value   | Property Description    |
|--------------- | --------------- |
| none   | No action will be performed   |
| create-only   | Database tables are only created   |
| drop   | Database tables are dropped   |
| create   | Database tables are dropped followed by database tables creation  |
| create-drop | Database tables are dropped followed by database tables creation. On application shutdown, drop the database tables |
| update | update the database tables schema |
| validate | Validate the database tables schema |

