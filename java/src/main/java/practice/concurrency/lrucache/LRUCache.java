package practice.concurrency.lrucache;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.locks.ReentrantLock;

public class LRUCache<K, V> {
    private final int capacity;
    
    private final ReentrantLock lock = new ReentrantLock();
    private final Map<K, Node<K, V>> cache = new HashMap<>();
    private final Node<K, V> head = new Node<>(null, null);
    private final Node<K, V> tail = new Node<>(null, null);

    public LRUCache(int capacity) {
        if (capacity <= 0) {
            throw new IllegalArgumentException("Capacity must be larger than 0");
        }
        this.capacity = capacity;
        head.next = tail;
        tail.prev = head;
    }

    public V get(K key) {
        lock.lock();
        try {
            Node<K, V> node = cache.get(key);
            if (node == null) {
                return null;
            }

            moveToFront(node);
            return node.value;
        } finally {
            lock.unlock();
        }
    }

    public void put(K key, V value) {
        lock.lock();
        try {
            Node<K, V> node = cache.get(key);
            if (node != null) {
                node.value = value;
                moveToFront(node);
                return;
            }

            Node<K, V> newNode = new Node<>(key, value);
            cache.put(key, newNode);
            addToFront(newNode);

            if (cache.size() > capacity) {
                Node<K, V> oldest = removeLast();
                cache.remove(oldest.key);
            }
        } finally {
            lock.unlock();
        }
    }

    private void moveToFront(Node<K, V> node) {
        remove(node);
        addToFront(node);
    }

    private void addToFront(Node<K, V> node) {
        node.next = head.next;
        node.prev = head;
        head.next.prev = node;
        head.next = node;
    }

    private void remove(Node<K, V> node) {
        node.prev.next = node.next;
        node.next.prev = node.prev;
    }

    private Node<K, V> removeLast() {
        Node<K, V> last = tail.prev;
        remove(last);
        return last;
    }

    private static class Node<K, V> {
        K key;
        V value;
        Node<K, V> prev;
        Node<K, V> next;

        Node(K key, V value) {
            this.key = key;
            this.value = value;
        }
    }
}
