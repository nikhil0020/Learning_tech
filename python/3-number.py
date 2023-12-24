# * we cannot convert complex number to any number type

# interger or int can be of unlimited length without decimals
# float contain one or more decimals
# float can also be scientific with an "e" to indicate the power of 10.
x = 35e3
y = 12E4
z = -87.7e-10

print(z)


# complex number

a = 3 + 5j
print(a)

# Python does not have a random() function to make a random number, 
# but Python has a built-in module called random that can be used to make random numbers:

import random

print(random.randrange(1,20))