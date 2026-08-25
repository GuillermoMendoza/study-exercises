package practice.concurrency.multithreadedfizzbuzz;
import java.util.function.IntConsumer;
/** LeetCode 1195. Four methods run concurrently and must emit values 1 through n in order. */
public final class FizzBuzz {
 public FizzBuzz(int n) { if (n < 1) throw new IllegalArgumentException(); }
 public void fizz(Runnable printFizz) throws InterruptedException { }
 public void buzz(Runnable printBuzz) throws InterruptedException { }
 public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException { }
 public void number(IntConsumer printNumber) throws InterruptedException { }
}
