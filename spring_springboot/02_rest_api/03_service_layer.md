### Purpose for service layer

- Service Facade design pattern.
- Intermediate layer for custom business logic.
- Integrate data from multiple sources (DAO / repositories).

![Service Layer](../assets/service_layer_1.png)

#### Spring provide @Service annotation

### Development process for Employee service layer

1. Define service interface.
2. Define service implementation.

- Inject the EmployeeDAO.

![Step 1 : Define service interface](../assets/step1_serviceLayer.png)

![Step 2 : Define service implementation](../assets/step2_serviceLayer.png)

![Service Layer Best Practices](../assets/serviceLayer_bestPractice.png)
