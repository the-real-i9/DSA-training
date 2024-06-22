# Binary Search Tree

A BST, also called an ordered or sorted binary tree, is a binary tree with <u>the key of each internal node (inode) being greater than all keys in the respective node's left subtree and less than the ones in its right subtree.</u>

The time complexity of operations on the BST is linear with respect to the height of the tree.

**BTSs allow binary search for** fast lookup, addition, and removal of data items. Since the nodes in a BST are laid out so that each comparison skips about half of the remaining tree.

**The performance of a BST** is dependent on the order of insertioin of the nodes into the tree since arbitrary insertions may lead to degeneracy; several variations of the BST can be built with guaranteed worst-case performance.
BSTs with guaranteed worst-case complexities perform better than an unsorted array, which would require linear search time.

**The basic operations** include: search, traversal, insert and delete.

**The complexity analysis of BST** shows that, on average, the insert, delete and search takes $O(log~n)$ for $n$ nodes. In the worst-case, they degrade to that of a singly linked list: $O(n)$.

**To address the boundless increase of the tree height** with arbitrary insertions and deletions, self-balancing variants of BSTs are introduced to bound the worst lookup complexity to that of the binary logarithm. AVL trees were the first self-balancing BSTs.

**Binary search tree can be used to implement ADTs** such as dynamic sets, lookup tables and priority queues, and used in sorting algorithms such as tree sort.

---

A BST is a binary tree in which nodes are arranged in strict total order in which the nodes with keys greater than any particular node `A` is stored on the right sub-trees to that node `A` and the nodes with keys equal to or less than `A` are stored on the left sub-trees to `A`, satisfying the binary search property.

BSTS are also efficient for sortings and search algorithms. However, **the search complexity of a BST depends upon the order in which the nodes are inserted and deleted**; since in worst case successive operations in the BST may lead to degeneracy and form a singly linked list (or "unbalanced tree") like structure, thus has the same worst-case complexity as a linked list.

## Operations

### Searching

Searching in a binary tree for a specific key can be programmed recursively or iteratively.

Searching begins by examining the root node.

- If the tree is `nil`, the key being searched for does not exist in the tree.
- Otherwise, if the key equals that of the root, the search is successful and the node is returned.
- If the key is less than that of the root, the search proceeds by examining the elft subtree. Similarly, if the key is greater than that of the root, the search proceeds by examining the right subtree.
- The process is repeated until the key is found or the remaining subtree is `nil`.
- If the sarched key is not found after a `nil` subtree is reached, then the key is not present in the tree.

#### Recursive search

The recursive procedure continues until a `nil` or the `key` being sarched for are encountered.

#### Iterative search

The recursive version of the search can be "unrolled" into a while loop. <u>On most machines, the iterative version is found to be more efficient.</u>

<u>Iterative searching:</u> As long as node is non-`nil` and it's not equal to the target node, we continue to look to the left of a node if the target node is less than the node, else we look to the right. Eventually, we either get a `nil` (not found) or we get the found target node.

```go
func iterSearch(node, targetNode nodeT) nodeT {
  for node != nil && targetNode != node {
    if targetNode < node {
      node = node.left
    } else {
      node = node.right
    }
  }

  return node
}
```

Since the search may proceed till some leaf node, the running time complexity of BST search is $O(h)$ where $h$ is the height of the tree. However, the worst case for BST search is $O(n)$ where $n$ is the total number of nodes in the BST, because an unbalanced BST my degenerate to a linked list. However, if the BST is height-balanced the search is $O(log~n)$.

#### Successor and predecessor

For certain operations, given a node `x`, finding the successor or predecessor of `x` is crucial. Assuming all the keys of a BST are distinct,

- the successor of a node `x` in a BST is the node with the smallest key greater than `x`'s key. On the other hand,
- the predecessor of a node `x` in a BST is the node with the largest key smaller than `x`'s key.

<u>Finding the successor:</u> Since all nodes greater than a node's key are found on the right subtree of a node, the minimum node on the right subtree of that node is its successor.

```go
succ := node.right.findMin()
```

<u>Finding the predecessor:</u> Since all nodes less than a nodes key are found on the left subtree of a node, the maximum node on the left subtree of that node is its predecessor.

```go
succ := node.left.findMax()
```

