package com.nikhil.springboot.cruddemo.dao;

import com.nikhil.springboot.cruddemo.entity.Employee;
import jakarta.persistence.EntityManager;
import jakarta.persistence.TypedQuery;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public class EmployeeDAOJpaImpl implements EmployeeDAO{

    private EntityManager entityManager;

    @Autowired
    public EmployeeDAOJpaImpl(EntityManager theentityManager) {
        entityManager = theentityManager;
    }

    @Override
    public List<Employee> findAll() {
        TypedQuery<Employee> theQuery = entityManager.createQuery("from Employee",Employee.class);
        return theQuery.getResultList();
    }

    @Override
    public Employee findById(int id) {
        return entityManager.find(Employee.class,id);
    }

    @Override
    public Employee save(Employee employee) {
        return entityManager.merge(employee);
    }

    @Override
    public void deleteById(int id) {
        Employee employee = entityManager.find(Employee.class,id);
        entityManager.remove(employee);
    }
}
