package practice.concurrency.streamdecoder;

import java.nio.ByteBuffer;
import java.nio.ByteOrder;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.locks.ReentrantLock;
import java.util.function.Consumer;

public final class StreamDecoder {
    private static final int HEADER_SIZE = 4;
    private static final int MAX_MESSAGE_SIZE = 16 * 1024 * 1024;

    private byte[] buffer = new byte[0];

    private final ReentrantLock lock = new ReentrantLock();
    private final Consumer<byte[]> accept;

    public StreamDecoder(Consumer<byte[]> accept) {
        this.accept = accept;
    }

    public void receive(byte[] chunk) {
        List<byte[]> completedMessages = new ArrayList<>();
        
        lock.lock();
        try {
            append(chunk);
            int offset = 0;

            while (buffer.length - offset >= HEADER_SIZE) {
                int messageLength = ByteBuffer
                    .wrap(buffer, offset, HEADER_SIZE)
                    .order(ByteOrder.BIG_ENDIAN)
                    .getInt();

                if (messageLength < 0 || messageLength > MAX_MESSAGE_SIZE) {
                    throw new IllegalArgumentException(
                            "Invalid message length: " + messageLength
                    );
                }

                int frameLength = messageLength + HEADER_SIZE;

                // Header is present but message is incomplete
                if (buffer.length - offset < frameLength) {
                    break;
                }

                completedMessages.add(Arrays.copyOfRange(
                        buffer,
                        offset + HEADER_SIZE,
                        offset + frameLength
                ));
                offset += frameLength;
            }
            
            // Retain only missing bytes from message
            buffer = Arrays.copyOfRange(buffer, offset, buffer.length);
        } finally {
            lock.unlock();
        }

        for (byte[] message : completedMessages) {
            accept.accept(message);
        }
    }

    private void append(byte[] chunk) {
        byte[] combined = Arrays.copyOf(this.buffer, this.buffer.length + chunk.length);
        System.arraycopy(chunk, 0, combined, this.buffer.length, chunk.length);
        buffer = combined;
    }
}
