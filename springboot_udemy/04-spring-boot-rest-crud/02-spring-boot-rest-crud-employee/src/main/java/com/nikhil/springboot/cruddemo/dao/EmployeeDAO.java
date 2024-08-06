package com.nikhil.springboot.cruddemo.dao;

import com.nikhil.springboot.cruddemo.entity.Employee;
import java.util.List;
public interface EmployeeDAO {
    List<Employee> findAll();
    Employee findById(int id);
    Employee save(Employee employee); // will be used for add and update employee
    // when id == 0 , add the employee else update it.
    void deleteById(int id);
}
