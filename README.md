
# LinkedList Implementation in Go

This project provides a simple implementation of a singly linked list in the Go programming language. The linked list supports various common operations such as adding, deleting, and accessing nodes, as well as sorting the list using quicksort.

## Table of Contents

- [Features](#features)
- [Usage](#usage)
- [Functions](#functions)
- [Example](#example)
- [How to Run](#how-to-run)

## Features

- Add elements to the beginning or the end of the list.
- Remove elements by value or index.
- Access elements by value or index.
- Swap nodes in the list.
- Sort the list using QuickSort.
- Print the contents of the list.
- Get the length of the list.

## Usage

To use the linked list, simply import the `main` package and create a new instance of `LinkedList`. You can then use the provided methods to manipulate the list.

### Basic Operations

- **Adding Elements**
  - `addFirst(value int)` — Adds a new element to the beginning of the list.
  - `addLast(value int)` — Adds a new element to the end of the list.

- **Deleting Elements**
  - `delValue(value int)` — Removes the first node with the specified value.
  - `delIndex(index int)` — Removes the node at the specified index.

- **Accessing Elements**
  - `value(index int)` — Retrieves the value of the node at the specified index.
  - `index(value int)` — Retrieves the index of the first node with the specified value.

- **Other Operations**
  - `swap(left, right int)` — Swaps two nodes by their index.
  - `sort()` — Sorts the list using QuickSort.
  - `printList()` — Prints all elements of the list to the console.

## Functions

### LinkedList Methods

- `addFirst(value int)` — Adds a new node with the given value at the start of the list.
- `addLast(value int)` — Adds a new node with the given value at the end of the list.
- `delValue(value int)` — Deletes the first occurrence of a node with the specified value.
- `delIndex(index int)` — Deletes the node at the given index.
- `node(index int)` — Returns the node at the specified index.
- `index(value int)` — Returns the index of the first occurrence of the value.
- `value(index int)` — Returns the value at the specified index.
- `insert(index, value int)` — Inserts a node at the specified index.
- `swap(left, right int)` — Swaps two nodes in the list by their indices.
- `sort()` — Sorts the list using the QuickSort algorithm.
- `printList()` — Prints the elements of the list.

### Helper Functions

- `quickSortByCenter(l *LinkedList, start, end int)` — Helper function for performing QuickSort on the linked list.

## Example

```go
package main

import "fmt"

func main() {
    // Initialize a new linked list
    list := LinkedList{}
    
    // Add elements to the list
    list.addFirst(3)
    list.addFirst(2)
    list.addFirst(8)
    list.addLast(4)
    list.addLast(7)
    list.addLast(6)

    // Print the list
    list.printList()

    // Delete elements by value and by index
    list.delValue(3)
    list.delIndex(2)

    // Print the modified list
    list.printList()

    // Sort the list
    list.sort()
    list.printList()
}
```

## How to Run

1. Clone the repository or copy the code to your local machine.
2. Create a new Go file (e.g., `main.go`) and paste the code.
3. In the terminal, navigate to the directory where the file is located and run the following command:

    ```bash
    go run main.go
    ```

4. You should see output in the console that demonstrates the various linked list operations, including adding, deleting, and sorting elements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
