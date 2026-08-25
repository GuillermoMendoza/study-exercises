package practice.concurrency.frequencymap;

import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;

/** Implement each nested map, progressing from unsafe to monitor and concurrent-collection variants. */
public final class FrequencyMaps {
    private FrequencyMaps() { }

    public interface FrequencyMap { 
        void increment(String key); 
        int count(String key); 
        Map<String, Integer> snapshot(); 
    }

    public static final class UnsafeFrequencyMap implements FrequencyMap {
        private final Map<String, Integer> frequencies = new HashMap<>();
        
        @Override
        public void increment(String key) {
            Objects.requireNonNull(key, "key must not be null");
            frequencies.put(key, frequencies.getOrDefault(key, 0) + 1);
        }

        @Override
        public int count(String key) { 
            Objects.requireNonNull(key, "key must not be null");
            return frequencies.getOrDefault(key, 0);
        }

        @Override
        public Map<String, Integer> snapshot() { 
            return Map.copyOf(frequencies); 
        }
    }

    public static final class LockedFrequencyMap implements FrequencyMap {
        private final Map<String, Integer> frequencies = new HashMap<>();
        
        @Override
        public synchronized void increment(String key) { 
            Objects.requireNonNull(key, "key must not be null");
            frequencies.put(key, frequencies.getOrDefault(key, 0) + 1);
        }

        @Override
        public synchronized int count(String key) { 
            Objects.requireNonNull(key, "key must not be null");
            return frequencies.getOrDefault(key, 0);
        }

        @Override
        public synchronized Map<String, Integer> snapshot() { 
            return Map.copyOf(frequencies); 
        }
    }

    public static final class ConcurrentFrequencyMap implements FrequencyMap {
        private final ConcurrentHashMap<String, Integer> frequencies = new ConcurrentHashMap<>();
        
        @Override
        public void increment(String key) { 
            Objects.requireNonNull(key, "key must not be null");
            frequencies.merge(key, 1, Integer::sum);
        }

        @Override
        public int count(String key) { 
            Objects.requireNonNull(key, "key must not be null");
            return frequencies.getOrDefault(key, 0);
        }

        @Override
        public Map<String, Integer> snapshot() { 
            return Map.copyOf(frequencies); 
        }
    }
}
