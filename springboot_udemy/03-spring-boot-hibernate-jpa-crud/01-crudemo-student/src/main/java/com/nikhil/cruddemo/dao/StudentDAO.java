package com.nikhil.cruddemo.dao;

import com.nikhil.cruddemo.entity.Student;

import java.util.List;

public interface StudentDAO {
    void save(Student theStudent);
    Student findById(Integer id);

    List<Student> findAll();

    void deleteStudent(Integer id);

    void deleteStudentWithLastname(String lastname);
}
