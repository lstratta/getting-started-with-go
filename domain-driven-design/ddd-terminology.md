# Domain-Driven Design Terminology

## Domain Driven Design

Domain driven design is...

### **Boundry Context**

### **Ubiquitous Language**

### **Context Maps**


## Model Driven Design

Model driven design is...

### **Entities**
1. Can be uniquely indentified using an ID
2. Consists of value objects
3. Generally persisted as a row in a DB
4. Typically mutable
5. Generally implements some business logic

### **Layered Architecture**
1. Request handlers
2. Controllers
3. Business logic
4. Persistence

Advantages:
* You can accept requests faster
* Everything is organised and well designed
* Increases flexibility, maintainability
* Resuable components


### **Value Objects**

These are immutable, light-weight objects that don’t have any identity. Value objects reduce complexity by performing complex calculations, isolating heavy computational logic from entities.

A general purpose object designed to handle primitive values. They reduce the complexity and forces ubiquitous language.

Instead of having `int price` inside `class product`, you would create a value object (a new class) called `Currency`. So you would now have `Currency price`. This allows you to encapsulate functions within the `Currency` class and prevent having to rewrite commonly used functions. It also allows you to make your code clearer by using a value objects instead of raw values or "magic numbers" that have no meaning.

Another example would be an email address. You could just create a variable of `String email` but what if there are constraints that you want to attach to that variable, such as; `maxLength`, `validEmail()`, etc? You'd have to write the functions within the class you are attributing the value to. Instead, you could create an `Email` class that contains all of these functions. It also allows you to easilt reuse the value object anywhere else without repeating the code.

Advantages:
* Don't care about uniqueness: The object can be varied and would still be valid
* Always immutable: Values cannot be changed, giving predictable results
* Rich domain logic & Auto-validating: All logic and validation associated with that object will done inside the object class
* Strong equality: ~*Need to check this definition*~
* Thread safe: Because they are immutable, they ill work even if many threads are executing it simultaneously

### **Services**

A service is the functionality between entities and value objects. Service objects (interfaces) contain operations that don't belong to any specific object. 

### **Aggregates**

### **Repositories**

Used in helping to persist aggregates. We should always create a repository per aggregate root but not for all entities.

### **Factories**

Factories help in managing beginning of lifecycle of aggregates whereas Repositories help in managing middle and end of lifecycle of an aggregate. Factories help in creating aggregates whereas Repositories help in persisting aggregates.
