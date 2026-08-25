package practice.concurrency.readerswritercache;
import java.util.function.Function;
/** Implement with reader/writer locking; concurrent loads for a missing key must run once. */
public final class ReadersWriterCache<K,V> {
 public V get(K key){return null;} public void put(K key,V value){} public V getOrLoad(K key, Function<? super K,? extends V> loader){return null;}
}
