# Three Cursors of Java

**If we want to retrieve Objects one by one from the collection then we should go for Cursors**

1. Enumeration
2. Iterator
3. ListIterator

## Enumeration

* We can use Enumeration to get Objects one by one from the old Collection Objects (Legacy Collection)
* We can create Enumeration Object by using **elements()** method of Vector Class
```java
Enumeration e = v.elements();
// v is a vector object
``` 

### Methods in Enumeration

```java
public boolean hasMoreElements();
public Object nextElement();
// The return type of nextElement method is Object, so we have to type cast that into our preferred type
```

### Problem with enumeration

1. Only applicable to legacy Collection.
2. We can only perform read operation using Enumeration , so we cannot remove or replace or update.


## Iterator

1. Iterator is univeral cursor, means it is applicable to any Collection.
2. We can perform read and remove operation with iterator.

```java
public Iterator iterator();

// Example

Iterator itr = c.iterator();

// c is any Collection Object.
```

### Methods

```java
public boolean hasNext();
public Object next();
public void remove();
```

### Limitation of iterator

1. We only move in forward direction using Iterator and Enumeration.
2. By using Iterator we can perform only read and remove operations and we cannot perform replacement of new Objects.

NOTE: To overcome above limitations of iterator we should go for ListIterator.

## LinkedList

1. We can more forwward and backward using ListIterator.
2. We can perform read, remove, replacement and addition of new objects.

```java
public ListIterator listIterator();

// Example

ListIterator itr = l.listIterator();

// Here l is any List Object(ArrayList , LinkedList , Vector or Stack)
```

### Methods

```java
public boolean hasNext();
public Object next();
public int nextIndex(); // First 3 method are for forward movement
public boolean hasPrevious();
public Object previous();
public int previousIndex(); // First 3 method are for backward movement
public void remove();
public void set(Object new);
public void add(Object new);
```

|**Properties**|**Enumeration**|**Iterator**|**ListIterator**|
|:-----|:------|:-------|:------|
|**Application for**|Only Legacy classes|Any Collection class|Only List Class|
|**Movement**|Forward|Forward|Bidirectional|
|**Accessibility**|Read|Read and remove|Read, remove, Replace, Addition of new Object|
|**How to get it?**|elements() method of vector class|iterator() method of Collection(I)|listIterator() method of List(I)|
|**Methods**|2 methods|3 methods|9 methods|
|**Is it legacy**|yes|no|no|