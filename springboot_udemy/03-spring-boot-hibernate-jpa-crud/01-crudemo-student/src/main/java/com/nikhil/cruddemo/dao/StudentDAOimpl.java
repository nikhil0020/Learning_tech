package com.nikhil.cruddemo.dao;

import com.nikhil.cruddemo.entity.Student;
import jakarta.persistence.EntityManager;
import jakarta.persistence.Query;
import jakarta.persistence.TypedQuery;
import jakarta.transaction.Transactional;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class StudentDAOimpl implements StudentDAO {
    // define field for entity manager
    private EntityManager entityManager;

    // inject entity manager using constructor injection

    @Autowired
    public StudentDAOimpl(EntityManager entityManager) {
        this.entityManager = entityManager;
    }

    // implement save method
    @Override
    @Transactional
    public void save(Student theStudent){
        entityManager.persist(theStudent);
    }

    // implement find method
    @Override
    public Student findById(Integer id){
      return entityManager.find(Student.class,id);
    }

    @Override
    public List<Student> findAll(){
        TypedQuery<Student> theQuery =  entityManager.createQuery("FROM Student", Student.class);
        return theQuery.getResultList();
    }

    @Override
    @Transactional
    public void deleteStudent(Integer id){
        Student myStudent = entityManager.find(Student.class,id);
        entityManager.remove(myStudent);
    }

    @Override
    @Transactional
    public void deleteStudentWithLastname(String lastName){
        Query query = entityManager.createQuery("DELETE FROM Student WHERE lastName=:lastName");
        query.setParameter("lastName",lastName);
        int numOfRowsDeleted = query.executeUpdate();
    }
}
