
#* Text Type:	str
#* Numeric Types:	int, float, complex
#* Sequence Types:	list, tuple, range
#* Mapping Type:	dict
#* Set Types:	set, frozenset
#* Boolean Type:	bool
#* Binary Types:	bytes, bytearray, memoryview
#* None Type:	NoneType


# str 
x = "Hello World"

# int 
x = 20

# float
x = 20.5

# double
x = 1j

# list
x = ["apple" , "banana" , "cherry"]

# tuple
x = ("apple", "banana" , "cherry")
res = x.count("apple")
print(res)

# dict
x = {
    "name" : "Nikhil",
    "age" : 36
}

# set
x = {"apple" , "apple" , "banana" , "cherry"}
# Printing order can be different
print(x)

# frozenset
x = frozenset({"apple" , "banana" , "cherry"})

# bool
x = True | False

# bytes
x = b"Hello"

# bytearray
x = bytearray(5)

# memoryview
x = memoryview(bytes(5))

# NoneType
x = None