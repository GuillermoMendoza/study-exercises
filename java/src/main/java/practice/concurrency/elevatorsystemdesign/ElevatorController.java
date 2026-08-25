package practice.concurrency.elevatorsystemdesign;

import java.util.List;
import java.util.concurrent.locks.ReentrantLock;

public class ElevatorController {
    private final ReentrantLock lock = new ReentrantLock();

    private final List<Elevator> elevators = List.of(
        new Elevator("A", (short) 0),
        new Elevator("B", (short) 0)
    );

    public String requestElevator(Short requestedFloor) {
        lock.lock();
        try {
            Elevator selected = elevators.stream()
            .min((optionA, optionB) -> Integer.compare(
                optionA.estimatedCost(requestedFloor),
                optionB.estimatedCost(requestedFloor)
            ))
            .orElseThrow();
            selected.addStop(requestedFloor);
            return selected.getId();
        } finally {
            lock.unlock();
        }
    }

    public void runElevator(Elevator elevator) {
        elevator.moveToNextStop();
    }

}
