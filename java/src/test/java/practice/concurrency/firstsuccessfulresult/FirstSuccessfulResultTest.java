package practice.concurrency.firstsuccessfulresult;

import static org.junit.jupiter.api.Assertions.*; 
import java.time.Duration; 
import java.util.List; 
import org.junit.jupiter.api.*;

class FirstSuccessfulResultTest { 
    @Test void ignoresFailureAndReturnsSuccess() throws Exception { 
        assertEquals(
            "ok",
            FirstSuccessfulResult.get(
                List.of(() -> { throw new IllegalStateException(); }, () -> "ok"),
                Duration.ofSeconds(1)
            )
        ); 
    } 
}
