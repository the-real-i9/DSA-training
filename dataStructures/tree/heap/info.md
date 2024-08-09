# Heap

A heap is a tree-based data structure that satisfies the heap property: In a max heap, the key (value) of each child node is less than its parent's. In a min heap the key (value) of each child node is greater than its parent's.

The heap is one maximally efficient implementation of an abstract data type called a priority queue.

In a heap, the highest priority element (for max heap) or the lowest priority element (for min heap) is always stored at the root. However, a heap is not a sorted structure; it can be regarded as being partially ordered.

A heap is a useful data structure when it is necessary to repeatedly remove the object with the highest (or lowest) priority, or when insertions need to be interspersed with removals of the root node.

A common implementation of a heap is the binary heap, in which the tree is a complete binary tree. When a heap is a complete binary tree, it has the smallest possible height $-$ a heap with $N$ nodes and $a$ branches for each node always has $log_aN$ height. Thus, it can be accessed in logarithmic time complexity.

Heaps are typically constructed in-place in the same array where the elements are stored, with their structure being implicit in the access pattern of the operations.

## Operations

- Basic

  - **find-max (or find-min)**: find a maximum item of a max-heap, or a minimum item of a min-heap, respectively (a.k.a peek).
  - **insert**: adding a new key to the heap (a.k.a push)
  - **extract-max (or extract-min)**: returns the node of the maximum value from a max heap (or minimum value from a min heap) after removing it from the heap (a.k.a. pop)
  - **delete-max (or delete-min)**: removing the root node of a max heap (or min heap), respectively
  - **replace**: pop root and push a new key
- Creation
  - **create-heap**: create an empty heap
  - **heapify**: create a heap out of a given array of elements
  - **merge (union)**: joining two heaps to form a valid new heap containing all the elemtns of both, preserveing the original heaps.
  - **meld**: joining two heaps to form a valid new heap containing all the elemtns of both, destroying the original heaps.
- Inspection
  - **size**: return the number of items in the heap
  - **is-empty**: return true if the heap is empty, false otherwise
- Internal
  - **increase-key or decrease-key**: updating a key within a max- or min-heap, respectively
  - **delete**: delete an arbitrary node (followed by moving last node and sifting to maintain heap)
  - **sift-up**: move a node up in the tree, as long as needed; used to restore heap condition after insertion
  - **sift-down**: move a node down in the tree; used to restore heap condition after deletion or replacement

## Implementation

Heaps are <u>usually implemented with an **array**</u>.

- Each element in the array represents a node of the heap
- The parent/child relationship is defined implicityly by the elements' indices in the array

For a binary heap, in the array,

- the first index contains the root element.
- The next two indices of the array contain the root's children.
- The next four indices contain the four children of the root's child nodes, and so on.

Therefore, given a node at index $i$, its children are at indices $2i + 1$ (left) and $2i + 2$ (right), and its parent is at index $[(i-1)/2]$. <u>This simple indexing scheme makes it efficient to move "up" or "down" the tree</u>.

Balancing a heap is done by sift-up or sift-down operations (swapping elements which are out of order). As we can build a heap from an array without requiring extra memory (for the nodes, for example), <u>heapsort can be used to sort an array in-place</u>.

After an element is inserted into or deleted from a heap, the heap property may be violated, and <u>the heap must be re-balanced by swapping elements within the array.</u>

**Insertion:** <u>Add the new element at the end of the heap, in the first available free space.</u> If this will violate the heap property, sift up the new element (*swim* operation) until the heap property has been reestablished.

**Extraction:** <u>Remove the root and insert the last element of the heap in the root</u>. If this will violate the heap property, sift down the new root (*sink* operation) to reestablish the heap property.

**Replacement:** Remove the root and put the new element in the root and sift down. When compare to extraction followed by insertion, this avoids a sift up step.
