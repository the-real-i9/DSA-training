# Types
## Binary Tree
A tree in which every node has at most two children.
> Types
- **Full binary tree:** Every *node* has either 0 or 2 children
- **Complete binary tree:** All *levels* are completely filled, except possibly the last level, which is filled from left to right
- **Perfect binary tree:** All *levels* are completely filled, and all leaves are at the same depth
- **Balanced binary tree:** The *height* of the left and right subtrees of any node differs by at most one. (The height of the left subtree of any node is one level short of its right, or vice-versa).
- **Degenerate tree:** Each parent node has only one associated child node. (Linked List)

Observe that the type of a binary tree is based on the nature of the real-world data representing it. (e.g. Family tree, Abstract Syntax Tree, Expression tree etc.). It's not a possible solution to a problem. There's no notion like, *"let's use this type of binary tree"*.

---

## Binary Search Tree (BST)
A binary tree in which <u>the left child of a node contains values less than that of the node</u>, while <u>the right child contains values greater than that of the node</u>.
> Additional Operations
- Finding the Mininum and Maximum value nodes
- Successor and Predecessor

## Balanced Tree
- [As defined above].
- The goal of maintaining balance is to ensure that the tree remains relatively shallow, which helps in achieving efficient search, insertion, and deletion operations.
- <u>The balance property helps prevent the tree from degenerating into a linked list</u>, which could happen in an unbalanced tree.
- The **balance factor** of a node in a balanced tree is defined as `balanceFactor = height(leftSubtree) - height(rightSubtree)`. A balanced tree has a balance factor of `-1`, `0` or `1` for every node.
- The height of a balanced tree is logarithmic in the number of nodes, making search, insertion, and deletion operations have logarithmic time complexity.

> Types 
- **AVL Tree:** A <u>self-balancing</u> BST. If a modification operation disrupts the balance, AVL tree uses rotations to restore balance.
- **Red-Black Tree:** A <u>self-balancing</u> BST. Each node in a Red-Black tree is assigned a color, and certain properties are maintained to ensure balance. The tree is restructured during modifications to maintain balance.

> AVL Tree Operations:
- Left Rotation
- Right Rotation
- Left-Right Rotation
- Right-Left Rotation

Balanced trees are crucial in scenarios where search, insertion, and deletion operations need to be performed efficiently, and maintaining a balanced structure is essential for consistent performance.

## Heap
A specialized tree-based data structure that satisfies **the heap property**. Heaps are <u>commonly used to implement **priority queues**</u> and efficiently find the maximum or minimum element in a collection.

A heap is a complete binary tree.

**The Heap Property:**
- In a **max-heap**, the value of each node is greater than or equal to the values of its children.
- In a **min-heap**, the value of each node is less than or equal to the values of its children.

> Types
- **Binary Heap:** A binary tree with the heap property.
- **Binomial Heap:** A collection of binomial trees where each tree follows the heap property.
- Fibonacci Heap

> Operations
- **Heapify:** Converting an array into a heap
- **Heap sort:** Using a heap to sort an array efficiently
- Merging two heaps into a single heap
- General Tree operations, while maintaining the Heap Property.

## Trie

## B-Tree

---

## Operations
- Enumerating all the items (Traversal)
- Enumerating a section of a tree
- Searching for an item (Search)
- Adding a new item at a certain position on the tree
- Deleting an item
- Pruning: Removing a whole section of a tree
- Grafting: Adding a whole section to a tree
- Finding the root for any node
- Finding the lowest common ancestor of two nodes

## Tree Traversal