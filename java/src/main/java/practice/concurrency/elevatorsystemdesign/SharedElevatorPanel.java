package practice.concurrency.elevatorsystemdesign;

public class SharedElevatorPanel {
    private final ElevatorController elevatorController;

    public SharedElevatorPanel(ElevatorController elevatorController) {
        this.elevatorController = elevatorController;
    }

    public void requestElevator() {
        String elevatorId = elevatorController.requestElevator((short) 5);
        System.out.println("Please use elevator " + elevatorId);
    }
}
