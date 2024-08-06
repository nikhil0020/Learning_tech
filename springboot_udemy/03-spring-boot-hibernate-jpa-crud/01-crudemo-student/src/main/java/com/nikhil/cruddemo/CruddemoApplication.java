package com.nikhil.cruddemo;

import com.nikhil.cruddemo.dao.StudentDAO;
import com.nikhil.cruddemo.entity.Student;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import java.util.List;

@SpringBootApplication
public class CruddemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(CruddemoApplication.class, args);
	}

	@Bean
	public CommandLineRunner commandLineRunner(StudentDAO studentDAO){
		return runner -> {
			// createStudent(studentDAO);
			// readStudent(studentDAO);
//			queryForStudents(studentDAO);
//			removeStudent(studentDAO);
			removeStudentWithLastname(studentDAO);
		};
	}

	private void createStudent(StudentDAO studentDAO){
		// create the student object
		System.out.println("Creating a new student object ... ");
		Student tempStudent = new Student("Nikhil","Singh","code.nikhil20@gmail.com");
		// save the student object

		System.out.println("Saving the student ... ");
		studentDAO.save(tempStudent);

		// display id of the saved student
		System.out.println("Saved Student , Generate ID : " + tempStudent.getId());
	}

  private void readStudent(StudentDAO studentDAO){
		//create the student object
		System.out.println("Creating a new student object ... ");
		Student tempStudent = new Student("Yash","Agarwal","yashubhai@gmail.com");
		// save the student object

		System.out.println("Saving the student ... ");
		studentDAO.save(tempStudent);

		// display id of the saved student
		System.out.println("Saved Student , Generate ID : " + tempStudent.getId());
    
    	Student myStudent = studentDAO.findById(tempStudent.getId());

    	System.out.println("Student = " + myStudent);
  }

  private void queryForStudents(StudentDAO studentDAO){
		List<Student> studentsList = studentDAO.findAll();

		for (Student st : studentsList){
			System.out.println(st);
		}
  }

  private void removeStudent(StudentDAO studentDAO){
		int id = 1;
		studentDAO.deleteStudent(id);
  }

  private void removeStudentWithLastname(StudentDAO studentDAO){
		String lastName = "Dev";
		studentDAO.deleteStudentWithLastname(lastName);
	}
}
