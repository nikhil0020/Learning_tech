## Java JSON Data Binding


* Data binding is the process of converting JSON data to a Java POJO and vice versa.
* Also known as Mapping , Serialization/Deserialization , Marshalling/Unmarshalling.

### JSON Data Binding with Jackson

* Spring uses the **Jackson Project** behind the scenes.
* Jackson handles data binding between JSON and Java POJO.
* Springboot starter web automatically includes dependency for Jackson.
* POJO -> Plain Old Java Object
### Jackson Data binding

* By Default, Jackson wwill call appropriate getter/setter method.


### Development process

1. Create Java POJO class for Student.
2. Create Spring REST Service using @RestController.


## Path variables

* Retrieve a single student by id.

* GET -> /api/students/{studentId}  (Retrieve a single student)

```java
@GetMapping("/students/{studentId}")
public Student getStudent(@PathVariable int studentId){
  // by default variables should match
    return theStudents.get(studentId);
}
```



