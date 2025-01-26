# Go-code-search-engine
# Code Search Engine

## Overview  
Code Search Engine is a backend system designed to process Java projects and provide powerful search functionalities. It enables developers to analyze code by searching for specific elements such as classes, methods, variables, data types, and loops. Additionally, it supports exploring inheritance hierarchies, composition relationships, and locating specific patterns within Java code.

---

## Features  

### Project Upload and Processing  
- Accepts Java projects containing multiple `.java` files.  
- Parses files to extract:  
  - Classes  
  - Methods  
  - Variables  
  - Data Types (e.g., int, String, List, custom types).  
  - Loops (for, while, do-while).  
- Detects relationships:  
  - **Inheritance** (e.g., `Class A extends Class B`).  
  - **Composition** (e.g., `Class A contains instances of Class B`).  

### Search Functionalities  
- **Classes**: Search by class name.  
- **Methods**: Search by method name.  
- **Variables**: Search by variable name.  
- **Data Types**: Locate occurrences of specific data types.  
- **Loops**: Find all occurrences of loops (for, while, do-while).  
- Supports:  
  - Exact match.  
  - Substring match.  

### Class Hierarchy and Composition  
- Retrieve inheritance hierarchy of a class.  
  - Example: `Class A → inherits from Class B → inherits from Class C`.  
- Identify composition relationships for a class.  
  - Example: `Class A contains instances of Class B, Class C`.  

### Substring Search  
- Search for any substring within the Java files.  
- Return all occurrences and their locations.  

### Find Data Types  
- Locate all occurrences of specified data types (e.g., `int`, `String`, `List`, custom types).  
- Provide locations and usage contexts.  

### Find Loops  
- Extract all loops:  
  - `for`, `while`, and `do-while`.  
- Provide the context (code block or line numbers).  


