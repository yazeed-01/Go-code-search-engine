// Class Employee - demonstrates inheritance
public class Employee extends Person {
    // Additional field specific to Employee
    private String position;

    // Constructor for Employee (calls the super class constructor)
    public Employee(String name, int age, String position) {
        super(name, age);  // Calling the parent class constructor
        this.position = position;
    }

    // Getter and Setter methods for position
    public String getPosition() {
        return position;
    }

    public void setPosition(String position) {
        this.position = position;
    }

    // Overriding the printDetails method from Person
    @Override
    public void printDetails() {
        super.printDetails(); // Call the parent class method
        System.out.println("Position: " + position);
    }
}
