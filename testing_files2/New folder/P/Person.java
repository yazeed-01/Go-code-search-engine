// Class Person - demonstrates variables, methods, and data types
public class Person {
    // Instance variables (fields)
    private String name;
    private int age;

    // Constructor to initialize the object
    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }

    // Getter and Setter methods
    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getAge() {
        return age;
    }

    public void setAge(int age) {
        this.age = age;
    }

    // Method to print person details
    public void printDetails() {
        System.out.println("Name: " + name + ", Age: " + age);
    }
}
