## How to use actuator 

* by default only /health route is visible , other routes are hidden.
* To see difference routes , change the application.properties files

#### Insert the following in application.properties

* management.endpoints.web.exposure.include=health,info (this will make health , info routes visible)
* management.endpoints.web.exposure.include=* (this will make all the routes visible)
