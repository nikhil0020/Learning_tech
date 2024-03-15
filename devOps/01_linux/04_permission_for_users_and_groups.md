## Ownership and file permission

- ls = shows files in a dir
- ls -a = shows hidden files
- ls -l = shows permission associated with a file

## Ownership vs permission

### Ownership

- Who owns the file/directory?

- Ownership levels of a file/directory
  - Which user owns the file?
  - Which group owns the file?

* Owner is the user, who created the file
* Owning group is the primary group of that user.

#### Change Ownership

- sudo chown \<username\> : \<groupname\> \<filename\> = change Ownership

- sudo chown \<username\> \<filename\> = only change user

- sudo chgrp \<groupname\> \<filename\> = change group

### Permission

- drwxrwxr-x

- File Types:

  - \- regular file
  - d directory
  - c character device file
  - l symbolic link
  - etc.

- First rwx is for Owner

  - r Read
  - w write
  - x Execute , Execute a script/executable
  - \- No permission

- Second rwx is for group

- Dash instead of the letter means, this permission is absent

- Third block belongs to others , who are not the file owner or don't belong to group owner.

#### Modifying permission

- sudo chmod -x api (take away the executable permission from api folder, this will remove permission from all user)

* How to take permission for specified
  | Owner | Group | Other |
  |---- | ---- | --- |
  | u | g | o |

  - sudo chmod g-w config.yaml (remove write permission from group)

* How to add permission
  - sudo chmod g+x config.yaml
  - sudo chmod g=rwx script.sh (gives read, write, execute access to group for script.sh)
  - sudo chmod g=r-- index.html (gives only read access to group for index.html)
  - sudo chmod a=rw- styles.css (gives read, write to everyone for styles.css)

#### setting permission using numeric values

- sudo chmod 476 index.html (it means read for user, rwx for group ,rw for others)

| value | access |
| ----- | ------ |
| 0     | ---    |
| 1     | --x    |
| 2     | -w-    |
| 3     | -wx    |
| 4     | r--    |
| 5     | r-x    |
| 6     | rw-    |
| 7     | rwx    |

### Three ways of setting permission

1. Symbolic Mode

- \+ add
- \- remove

2. Set permission

- = set the permission and override the permissions set earlier.

3. Numeric Mode

- 4 read , 2 write , 1 Execute , 0 permission

#### We can change permission of hidden files and folders just like we did above
