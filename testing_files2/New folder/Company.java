import java.util.List;
import java.util.ArrayList;

// Class Company - demonstrates composition
public class Company {
    private String companyName;
    private List<Employee> employees;

    // Constructor for Company
    public Company(String companyName) {
        this.companyName = companyName;
        this.employees = new ArrayList<>();
    }

    // Method to add an employee
    public void addEmployee(Employee employee) {
        employees.add(employee);
    }

    // Method to print details of all employees in the company
    public void printEmployeeDetails() {
        System.out.println("Company: " + companyName);
        for (Employee employee : employees) {
            employee.printDetails();  // Calling the printDetails of Employee
            System.out.println();
        }
    }
}
