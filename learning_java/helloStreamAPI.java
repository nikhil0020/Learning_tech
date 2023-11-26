package learning_java;
import java.util.Arrays;
import java.util.List;
import java.util.function.Predicate;

public class helloStreamAPI {
    public static void main(String[] args){

        List<Integer> nums = Arrays.asList(1,5,3,7,2);
        
        Predicate<Integer> predicate = new Predicate<Integer>() {
            @Override
            public boolean test(Integer integer) {
                return integer % 2 == 1;
            }
        };

        nums.stream().filter(predicate).sorted().map(n -> n*2).forEach(n -> System.out.println(n));
    }
}