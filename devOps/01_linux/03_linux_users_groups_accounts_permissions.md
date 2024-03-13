### Types of User categories

1. Superuser Account
2. User account
3. Service Account
   - Relevant for linux server distros
   - Each service will get its own user
   - E.g. mysql user will start mysql application.
   - Best Practice for Security
   - Don't run services with root user

### Multiple Users on a server

Why having a user for each team member is important

- They need a non root user
  - Permission per team member
  - Traceability - who did what on the system?
- Admin create users

## GROUPS AND PERMISSION

### User Level

- Give Permission to user directly

### Group Level

- Group Users into Linux groups.
- give Permission to the group

### How to manage Permissions

/etc/passwd

- Stores user account information
- Everyone can read it, but only root user can change the file.

Example - ninja:x:1003:1003:Test User,,,:/home/ninja:/bin/bash
USERNAME:PASSWORD:UID:GID:GECOS:HOMEDIR:SHELL

password: x means, that encrypted password is stored in /etc/shadow

User ID (UID): Each user has a unique ID. UID 0 is reserved for root.

Group ID (GID) : The primary group ID (Stored in /etc/group file)

User ID info : Comment Field for extra information about users.

Home directory : Absolute path of user's home directory.

SHELL : Absolute path of a SHELL

#### Managing Users

- Do not edit these access control files with a text editor.
- Instead you should use the dedicated commands.

* **adduser \<username\>** = Create a new user
* This is an administrative command so we have to use sudo

  - Automatically choose policy-conformant UID and GID values
  - Automatically creates a home directory with skeletal configuration.
  - when we create a user, it also creates a group with same name as user.

* **passwd <username>** to change password of a user , and it is a system command so we have to use sudo

* su = means switch user

* **su - kram** = change the user to kram
* **su -** = change to the root user

* groupadd <groupname> = Create new group
  - sudo groupadd devops
  - By default, system assigns the next available GID from the range of group IDs specified in the login.defs file

#### Different User and Group commands

| adduser , addgroup | useradd , groupadd |
| ------------------ | ------------------ |
| Interactive        |                    |
| More user Friendly |                    |

#### how to change primary group of a user

- usermod [OPTIONS] <username> = Modify a user Account
