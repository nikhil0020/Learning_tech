#### Step 1: Define a methond in DAO interface

```java
import com.nikhil.cruddemo.entity.Student;

public interface StudentDAO{
    void save(Student theStudemt);
    Student findById(Integer id);
}
```


#### Step 2: Override this method in DAO implementation


```java
import com.nikhil.cruddemo.entity.Student;
import jakarta.persistence.EntityManager;

@Repository
public class StudentDaoImpl implements StudentDAO{

    private EntityManager entityManager;

    @AutoWired
    public StudentDAOImpl(EntityManager theEntityManager){
        entityManager = theEntityManager;
    }

    // ...
    // other methods
    // ...
    
    @Override
    public Student findById(Integer id){
      return entityManager.find(Student.call,id);
    }
}
```

#### Step 3: use it in you application.

```java
public class CruddemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(CruddemoApplication.class, args);
	}

	@Bean
	public CommandLineRunner commandLineRunner(StudentDAO studentDAO){
		return runner -> {
			readStudent(studentDAO);
		};
	}

	private void readStudent(StudentDAO studentDAO){
		// create the student object
		System.out.println("Creating a new student object ... ");
		Student tempStudent = new Student("Nikhil","Singh","code.nikhil20@gmail.com");
		// save the student object

		System.out.println("Saving the student ... ");
		studentDAO.save(tempStudent);

		// display id of the saved student
		System.out.println("Saved Student , Generate ID : " + tempStudent.getId());
	  
    // find the student in database.
    
    Student myStudent = studentDAO.findById(tempStudent.getId());

    System.out.println("Found a student" + myStudent);

    // we have defined toString method that will be called here
  }
}
```


### JPA Query Language (JPQL)

1. Query language for retrieving objects.
2. Similar ini concept to SQL
  1. Where, like , order by , join , in , etc.
3. However, JPQL is based on **entity name** and **entity fields** (IMP)


```java
TypedQuery<Student> theQuery = entityManager.createQuery("FROM Student", Student.class);
<List> Student = theQuery.getResultList();
// Student is JPA Entity class name
```

#### Note : Student is NOT the name of the database table , All JPQL syntax is based on entity name and entity fields.
```java

TypedQuery<Student> theQuery = entityManager.createQuery("FROM Student WHERE lastName='Doe'", Student.class);

List<Student> students = theQuery.getResultList();

// here lastName is field of JPA Entity

```

### JPQL - Named Parameters

```java
public List<Student> findByLastName(String theLastName) {
  TypedQuery<Student> theQuery = entityManager.createQuery("FROM Student WHERE lastName=:theData", Student.class);
  theQuery.setParameter("theData",theLastName);
  return theQuery.getResultList();
}
// :theData -> we use : notation for dynamic query.
```


## Development process

1. Add new method to DAO interface.
2. Add new method to DAO implementation.
3. Update main app.


#### Step 1: Add new method to DAO interface.

```java
import com.nikhil.cruddemo.entity.Student;

public interface StudentDAO{
    void save(Student theStudemt);
    Student findById(Integer id);
    List<Student> findAll();
}
```


#### Step 2: Add new method to DAO implementation



```java
import com.nikhil.cruddemo.entity.Student;
import jakarta.persistence.EntityManager;

@Repository
public class StudentDaoImpl implements StudentDAO{

    private EntityManager entityManager;

    @AutoWired
    public StudentDAOImpl(EntityManager theEntityManager){
        entityManager = theEntityManager;
    }

    // ...
    // other methods
    // ...

    @Override
    public List<Student> findAll(){
      TypedQuery<Student> theQuery = entityManager.createQuery("FROM Student", Student.class);
      return theQuery.getResultList();
    }
}
```


#### Step 3: Update main app.


```java
public class CruddemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(CruddemoApplication.class, args);
	}

	@Bean
	public CommandLineRunner commandLineRunner(StudentDAO studentDAO){
		return runner -> {
		  queryForStudent(studentDAO);
		};
	}

  private void queryForStudent(Student studentDAO){
    // get list of students
    List<Student> theStudents = studentDAO.findAll();

    // display list of students

    for (Student tempStudent : theStudents){
      System.out.println(tempStudent);
    }
  }
}
```

### Update

```java
Student theStudent = entityManager.find(Student.class,1);

// Change first name to "Scooby"
theStudent.setFirstName("Ram");

entityManager.merge(theStudent); // This will update the entity
```

#### Update last name for all students

```java
int numRowsUpdated = entityManager.createQuery("UPDATE Student SET lastName='Testter'").executeUpdate();
```
