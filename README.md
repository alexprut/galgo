<br />
<br />

<div align="center">

<h3>Classic Algorithms and Data Structures implemented in Go</h3>

[![MIT](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/alexprut/galgo/blob/master/LICENSE)


</div>
<br />

Algorithms
==========

|Algorithm|Worst-case time complexity|Average-case time complexity|Best-case time complexity|Auxiliary space complexity|
|:---|:---|:---|:---|:---|
|[Insertion Sort](https://github.com/alexprut/galgo/algorithms/sorting/insertion_sort.go)|Θ(n^2)|Θ(n^2)|O(n)|-|
|[BubbleSort](https://github.com/alexprut/galgo/algorithms/sorting/bubble_sort.go)|O(n^2)|O(n^2)|O(n)|-|
|[MergeSort](https://github.com/alexprut/galgo/algorithms/sorting/merge_sort.go)|Θ(nlogn)|Θ(nlogn)|Θ(nlogn)|-|
|[HeapSort](https://github.com/alexprut/galgo/datastructures/max_heap.go#L82)|Θ(nlogn)|-|-|-|
|[QuickSort](https://github.com/alexprut/galgo/algorithms/sorting/quick_sort.go)|Θ(n^2)|Θ(nlogn)|Θ(nlogn)|-|
|[CountingSort](https://github.com/alexprut/galgo/algorithms/sorting/counting_sort.go)|Θ(k + n)|Θ(k + n)|Θ(n)|Θ(n)|
|[RadixSort](https://github.com/alexprut/galgo/algorithms/sorting/radix_sort.go)|Θ(d(k + n))|Θ(d(k + n))|Θ(n)|-|
|[Floyd-Warshall](https://github.com/alexprut/galgo/algorithms/graph/floyd_warshall.go)|Θ(V^3)|Θ(V^3)|Θ(V^3)|Θ(V^2)|
|[Kruskal](https://github.com/alexprut/galgo/algorithms/graph/mst/kruskal.go)|O(\|E\|log\|V\|)|-|-|-|
|[Prim](https://github.com/alexprut/galgo/algorithms/graph/mst/prim.go)|O(\|E\|log\|V\|)|-|-|-|
|[Bellman–Ford](https://github.com/alexprut/galgo/algorithms/graph/bellman_ford.go)|Θ(\|E\|\|V\|)|-|-|Θ(V)|
|[Dijkstra](https://github.com/alexprut/galgo/algorithms/graph/dijkstra.go)|O(\|E\| + \|V\|log\|V\|)|-|-|-|
|[Maximum SubArray](https://github.com/alexprut/galgo/algorithms/maximum_sub_array.go)|Θ(n)|Θ(n)|Θ(n)|Θ(1)|
|[Knuth-Morris-Pratt](https://github.com/alexprut/galgo/algorithms/knuth_morris_pratt.go)|Θ(n + m)|Θ(n)|Θ(n)|Θ(n)|
|[Rabin-Karp](https://github.com/alexprut/galgo/algorithms/RabinKarp.go)|Θ((n - m + 1)m)|Θ(n)|Θ(n)|-|
|[Greatest common divisor (GCD)](https://github.com/alexprut/galgo/algorithms/math/math.go#L13)|-|-|-|-|
|[Binary Search](https://github.com/alexprut/galgo/algorithms/search/binary_search.go)|O(nlogn)|O(nlogn)|O(1)|Θ(1)|
|[Breadth First Search (BFS)](https://github.com/alexprut/galgo/algorithms/search/breadth_first_search.go#L43)|O(\|E\|+\|V\|)|O(\|E\|+\|V\|)|-|-|
|[Depth First Search (DFS)](https://github.com/alexprut/galgo/algorithms/search/depth_first_search.go#L10)|O(\|E\|+\|V\|)|O(\|E\|+\|V\|)|-|-|
|[Topological Sort (DFS)](https://github.com/alexprut/galgo/algorithms/search/depth_first_search.go#L91)|O(\|E\|+\|V\|)|O(\|E\|+\|V\|)|-|-|
|[Longest Increasing Subsequence (LIS)](https://github.com/alexprut/galgo/utils.go#L98)|Θ(n^2)|-|-|Θ(n)|
|[Longest Common Subsequence (LCS)](https://github.com/alexprut/galgo/utils.go#L128)|Θ(n^2)|-|-|Θ(n^2)|

Data Structures
===============
|Data Structure|Methods|
|--------------|-------|
|[Max Heap](https://github.com/alexprut/galgo/datastructures/max_heap.go)|```max()``` - Θ(1), ```extractMax()``` - O(nlogn), ```increaseKey()``` - O(logn), ```insert()``` - O(logn), ```heapify()``` - O(logn), ```heapsort()``` - O(nlogn)|
|[Min Heap](https://github.com/alexprut/galgo/datastructures/min_heap.go)|```min()``` - Θ(1), ```extractMin()``` - O(nlogn), ```insert()``` - O(logn), ```heapify()``` - O(logn), ```heapsort()``` - O(nlogn)|
|[MinMax Heap](https://github.com/alexprut/galgo/datastructures/min_max_heap.go)|```min()``` - Θ(1), ```max()``` - Θ(1), ```extractMin()``` - O(nlogn), ```extractMax()``` - O(nlogn), ```insert()``` - O(logn), ```heapify()``` - O(logn)|
|[Disjoint Set](https://github.com/alexprut/galgo/datastructures/disjoint_set.go)|```makeSet()``` - Θ(1), ```findSet()``` - Θ(1), ```union()``` - Θ(1)|
|[Trie](https://github.com/alexprut/galgo/datastructures/trie.go)|```insert()``` - O(\|s\|), ```search()``` - O(\|s\|), ```searchPrefix()``` - O(\|s\|), ```remove()``` - O(\|s\|), ```size()``` - O(1)|
|[Stack](https://github.com/alexprut/galgo/datastructures/stack.go)|```push()``` - Θ(1), ```pop()``` - Θ(1), ```empty()``` - Θ(1), ```size()``` - Θ(1), ```peek()``` - Θ(1)|
|[Queue](https://github.com/alexprut/galgo/datastructures/queue.go)|```enqueue()``` - Θ(1), ```dequeue()``` - Θ(1), ```empty()``` - Θ(1), ```size()``` - Θ(1)|
|[Binary Search Tree](https://github.com/alexprut/galgo/datastructures/binary_search_tree.go)|```insert()``` - O(n), ```search()``` - O(n), ```delete()``` - O(n), ```contains()``` - O(n), ```minimum()``` - O(n), ```maximum()``` - O(n), ```size()``` - Θ(1), ```successor()``` - O(n), ```preOrderVisit()``` - O(n), ```inOrderVisit()``` - O(n), ```postOrderVisit()``` - O(n)|
|[Double Linked List](https://github.com/alexprut/galgo/datastructures/double_linked_list.go)|```insertFront()``` - Θ(1), ```removeFront()``` - Θ(1), ```insertBack()``` - Θ(1), ```removeBack()``` - Θ(1), ```head()``` - Θ(1), ```size()``` - Θ(1)|
|[Linked List](https://github.com/alexprut/galgo/datastructures/linked_list.go)|```insertFront()``` - Θ(1), ```removeFront()``` - Θ(1), ```head()``` - Θ(1), ```size()``` - Θ(1)|
|[Graph](https://github.com/alexprut/galgo/datastructures/graph.go)|```buildAdjacencyMatrix()``` - Θ(\|V\|^2), ```buildAdjacencyList()``` - Θ(\|V\| + \|E\|), ```addEdge()``` - Θ(1)|
|[Red-Black Tree](https://github.com/alexprut/galgo/datastructures/red_black_tree.go)|```insert()``` - O(logn), ```search()``` - O(logn), ```delete()``` - O(logn), ```minimum()``` - O(logn), ```maximum()``` - O(logn), ```successor()``` - O(logn)|
|[Interval Tree](https://github.com/alexprut/galgo/datastructures/interval_tree.go)|```insert()``` - O(logn), ```search()``` - O(logn), ```find()``` - O(logn), ```findAll()``` - O(n), ```delete()``` - O(logn), ```minimum()``` - O(logn), ```maximum()``` - O(logn), ```successor()``` - O(logn)|
|[Segment Tree](https://github.com/alexprut/galgo/datastructures/segment_tree.go)|```build()``` - O(n), ```update()``` - O(logn), ```search()``` - O(logn)|
|[AVL Tree](https://github.com/alexprut/galgo/datastructures/avl_tree.go)|```insert()``` - O(logn), ```search()``` - O(logn), ```delete()``` - O(logn), ```minimum()``` - O(logn), ```maximum()``` - O(logn), ```successor()``` - O(logn)|
|[B-Tree](https://github.com/alexprut/galgo/datastructures/b_tree.go)|```insert()``` - O(th), ```search()``` - O(th), ```delete()``` - O(th), ```successor()``` - O(th), ```predecessor()``` - O(th)|
|[Fibonacci Heap](https://github.com/alexprut/galgo/datastructures/fibonacci_heap.go)|```insert()``` - O(1), ```minimum()``` - O(1), ```extractMin()``` - O(logn), ```decreaseKey()``` - O(1), ```delete()``` - O(logn)|
|[Merkle Tree](https://github.com/alexprut/galgo/datastructures/merkle_tree.go)|```build()``` - O(n), ```verify()``` - O(logn), ```getProofPath()``` - O(logn)|

---

License
=======
Licensed under [MIT](https://github.com/alexprut/galgo/blob/master/LICENSE).