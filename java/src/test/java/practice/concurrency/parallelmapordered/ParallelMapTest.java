package practice.concurrency.parallelmapordered;
import static org.junit.jupiter.api.Assertions.*; import java.util.*; import org.junit.jupiter.api.*;
class ParallelMapTest { @Test void retainsInputOrder() throws Exception { assertEquals(List.of(2,4,6),ParallelMap.map(List.of(1,2,3),2,n->n*2)); } }
