## Input, Output & Pipes in Linux

Every command has: Input and output

- Output from one program can become the input of another command.

- Syntax of passing output as parameter
- "pipe" command : |
- Pipes the output of the previous command as an input to the next command

- Ex => cat /var/log/syslog | less

* less command
  - Displays the contents of a file or a command output, one page at a time.
  - Allows you to navigate forward and backward through the file.
  - Mostly used for opening large files, as less doesn't read the entire file, which results in faster load times.

#### pipe and less

- ls /usr/bin | less (This will give us page by page view of directory content)

#### Filtering

- using grep
- Stands for: Globally Search for Regular Expression and Print out
- Searches for a particular pattern of characters and displays all lines that contain that pattern.
- ls /usr/bin | grep python (get all the files or folder which contain python in their name)
- history | grep sudo (get all the command that contain sudo , that we previously executed)

#### Redirect in Linux

- Redirect results of a command into a file.
- We can do that by

  - \> character is the redirect operator
  - Takes the output from the previous command and sends it to a file that you give.

- Ex - history | grep sudo > sudo-commands.txt (This will store the result of grep in sudo-commands.txt file)

- we can also redirect content of one file into another file.
- Ex - cat sudo-commands.txt > sudo-backup.txt

#### Standard I/O

Standard Input and Standard Output

- Every program has 3 built-in streams:
  - STDIN(0) = standard input
  - STDOUT(1) = Standard Output
  - STDERR (2) = Standard Error

* We pipe or redirect the standard output from one command to the standard output of another command.
