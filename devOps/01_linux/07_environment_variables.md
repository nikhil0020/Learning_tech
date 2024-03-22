# Environment variables

#### printenv command is used to get all env

* Ex -> printenv (prints all env)
* Ex -> printenv | grep USER (using pipes and global search , we can get all the env that contain user)


### For application

* use case : Sensitive data for application
    * Set these data as env vars on server
        * Ex -> DB_PWD=secretpassword
    * Apps can read those env vars
    * By creating these env vars, we make them available in the environment

* How to access these env variables ??
    * Every programming language allows you to access environment variables

* Use case : Make application more flexible
    * Suppose we have different server
        * DEV server
        * TESTING server
        * PRODUCTION server
    * We can set different env variables for more flexibility

* How to set env variables??
    * export DB_USERNAME=dbuser (in terminal or shell script)

* Delete env variables??
    * unset DB_USERNAME

* Persisting env variables
    * When we set env using export command , they are available only in that particular session or terminal
    * How do we set them permanently??
        * we have to add them to shell specific configuration file
        * Ex -> .bash_profile , .bashrc etc , that run when we create new session , whenever a bash login shell is entered

* Persisting env variables (System wide)
    * .bashrc is user specific
    * a file located in /etc folder called "environment"
    * vi /etc/environment

### ADD a custom command/program

* append the path of the command or program to PATH variable in .bashrc
* run it from anywhere