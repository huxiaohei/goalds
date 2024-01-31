# GOALDS

GOALDS (Golang Algorithms and Data Structures) aims to enrich the data structures and algorithms in Golang.

## array

Array is a container that encapsulates fixed size arrays.

## list

List is a container that supports constant time insertion and removal of elements from anywhere in the container. Fast random access is not supported. It is implemented as a doubly-linked list.

## deque

Deque (double-ended queue) is an indexed sequence container that allows fast insertion and deletion at both its beginning and its end. In addition, insertion and deletion at either end of a deque never invalidates pointers to the rest of the elements. Deque actually implements APIs for inserting and deleting from any position. If you need to frequently insert and delete at intermediate positions while also requiring fast random access, then deque is a good choice. Indeed, when it comes to sorting, deque may not perform as well as vector. Vector provides efficient contiguous memory access, which can enhance sorting performance compared to deque, especially for large datasets. So, if sorting is a critical operation for your use case, vector would be a better choice.

## segment

A segment is a fixed capacity ring. In theory, you should not directly use it.

## queue

Queue is a container adapter that provides first-in-first-out (FIFO) data structures for insertion and deletion. By default it is implemented as an adapter on top of the deque. You can also specify a Container, as long as the container implements the Container interface

## priorityqueue

The priority queue is a container adaptor that provides constant time lookup of the largest (by default) element, at the expense of logarithmic insertion and extraction. By default it is implemented as an adapter on top of the heap.

## stack

Stack is a container adaptor that provides a LIFO (last-in-first-out) data structure for insertion and deletion. It is implemented as an adapter on top of deque.

## bitmap

Bitmap is a container that encapsulates a bit array. It is implemented as an adapter on top of the []byte. It is mainly used to solve the problem of deduplication of large data sets.

## bloomfilter

Bloomfilter is a probabilistic data structure that can quickly determine whether an element is in a set. It is implemented as an adapter on top of the bitmap. It is mainly used to solve the problem of deduplication of large data sets. Compared with bitmap, bloomfilter can save more space, but there is a certain probability of false positives. The false positive rate is related to the number of elements in the set and the size of the bitmap.
