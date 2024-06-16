# Binary Tree

A tree in which every node has at most two children. Forms of binary tree include:

- **Full binary tree:** Every *node* has either 0 or 2 children.
- **Complete binary tree:** All *levels* are completely filled, except possibly the last level, which is filled from left to right.
- **Perfect binary tree:** All *levels* are completely filled, and all leaves are at the same depth.
- **Balanced binary tree:** The *height* of the left and right subtrees of any node differs by ***at most one***. (The height of the left subtree of any node is at most one level short of its right, or vice-versa).
- **Degenerate tree:** Each parent node has only one associated child node (Linked List).

Observe that the type of a binary tree is based on the nature of the real-world data representing it. (e.g. Family tree, Abstract Syntax Tree, Expression tree etc.). It's not a possible solution to a problem. There's no notion like, *"let's use this type of binary tree"*.

## Binary Search Tree (BST)

A binary tree in which <u>the left child of a node contains values less than that of the node</u>, while <u>the right child contains values greater than that of the node</u>.

### Additional Operations

- Finding the Mininum and Maximum value nodes
- Successor and Predecessor

## Balanced Tree

[As defined above].

- The goal of maintaining balance is to ensure that the tree remains relatively shallow, which helps in achieving efficient search, insertion, and deletion operations.
- <u>The balance property helps prevent the tree from degenerating into a linked list</u>, which could happen in an unbalanced tree.
- The **balance factor** of a node in a balanced tree is defined as `balanceFactor = height(leftSubtree) - height(rightSubtree)`. A balanced tree has a balance factor of `-1`, `0` or `1` for every node.
- The height of a balanced tree is logarithmic in the number of nodes, making search, insertion, and deletion operations have logarithmic time complexity.

### AVL Tree

A <u>self-balancing</u> BST. If a modification operation disrupts the balance, AVL tree uses rotations to restore balance.

#### AVL Tree Operations

- Left Rotation
- Right Rotation
- Left-Right Rotation
- Right-Left Rotation

### Red-Black Tree

A <u>self-balancing</u> BST. Each node in a Red-Black tree is assigned a color, and certain properties are maintained to ensure balance. The tree is restructured during modification to maintain balance while adhering to coloring rules.

### Advantages of Balances Trees

- **Efficient Operations:** Search, insertion, and deletion operations have logarithmic time complexity, $O(logN)$, ensuring efficient perfomance.
- **Prevention of Degeneracy:** By maintaining balance, the tree avoids  degenerating into a linked list, <u>ensuring that the worst-case time complexity remains logarithmic.</u>
- **Predictable Performance:** Balanced tress provide predictable and consistent performance for various operations.

## Heap

A specialized tree-based data structure that satisfies **the heap property**. Heaps are <u>commonly used to implement **priority queues**</u> and efficiently find the maximum or minimum element in a collection.

A heap is a complete binary tree.

### "The Heap Property"

- In a **max-heap**, the value of each node is greater than or equal to the values of its children.
- In a **min-heap**, the value of each node is less than or equal to the values of its children.

### Types of Heap

- **Binary Heap:** A binary tree with the heap property. It is commonly represented using an array.
- **Binomial Heap:** A collection of binomial trees where each tree follows the heap property.
- **Fibonacci Heap:** Used particularly in graph algorithms, such as Dijkstra's algorithm.

### Operations on Heaps

- **Heapify:** Converting an array into a heap.
- **Heap sort:** Using a heap to sort an array efficiently
- Merging two heaps.
- General Tree operations, while maintaining the Heap Property.

## Trie (Prefix Tree)

A tree-like data structure used for efficiently <u>storing and retrieving a dynamic set or associative array</u> where the *keys are usually strings*. The term "Trie" comes from the word "retrieval".

### Key Characteristics of Tries

- **Tree Structure:** A tree structure where each node represents <u>a character in a key</u> or <u>part of a key</u>.
- **Node:** Each node contains a character and may have links to its children nodes. The links represent the next possible characters in the key.
- **Root Node:** The topmost node, representing an empty string or the common prefix of all keys.
- **Leaf Nodes:** They represent the end of a key.
- **Path to a Node:** The path from the root to a node represents a key.
- **Prefix Property:** <u>Tries excel at handling prefixes,</u> making them <u>suitable for tasks like **autocomplete** with partial keys</u>.

### Operations on Tries

- **Insertion:** Adding a new key to the Trie involves creating nodes for each character in the key and linking them appropriately.
- **Search:** Checking whether a key is present in the Trie involves traversing the Trie along the characters of the key.
- **Deletion:** Removing a key from the Trie requires traversing the Trie to find the key ad then removing the nodes corresponding to that key.
- **Prefix Search:** Finding all keys with a given prefix involves traversing the Trie until the end of the prfix and then collecting all descendants of that node.
- **Autocomplete:** Tries are often used in autocomplete systems to efficiently retrieve suggestions based on a prefix.

They are often used in scenarios like autocomplete systems, spell checkers, and IP routers.

## B-Tree

A B-Tree (Balanced Tree) is a self-balancing tree data structure that maintains sorted data and allows for efficient search, insertion, and deletion oeprations.

B-Trees are widely used in databases and file systems where large amounts of data need to be stored and efficiently retrieved.

They are designed to provide good performance for both random and sequential access.

### Key Characteristics

- **Balanced Structure:** B-Trees are designed to be balanced, ensuring that the depth of the tree remains relatively constant. This property allows for efficient search operations.
- **Sorted Keys:** The keys in each node are stored in sorted order, allowing for efficient search operations using binary search.
- **Multiple Keys in a Node:** Unline BSTs, B-Trees can have multiple keys in each node. This allows B-Trees to store more keys in a node, reduding the height of the tree.
- **Balancing Mechanism:** B-Trees use a balancing mechanism during insertions and deletions to ensure that the tree remains balanced.
- **Efficient Disk Access:** B-Trees are designed with disk storage in mind, and their structure allows for efficient disk access. Each node in a B-Tree corrensponds to a disk block.

### Additional Operations

- Range Queries

# Why Trees?

- Tress are a natural way to **represent hierarchical relationships** in real-world scenarios.
- They enable efficient search, insertion, and deletion operations. The hierarchical nature allows for **logarithmic time complexity** in these operations - that's fast.
- **Sorting:** Can be used to efficiently maintain a sorted collection of elements. In-order traversal of a BST results in a sorted sequence of elements.
- Trees are **efficient for hierarchical queries**. These makes them suitable for scenarios where hierarchical relationships need to be analyzed.
- Efficient memory utilization (Tries)
- Efficient **range queries** (B-Trees)
- Trees are fundamental for **implementing symbol tables, dictionaries, and associative arrays**. Search trees provide efficient key-based retrieval.
- **Expression Trees:** Trees are used to represent mathematical expressions in a hierarchical manner. Expression trees are valuable in compilers and interpreters for evaluating and optimizing mathematical expressions.

# Problems and their Tree-based solutions

| Problems | Tree |
| --- | ---- |
| Search and Retrieval Operations | BSTs: AVL, Red-black |
| Range Queries and Sequential Access | B-Trees |
| Hierarchical Relationships | Trees |
| Autocomplete and Prefix Queries | Tries |
| Expression Parsing and Evaluation | Expression trees |
| Efficient memory utilization for strings | Ternary Search Trees (TSTs) or comressed Trie structures. |
| Dynamic set operations | Splay trees. For their simplicity and amortized logarithmic time complexity in various operations |

---

# General Operations

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
