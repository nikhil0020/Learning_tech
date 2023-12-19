# Set

* Set is a child interface of Collection.
* Set interface doesn't contain any new methods. so we have to use only Collection interface methods.

## HashSet

* Based on Hash table
* Duplicates are not allowed
* Insertion order not preserved
* Heterogeneous objects allowed
* Null allowed
* Serializable and cloneable
* Best for Search Operations
* add() method wiwll return false if element already present in HashSet.


### Constructors

```java
HashSet h = new HashSet();
// Default initial capacity is 16
// Fill Ratio is 0.75 (Load Factor, if it reach max capacity, it tells with how much capacity hashset increase)
HashSet h = new HashSet(int initialCapacity);
HashSet h = new HashSet(int initialCapacity,float fillRatio);
HashSet h = new HashSet(Collection c);
```

## LinkedHashSet

* When we want to preserve the insertion order.
* Underlying data structure is Hash table + linked list

## SortedSet( Interface )

### Methods
```java
Object first();
Object last();
SortedSet headSet(Object obj); // return the SortedSet whose elements are < obj
SortedSet tailSet(Object obj); // return the SortedSet whose elements are >= obj
SortedSet subSet(Object obj1, Object obj2); // return the SortedSet whose are >= obj1 & < obj2
Comparator comparator(); // return Comparator object that describes underlying sorting technique, if we are using default natural sorting order then we will get null
```

## TreeSet

* Underlying data structure is balanced tree.
* Duplicates are not allowed
* Insertion order not allowed
* Sorting order
* Heterogeneous Objects are not allowed, will through exception
* Null acceptance is there but only once

### Constructors

```java
TreeSet t = new TreeSet(); // Element will be inserted according to default natural sorting order.
TreeSet t = new TreeSet(Comparator c); // when we want customized sorting
TreeSet t = new TreeSet(Collection c);
TreeSet t = new TreeSet(SortedSet s);
```

