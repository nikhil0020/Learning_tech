# Indentation matters a lot in python

# In Python, variables are created when you assign a value to it:

x = 5
y = "Hello , World!"

# There is no multiline comment in python
# but we can use string literals because python will ignore
# unassigned a string literal which is not assigned to a variable

"""
This is a comment
written in
more than just one line
"""

# Variables do not need to be declared with any particular type, and can even change type after they have been set.

# Type casting

x = str(4)  # x will be '4'
y = int(3)  # y will be 3
z = float(6)  # z will be 6.0

# We can use type function to get type of a variable

print(type(y))


# String can be declared using single and double quotes

# Variable are case sensitive

# Camel Case
myVariableName = "Sita"

# pascal case
MyVariableName = "Geeta"

# Snake Case
my_variable_name = "Ram"


# * Assign Many values to Multiple variables

a, b, c = "Hello world ", 30, False

print(a, b, c)

# * One value to muliple variables

x = y = z = "Namaste"
print(x, y, z)

# * Unpacking a collection

lectures = ["dataStructure", "networks", "OS"]

l1, l2, l3 = lectures

print(l1, l2, l3)

# ! We cannot add a string to a number

# print(a + b) #! This will give error

# * Global Variables

x = "awesome"


def myfunc():
    print("Python is " + x)


myfunc()

# we can use global keyword to make a variable global

def changeX():
  global x
  x = "fantastic"

changeX()

print("Python is " + x)
