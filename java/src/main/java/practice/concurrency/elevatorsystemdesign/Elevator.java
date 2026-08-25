package practice.concurrency.elevatorsystemdesign;

import java.util.Comparator;
import java.util.NavigableSet;
import java.util.TreeSet;
import java.util.concurrent.locks.ReentrantLock;

public class Elevator {
    private final String id;
    private final ReentrantLock lock = new ReentrantLock();
    
    private ElevatorDirection direction = ElevatorDirection.IDLE;
    private Short currentFloor;
    
    private final NavigableSet<Short> upStops = new TreeSet<>();
    private final NavigableSet<Short> downStops = new TreeSet<>(Comparator.reverseOrder());

    public Elevator(String id, Short initialFloor) {
        this.id = id;
        this.currentFloor = initialFloor;
    }

    public String getId() {
        return this.id;
    }

    public ElevatorDirection getDirection() {
        lock.lock();
        try {
            return this.direction;   
        } finally {
            lock.unlock();
        }
    }

    public Short getCurrentFloor() {
        lock.lock();
        try {
            return this.currentFloor;
        } finally {
            lock.unlock();
        }
    }

    public void addStop(Short floor) {
        lock.lock();

        try {
            if (floor == this.currentFloor) {
                return;
            }

            if (floor > this.currentFloor) {
                upStops.add(floor);
            }else {
                downStops.add(floor);
            }

            if (this.direction == ElevatorDirection.IDLE) {
                this.direction = floor > this.currentFloor ? ElevatorDirection.UP : ElevatorDirection.DOWN;
            }
        } finally {
            lock.unlock();
        }
    }

    public Short moveToNextStop() {
        lock.lock();
        try {
            while (this.direction != ElevatorDirection.IDLE) {
                if (this.direction == ElevatorDirection.UP) {
                    Short nextFloor = this.upStops.pollFirst();
                    if (nextFloor != null) {
                        this.currentFloor = nextFloor;
                        return nextFloor;
                    }

                    this.direction = this.downStops.isEmpty() ? ElevatorDirection.IDLE : ElevatorDirection.DOWN;
                } else {
                    Short nextFloor = this.downStops.pollFirst();

                    if (nextFloor != null) {
                        this.currentFloor = nextFloor;
                        return nextFloor;
                    }

                    this.direction = this.upStops.isEmpty() ? ElevatorDirection.IDLE : ElevatorDirection.UP;
                }
            }
            return null;
        } finally {
            lock.unlock();
        }
    }

    public int estimatedCost(Short requestedFloor) {
        lock.lock();
        try {
            int distance = Math.abs(this.currentFloor - requestedFloor);

            if (this.direction == ElevatorDirection.IDLE) {
                return distance;
            }

            boolean isMovingToward = (requestedFloor <= this.currentFloor && this.direction == ElevatorDirection.DOWN) 
            || (requestedFloor >= this.currentFloor && this.direction == ElevatorDirection.UP);

            // Penalize Elevator that is not in the right direction
            return isMovingToward ? distance : distance + 1000;
        } finally {
            lock.unlock();
        }
    }
}
