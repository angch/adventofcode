Benchmark
---------

```
$ ./day17.py
Part 1: [0.280] 419
Part 2: [63.713] 46038988
```

Analysis
--------

Solution shamelessly stolen from https://www.reddit.com/r/adventofcode/comments/7kc0xw/2017_day_17_solutions/drd6kck/, so I did not come up with the idea myself.

`collection.deque` version is faster than plain old manual index modulus math simply due to the requirement of _inserting a value in the middle of a list_.

Inserting value in middle of list is `O(n)` complexity, so this will very quickly add up over 5 million iterations. Contrast this with `collection.deque` rotation trick, which is only `O(k)` complexity, where k = number of steps for spinlock to step forward.

See also: https://wiki.python.org/moin/TimeComplexity
