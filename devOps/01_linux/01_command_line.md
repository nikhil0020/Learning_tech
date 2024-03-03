1. pwd - print working directory
2. ls - list
3. cd - change directory
4. mkdir - make directory
    * mkdir folderName
5. touch - to create a file
    * touch fileName.type
    * touch main.py
6. rm - remove a file
    * to remove a directory that contains some files and folder, then we have to use 'rm -r folderName'
    * -r stands for recursive
7. clear - clears the terminal
8. Everything in linux is a file
    - directories
    - text documents , Pictures etc.
    - Commands like pwd, ls etc.
    - devices , like Printer, Keyboard, USB , CD
    - you can move them like files , show content, copy them, etc.


#### Navigating to file system

1. cd /usr/local/bin
2. mv [filename] [new_filename] = rename the file to a new filename
3. cp - to copy a file or directory
    * cp -r [dirname] [new_dirname] = copy dirname to new_dirname

#### Some useful commands

1. I want to see content inside a directory and it's sub directories.
    * ls -R [dirName] 
    * R is recursive flag
2. history
    * gives a list of all past commands typed in the current terminal session.
3. Ctrl + r = search in history
4. Ctrl + c = stop current command
5. Ctrl + shift + v = to paste command in terminal
6. ls -a = to show hidden files
7. cat = concatination (display file content)

#### Executing command as a superuser

1. sudo = allow regular users to run programs with the security privileges of the superuser or root.
2. sudo adduser username = to add new user
3. sudo addgroup groupname = to add group
4. su - username = to switch user