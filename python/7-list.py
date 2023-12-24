
# * List

# Lists are used to store multiple items in a single variable.
# Lists are one of 4 built-in data types in Python used to store collections of data, 
# the other 3 are Tuple, Set, and Dictionary, all with different qualities and usage.


fruits = ["Apple" , "banana" , "cherry"]

# print(fruits)

# We can find length of a list by len() method

# List items can be of any data type.
# One list can contain multiple data types.

# * Creating List Using constructor

gods = list(("Ram" , "Hanuman" , "Krishna"))

# print(gods)

# * Accessing list items

# list items can be accesssed by square notation

# first index is 0
# We can also access by negative index , -1 for last element


# * Range of indexes

thislist = ["apple", "banana", "cherry", "orange", "kiwi", "melon", "mango"]
# print(thislist[2:5])

# * From begining

# print(thislist[:4])


# * Till end
# print(thislist[2:])

# Negative indexes

print(thislist[-4:-1])
# -4 is included and -1 is not included

# * Check if item exist or not

if "apple" in thislist:
    print("present")


# * Change a range of item values

thislist[1:3] = ["blackcurrent" , "watermelon"]

# print(thislist)

# We can also replace more or less item in a list and other items will shift accordingly

thislist[1:3] = ["guava"]

# print(thislist)

thislist[2:3] = ["kiwi" , "mango" , "grapes"]
# print(thislist)

# * Add items to list

# Inserting using insert method at a specific location

thislist.insert(2,"peach")
# print(thislist)


# Append Items

thislist.append("berry")
# print(thislist)

# Extend - To append elements from another list to current list, use extent() method

thislist.extend(fruits)

# print(thislist)

# * IMP --- we can add any iterable object (tuples , sets , dictionary)


# ? Remove list items

# * Remove specifc item

thislist.remove("kiwi")
# print(thislist) # this will remove the first occurance of 'kiwi'

# * Remove specific index

thislist.pop(1) # It will remove element at index 1 , by default remove the element at last index
# print(thislist) 

thislist.pop()
# print(thislist) 


# * del vs clear

newlist = ["apple" , "mango"]

del newlist

# print(newlist) #! This line will give error because list is completely deleted

anothernewlist = ["apple" , "mango"]

anothernewlist.clear()

# print(anothernewlist)
# clear will empty the list by list will not delete completely


# ? LOOPING IN A LIST

# for fruit in thislist:
#     print(fruit)

# * Loop through index

# for i in range(len(thislist)):
#     print(thislist[i])

# * Using a while loop

# i = 0
# while i < len(thislist):
#     print(thislist[i])
#     i += 1


# ** Looping using list comprehension

# Short hand for printing all elements

# [print(x) for x in thislist]

# The Syntax

# * newlist = [expression for item in iterable if condition == True]


newlist = [x for x in thislist if x != "kiwi"]

# print(newlist)

newlist = [x for x in range(10)]
newlist = [x for x in range(10) if x < 5]
newlist = [x if x != "banana" else "orange" for x in fruits]


# ? SORT LISTS

# List objects have a sort() method that will sort the list alphanumerically, ascending, by default:

thislist.sort()

# print(thislist)

# Reverse sort
thislist.sort(reverse=True)
# print(thislist)

# Customize sort function

nums = [100,50,65,82,23]

def myfunc(n):
    return abs(n-50)

nums.sort(key=myfunc)

print(nums)


# * Case insensitive sort

fruits = ["banana", "Orange", "Kiwi", "cherry"]
fruits.sort(key = str.lower)
# print(fruits)

# * Reverse a list

fruits = ["banana", "Orange", "Kiwi", "cherry"]
fruits.reverse()

# print(fruits)


# ? Copy Lists

# You cannot copy a list simply by typing list2 = list1, because: list2 will only be a reference to list1, 
# and changes made in list1 will automatically also be made in list2.

 
#  There are ways to make a copy, one way is to use the built-in List method copy().

newfruits = fruits.copy()
print(newfruits)

# Another way to make a copy is to use the built-in method list().

newfruits2 = list(fruits)

print(newfruits2)


# ? JOIN 2 list

# There are 3 methods
#   1 -> + operator
#   2 -> extend method
#   3 -> append every element of one list to other





