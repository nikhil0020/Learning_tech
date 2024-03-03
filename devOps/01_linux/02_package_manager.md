1. apt search openjdk 
    * this command will give all the versions available for a package.
2. apt install packagename = to install a package
3. apt remove packagename = to remove a package


#### Difference btw apt and apt-get

* apt is more user friendly then apt-get
* no search command is available in apt-get

#### Where are these files stored
    Repository

* Storage location, containing thousands of programs
* Package manager fetches the packages from these repositories.
* Always update the package index before upgrading or installing new packages.

#### How to install softwares that are not available on repository

1. Use ubuntu-software-center
2. Use snap package manager
3. apt-add-repository (add repository to /etc/apt/sources.list)