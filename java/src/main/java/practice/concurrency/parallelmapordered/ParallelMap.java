package practice.concurrency.parallelmapordered;
import java.util.*; import java.util.function.Function;
/** Apply mapper with no more than maxConcurrency active calls, retaining input order. */
public final class ParallelMap { public static <T,R> List<R> map(List<T> input,int maxConcurrency,Function<? super T,? extends R> mapper) throws Exception { return List.of(); } }
