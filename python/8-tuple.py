
# * Tuples

# Tuples are used to store multiple items in a single variable.
# A tuple is a collection which is ordered and unchangeable.
# Tuples are written with round brackets.


# *Tuple Items

# Tuple items are ordered, unchangeable, and allow duplicate values.
# Tuple items are indexed, the first item has index [0], the second item has index [1] etc.

# *Ordered

# When we say that tuples are ordered, it means that the items have a defined order, and that order will not change.

# *Unchangeable

# Tuples are unchangeable, meaning that we cannot change, add or remove items after the tuple has been created.

# *Allow Duplicates

# Since tuples are indexed, they can have items with the same value:


# ? ACCESS ITEMS IN TUPLE

# -> Accessed like we did in list

# ? UPDATES IN TUPLES

# * CHANGE TUPLE VALUE

# -> create list from tuple
# -> update into that list
# -> convert that list to tuple

# * ADD ITEM TO TUPLE

# -> create list from tuple
# -> append item to list
# -> convert list to tuple


# * ADD tuple to tuple

thistuple = ("apple", "banana", "cherry")
y = ("orange",)
thistuple += y

print(thistuple)


# * Remove items from tuple

# -> convert to list
# -> delete item from list
# -> convert list to tuple


# * Delete the tuple

# -> delete the tuple using del keyword

# ? Join tuples
# 
# * Join Two tuples

# 1 -> Using + operator
# tupleC = tupleA + tupleB

# 2 -> If we want one tuple to be added multiple times
# tupleA = tupleB * 2

