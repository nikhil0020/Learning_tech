### STREAM

```java
List<Integer> nums = Arrays.asList(4,5,6,7,8);

Stream<Integer> data = nums.stream();

data.forEach(n -> System.out.println(n));
data.forEach(n -> System.out.println(n)); // This line will give error , because we cannot use a stream twice.

```