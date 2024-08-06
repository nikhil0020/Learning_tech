## What are build and package manager tools

* app needs to be deployed on a production server.
* How to move code and it's dependencies??
* Package application into single movable file
    * that single packaged files is called application artifact.
    * packaging = "building the code"
* Building the code:
    * compiling
    * compress
    * Hundreds of file into 1 single file.
    * after building we keep artifact in storage
        * to deploy it multiple times, have backup etc.
        * deploy the artifact to dev , test and prod.
    * Storage we use to store artifact is called artifact repository.
        * Ex of such storage are : Nexus , Jfrog etc.

### What kind of file is the artifact?

* Artifact file looks different for each programming language.
    * Ex : For Java application 
        * JAR or WAR file are artifacts ( JAR => java archive)
        * include whole code plus dependencies
            * Spring framework
            * datetime libraries
            * pdf processing libraries.


## Build An artifact

### Build using maven

* In maven we have to add configuration in pom.xml related to build 