#### Maximum and minimum

Operations such as finding a node in a BST whose key is the maximum or minimum are critical in certain operations, such as determining the successor and predecessor of nodes.

A tree of size one, has the root as its minimum and maximum nodes.

<u>Finding the maximum:</u> Since all nodes greater than a node's key are found on the right subtree of a node, the last `non-nil` node on the right is the maximum node. Practically, we keep looking on the next right node until it is `nil`, in which case the current right node is the maximum node.

```go
func findMax(node nodeT) nodeT {

  for node.right != nil {
    node = node.right
  }

  return node
}
```

<u>Finding the minimum:</u> Since all nodes less than a node's key are found on the left subtree of a node, the last `non-nil` node on the left is the minimum node. Practically, we keep looking on the next left node until it is  `nil`, in which case the current left node is the minimum node.

```go
func findMin(node nodeT) nodeT {

  for node.left != nil {
    node = node.left
  }

  return node
}
```

### Insertion

Operations such as insertion and deletion cause the BST representation to change dynamically. The structure must be modified in such a way that the properties of BST continue to hold.

<u>Inserting into a BST:</u> *(Assuming a tree's initial state is one with an empty root node.)*

If the root is empty set the root's value to the new node's value and return. Else, walk the tree in this manner: If the new node's value is less that the current node's value, consider the node's left: if it is `nil`, insert new node to current node's left and return, else proceed to the left node (sub-tree) and repeat the step. But, if the new node's value is greater then the current node's value, consider the node's right: if it is `nil`, insert new node to the current node's right and return, else proceed to the right node (sub-tree) and repeat the step.

```go
// Assuming a tree's initial state is one with an empty root node.

func treeWalkInsert(currNode *nodeT, newNode nodeT) {
  if currNode.value == nil {
    // this is an empty root
    currNode.value = newNode.value
    return
  }

  // tree walk
  for { 
    if newNode.value < currNode.value {
      if currNode.left != nil {
        // proceed to the left node (sub-tree)
        currNode = currNode.left
      } else {
        // insert to the left and return
        currNode.left = newNode
        break
      }
    } else {
      if currNode.right != nil {
        // proceed to the right node (sub-tree)
        currNode = currNode.right
      } else {
        // insert to the right and return
        currNode.right = newNode
        break
      }
    }
  }
}
```

### Deletion

*[The algorithm is as you've implemented]*.

## Traversal

**A BST can be traversed through** three basic algorithms: inorder, preorder, and postorder tree walks.

- Inorder tree walk: Nodes from the left subtree get visited first, followed by the root node and right subtree. Such a traversal visits all the nodes in the order of non-decreasing key sequence.
- Preorder tree walk: The root node gets visited first, followed by left and right subtrees.
- Postorder tree walk: Nodes from the left subtree get visited first, followed by the right subtree, and finally, the root.

## Balanced BSTs

**Without rebalancing,** insertions and deletions in a binary search may lead to degeneration, resulting in a height $n$ of the tree (where $n$ is number of items in a tree), so that the lookup performance is deteriorated to that of a linear search.

**Keeping the search tree balanced** and height bounded by $O(log~n)$ is a key to the usefulness of the BST. This can be achieved by "self-balancing" mechanisms during the modification operations to the tree designed to maintain the tree height to the binary logarithmic complexity.

### Height-balanced trees

A tree is height-balanced if the heights of the left sub-tree are guaranteed to be related by a constant factor. The heights of all the nodes on the path from the root to the modified leaf node have to be observed and possibly corrected on every insert and delete operation to the tree.

### Types

There are several self-balanced BSTs, including T-tree, treap, AVL tree, red-black tree, B-tree, 2-3 tree, and Splay tree.

## Examples of applications

### Sort

BSTs are used in sorting algorithms such as tree sort, where <u>all the elements are inserted once and the tree is traversed at an in-order fashion</u>. BSTs are also used in quick sort.

### Priority queue operations

BSTs are used in implementing priority queues, using the node's key as priorities. Adding new elements to the queue follows the reqular BST operation but removal operation depends on the type of priority queue:

- If its an ascending order priority queue, removal of an element with the lowest priority is done through leftward traversal of the BST.
- If it is a descending order priority queue, removal of an element with the highest priority is done through rightward traversal of the BST.
