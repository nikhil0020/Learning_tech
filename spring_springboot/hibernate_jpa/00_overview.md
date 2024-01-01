### What is Hibernate?

* A framework for persisting / saving Java objects in a database.
    * www.hibernate.org/orm

### Benifits of Hibernate

* Hibernate handles all of the low-level SQL.
* Minimizes the amount of JDBC code you can have to develop.
* Hibernate provides the Object-to-Relational Mapping(ORM)

### ORM

* The developer defines mapping between Java class and database table.


### What is JPA? 

* Jakarta Persistence API (JPA)... previously known as Java Persistence API
    * Standard API for ORM

* Only a specification
    * Defines a set of interfaces
    * Requires an implementation to be usable

#### What are the benifits of JPA

* By having a standard API, you are not locked to vendor's implementation.

* Maintain portable, flexible code by coding to JPA spec (interfaces)

* Can theoretically which vendor implementations
    * For example, if Vendor ABC stops supporting their product.
    * You can switch to vendor XYZ without vendor lock in


#### Hibernate / JPA and JDBC 

* Hibernate / JPA uses JDBC for all database communications.

