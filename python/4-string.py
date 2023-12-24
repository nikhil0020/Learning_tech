
# ? String Basic


# * String can be repressented in single and double quotes .
# * Multiline string can be assigned using triple single and double quotes


a = """Lorem ipsum dolor sit amet,
consectetur adipiscing elit,
sed do eiusmod tempor incididunt
ut labore et dolore magna aliqua."""

print(a)

# * Strings are Arrays
# We can access any character of a string just like arrays

# print(a[3])

# * Looping through a String
for x in "banana":
    print(x)

# * Length of a string can be find by len() function

# * check string
txt = "The best things in life are free!"
print("free" in txt)

# * Check if NOT

print("life" not in txt)

# We can use these checks in if statements


# ? String Slicing

# we can do slicing in a string

b = "Hello World!"
print(b[2:5]) # This will print characters from position 2 to position 5(not including)

# Slicing from start

print(b[:5])

# Slicing to the End
print(b[2:])

# Negative Indexing
print(b[-5:-2])


# ? Modify string

# * Upper case
 
a = "Hello , World!"
print(a.upper())

# * Lower Case

print(a.lower())

# * Remove WhiteSpace

x = "    Jai Shree Ram   "
print(x.strip())

# * Replace string

print(a.replace("l","o")) # ? It will replace all the character 'l' to 'o';

# * Split

print(a.split(",")) # we can provide the seperator in function



# ? Concatenation

# we can use + operator to concat string

# ? Format Strings

age = 36
# txt = "My name is John, I am " + age
# print(txt)
# ! We cannot combine strings and number like this

# * We have to use format() method to combine numbers and strings

txt = "My name is Nikhil , and I am {}";

print(txt.format(10))

# * format method can take unlimited number of arguments

# We can use index numbers to be sure the argument are placed in correct order
quantity = 3
itemno = 567
price = 49.95
myorder = "I want to pay {2} dollars for {0} pieces of item {1}."
print(myorder.format(quantity, itemno, price))

# * Escape character

txt = "We are the so-called \"Vikings\" from the north."

# * There are various string methods present