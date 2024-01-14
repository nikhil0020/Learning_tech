### Development process

1. Create a custom error response class.
2. Create a custom exception class.
3. Update REST service to throw exception if student not found.
4. Add an exception handler method using @ExceptionHandler


#### STEP 1: Create a custom error response class.

* The custom error response class will be sent back to client as JSON.
* We will define as Java class (POJO)
* Jackson will handle converting it to JSON.

![Error Response](../assets/error_response.png)
