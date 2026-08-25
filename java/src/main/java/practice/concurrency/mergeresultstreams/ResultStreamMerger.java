package practice.concurrency.mergeresultstreams;
import java.util.*; import java.util.concurrent.BlockingQueue;
/** Each input ends with StreamEvent.complete(); preserve each input's local order. */
public final class ResultStreamMerger {
 public record StreamEvent<T>(T value, boolean complete){ public static <T> StreamEvent<T> value(T value){return new StreamEvent<>(value,false);} public static <T> StreamEvent<T> end(){return new StreamEvent<>(null,true);} }
 public static <T> List<T> merge(List<? extends BlockingQueue<StreamEvent<T>>> inputs) throws InterruptedException { return List.of(); }
}
