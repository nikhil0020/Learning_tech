## Comparable Interface

* This interface present in java.lang package it contains only one method compareTo().
    * `public int compareTo(Object obj)`
* Example
    * obj1.compareTo(obj2);
        * return -ve iff obj1 has to come before obj2
        * return +ve iff obj1 has to come after obj2
        * return 0 iff obj1 and obj2 are equal

* If we are not satisfied with default natural sorting order or if the default natural sorting order is not already available then we can define our own customized sorting by using Comparator.
* Comparable meant for Default Natural sorting order where as Comparator meant for customized sorting order.


## Comparator (Interface)

* Customized sorting order
* Prevent in java.util
* Contains two methods
    * ```public int compare(Object obj1,Object obj2)```
        * return -ve iff obj1 has to come before obj2
        * return +ve iff obj1 has to come after obj2
        * return 0 iff obj1 and obj2 are equal
    * ```public boolean equals()```
* Defining compare method is compulsory, but equals() is optional because it is defined in Object class
* Create a class by implementing comparator interface and then define compare() method, then pass Object of this class to Constructor of class whichc you want to custom sort.
* If we return +1 from compare() method, then elements will be added in insertion order.
* If we return -1 from compare() method then elements will be added in reverse order of their insertion.
* If we return 0, then only first element will be added, rest will not be inserted.

* String buffers cannot have default sorting order, so we can write our own comparator class for them and define compare method