# GOALDS

GOALDS (Golang Algorithms and Data Structures) aims to enrich the data structures and algorithms in Golang.

## array

Array is a container that encapsulates fixed size arrays.

## list

List is a container that supports constant time insertion and removal of elements from anywhere in the container. Fast random access is not supported. It is implemented as a doubly-linked list.

## vector

The storage of the vector is handled automatically, being expanded as needed. Vectors usually occupy more space than static arrays, because more memory is allocated to handle future growth. This way a vector does not need to reallocate each time an element is inserted, but only when the additional memory is exhausted. The total amount of allocated memory can be queried using Capacity() function. Extra memory can be returned to the system via a call to ShrinkToFit().

The complexity (efficiency) of common operations on vectors is as follows:

* Random access - constant $O(1)$.
* Insertion or removal of elements at the end - amortized constant $O(1)$.
* Insertion or removal of elements - linear in the distance to the end of the vector $O(n)$.

## segment

A segment is a fixed capacity ring. In theory, you should not directly use it.

## deque

Deque (double-ended queue) is an indexed sequence container that allows fast insertion and deletion at both its beginning and its end. In addition, insertion and deletion at either end of a deque never invalidates pointers to the rest of the elements. Deque actually implements APIs for inserting and deleting from any position. If you need to frequently insert and delete at intermediate positions while also requiring fast random access, then deque is a good choice. Indeed, when it comes to sorting, deque may not perform as well as vector. Vector provides efficient contiguous memory access, which can enhance sorting performance compared to deque, especially for large datasets. So, if sorting is a critical operation for your use case, vector would be a better choice.

## queue

Queue is a container adapter that provides first-in-first-out (FIFO) data structures for insertion and deletion. By default it is implemented as an adapter on top of the deque. You can also specify a Container, as long as the container implements the Container interface

## priorityqueue

The priority queue is a container adaptor that provides constant time lookup of the largest (by default) element, at the expense of logarithmic insertion and extraction. By default it is implemented as an adapter on top of the heap.

## stack

Stack is a container adaptor that provides a LIFO (last-in-first-out) data structure for insertion and deletion. It is implemented as an adapter on top of deque.

## skiplist

A skiplist is a data structure that allows for efficient search, insertion and deletion of elements in a sorted list. It is a probabilistic data structure, meaning that its average time complexity is determined through a probabilistic analysis. Skiplists have an average time complexity of $O(log_2n)$ for search, insertion and deletion, which is similar to that of balanced trees, such as AVL trees and red-black trees, but with the advantage of simpler implementation and lower overhead.

## bitmap

Bitmap is a container that encapsulates a bit array. It is implemented as an adapter on top of the []byte. It is mainly used to solve the problem of deduplication of large data sets.

## bloomfilter

Bloomfilter is a probabilistic data structure that can quickly determine whether an element is in a set. It is implemented as an adapter on top of the bitmap. It is mainly used to solve the problem of deduplication of large data sets. Compared with bitmap, bloomfilter can save more space, but there is a certain probability of false positives. The false positive rate is related to the number of elements in the set and the size of the bitmap.
