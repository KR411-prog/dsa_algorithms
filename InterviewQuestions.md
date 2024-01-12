Advantages of Go

Concurrency Support: Go has built-in support for concurrent programming through goroutines and channels. Goroutines are lightweight threads managed by the Go runtime, and channels provide a safe way for goroutines to communicate. This makes it easy to write concurrent and parallel programs.

Efficient Compilation: Go compiles to machine code, and its compilation process is fast, producing statically linked binaries. The resulting binaries are standalone and don't require external dependencies, simplifying deployment.

Garbage Collection: Go features automatic memory management with a garbage collector. This helps developers focus on writing code without the need to manage memory manually, reducing the risk of memory-related errors.

Simple Syntax: Go has a clean and simple syntax, making it easy to learn and read. The language design emphasizes readability and reduces unnecessary complexity, leading to more maintainable code.

Static Typing: Go is statically typed, which means that variable types are checked at compile-time. This can help catch errors early in the development process and enhances code reliability.

Strong Standard Library: Go comes with a comprehensive standard library that covers a wide range of functionalities, including networking, encryption, testing, and more. This reduces the need for third-party libraries in many cases.

========================================================
Can I just two functions with same name and a different number of arguments?
Go does not support method overloading

========================================================

OOP in Golang
Polymorphism - Allows different types to satisfy the same interface
Inheritance - Embedding interfaces. You can override the methods from the inherited (embedded) interface
Encapsulation - methods and structs can be defined with lower case first letter which make them private
========================================================
Memory allocation in Golang

In Go (Golang), memory allocation is managed by the language runtime, and it utilizes both the stack and the heap for different purposes. The decision of whether memory is allocated on the stack or the heap is made by the Go runtime based on certain criteria.

Stack Allocation: small and fixed data. : Small data with fixed sizes, like integers and pointers, are often allocated on the stack.

Heap Allocation:
Large Data or Data with Unknown Size: If the size of the data is not known at compile time or if it's too large to fit on the stack, it is allocated on the heap.

Its important to note that Go developers usually don't need to explicitly manage memory allocation and deallocation, thanks to the garbage collector and the language's automatic memory management features. Developers can focus on writing clean and efficient code, and the runtime takes care of memory allocation and deallocation.

========================================================

