### Methods in collection

1. `add(Object o)` -> add object  o from collection
2. `addAll(Collection C)` -> add a group of object in collection.
3. `remove(Object o)` -> remove a object o from collection.
4. `removeAll(Collection C)` -> remove a group of object from collection.
5. `clear()` -> This will remove all the element from the collection.
6. `retailAll(Collection C)` -> Except a group of object remove all the element from the collection.
7. `isEmpty()` -> return true or false
8. `size()` -> return the size of collection
9. `contains(Object o)` -> check whether a object is present or not.
10. `containsAll(Collection C)` -> check whether a group of collection is available or not.
11. `toArray()` -> convert object to Array
    * Ex -> Object[] a = c.toArray()
    * return type of toArray() method is Object[]
12. `iterator()` -> The iterator() method can be used to get an iterator for any collection .
    * Looping through a collection , use the `hasNext()` and `next()` methods of the iterator.
    * Iterator iterator()


## **List (I) :** 
    All the methods present in Collection Interface are available in List Interface.

Other method specific to List interface

1. `add(int index, Object o)`
2. `addAll(int index, Collection c)`
3. `remove(int index)`
4. `indexOf(Object o)` -> Return the first occurance of Object o
5. `lastIndexOf(Object o)` -> return last occurance of Object o
6. `get(int index)` -> get element at index
7. `set(int index, Object o)` -> set Objet o at index.
8. `listIterator()` -> return ListIterator object for iteration


##

    Everywhere heterogeneous objects are allowed , except TreeSet and TreeMap
    (because these classes use SortedSet and SortedMap respectively , so Object type need to be same)


## ArrayList
1. Resizeable list
2. Duplicate allowed
3. Insertion order preserved
4. Heterogeneous Objects allowed
5. `null` insertion allowed


### Constructor in ArrayList
1. `ArrayList al = new ArrayList()`;
    * This return reference to a empty array with default capacity 10
    * New Capacity = (old * 3/2) + 1
2. `ArrayList al = new ArrayList(int initialCapacity)`
3. `ArrayList al = new ArrayList(Collection C)`

##

    Usually we can use Collections to hold and transfer Objects from one place to another place, to provide support for this requires every Collection already implements Serializable and Clonable Interfaces.


### RandomAccess(I)
1. present in java.util
2. ArrayList and Vector implements this interface
3. It is a Marker Interface -> means it has no method



![ArrayList vs Vector](./assets/arraylist_vs_vector.png)



### How to get Synchronized array list

```java
    ArrayList L1 = new ArrayList();

    List L2 = Collection.synchronizedList(L1);

    // L1 is not synchronized list
    // L2 is synchronized list
```

    To make ArrayList synchronized we use Collection class method synchronizedList().
    
```java
public static List synchronizedList(List L);
```


## Linked List

1. Doubly Linked list
2. Insertion order preserved
3. Duplicates allowed
4. Heterogeneous objects are allowed
5. Null insertion is possible
6. LinkedList implements serializable and cloneable interface but not RandomAccess interface.
7. LinkedList is the best choice if frequent operation is insertion and deletion in the middle.
8. LinkedList is worst choice if our frequent operation is retreival operation.

### Methods available only to linked list

```java
void addFirst(Object o);
void addLast(Object o);
Object getFirst();
Object getLast();
Object removeFirst();
Object removeLast();
```

### Constructors

```java
LinkedList l = new LinkedList(); // creates ab Empty LinkedList Object
LinkedList l = new LinkedList(Collection c); // Create an equivalent LinkedList Object for the given Collection
```

### Difference between ArrayList and LinkedList

|**ArrayList** |**LinkedList** |
|:-------------|:--------------|
|Best for retrieval | Best choice for insertion and deletion in middle |
|Worst choice is insertion and deletion in middle | Worst choice for retrieval |
|Resizable and growable array | Double LL |
|Implements RandomAccess interface | Doesnot implements RandomAccess interface |



## Stack Class

1. Stack is child of vector
2. LIFO order

### Constructor

```java
Stack s = new Stack(); // Only one constructor
```

### Methods

```java
void push(Object o);
Object pop(); // return and remove top of stack
Object peek(); // return top of the stack
Boolean empty(); // to check if stack is empty or not
int search(Element e); // return offset of the element (offset means at which position from top the element is located), if element is not present then it will return -1
```