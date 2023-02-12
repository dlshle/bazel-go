package go.bazel.fib;

public class Fib {
    public static int calculateFib(int n) {
        if (n < 0) {
            throw new RuntimeException("invalid value for fib calculation " + n);
        }
        if (n < 3) {
            return 1;
        }
        return calculateFib(n - 1) + calculateFib(n - 2);
    }
}