### Scanning for Component Classes

* Spring will scan your java classes for special annotations.
    * @Component, etc ...

* Automatically register the beans in the Spring container.

* By Default, Spring Boot starts component scanning
    * From same package as your main Spring Boot application.
    * Also scans sub-packages recursively

* This implicitly defines a base search package
    * Allows you to leverage default component scanning.
    * No need to explicitly reference the base package name.

* Default scanning is fine if everything is under
    * com.nikhil.springcoredemo

* But what about my other packages?
    * com.nikhil.util
    * org.acme.cart
    * edu.cmu.srs
    ```java
    // Explicitly list base packages to scan
    package com.nikhil.springcoredemo;
    ...
    @SpringBootApplication(
            scanBasePackages={"com.nikhil.springcoredemo",
                "com.nikhil.util",
                "org.acme.cart",
                "edu.cmu.srs"
            }
    )

    public class SpringcoredemoApplication{
        ...
    }
    ```



## Annotations

* **SpringBootApplication** is composed of the following annotations.

|**Annotation**|**Description**|
|:-------------|:--------------|
|@EnableAutoConfiguration| Enables Spring Boot's auto-configuration support|
| @ComponentScan | Enables component scanning of current package. Also recursively scans sub-packages|
| @Configuration | Able to register extra beans with @Bean or import other configuration classes|


