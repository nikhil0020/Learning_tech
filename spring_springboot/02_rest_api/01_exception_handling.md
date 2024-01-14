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


#### STEP 2: Create a custom exception class

* The Custom student exception will used by our REST service.
* In our code, if we can't find student, then we'll throw an exception.
* Need to define a custom student exception class.
  * StudentNotFoundException

![Custom Exception](../assets/custom_exception.png)

#### STEP 3: Update rest service to throw exception

![Rest Exception](../assets/update_rest_exception.png)

#### STEP 4: Add exception handler method

* Define exception handler method(s) with @ExceptionHandler annotation.
* Exception handler will return a ResponseEntity.
* ResponseEntity is a wrapper for HTTP response object.
* ResponseEntity provides fine-grained control to specify:
  * HTTP status code, HTTP headers and Response Body.

![ExceptionHandler method 1](../assets/exceptionHandlerMethod1.png)

![ExceptionHandler method 2](../assets/exceptionHandlerMethod2.png)

![ExceptionHandler method 3](../assets/exceptionHandlerMethod3.png)

