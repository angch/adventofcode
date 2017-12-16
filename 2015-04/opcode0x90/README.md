This is a demonstration of how NOT to parallelize a solution.

Let's start with our baseline.

```
$ ./day4.py test.txt
Part 1: [0.8125] 'abcdef' = (609043, '000001dbbfa3a5c83a2d506429c7b00e')
Part 1: [1.4062] 'pqrstuv' = (1048970, '000006136ef2ff3b291c85725f17325c')
Part 2: [9.4062] 'abcdef' = (6742839, '000000072a1e4320d13deee9d934ae29')
Part 2: [7.8281] 'pqrstuv' = (5714438, '000000c76bdbbb114044ada5ad14523b')
```

And compare this to our `day4_bad_parallel.py` version:

```
$ ./day4_bad_parallel.py test.txt
Part 1: [26.922] 'abcdef' = (609043, '000001dbbfa3a5c83a2d506429c7b00e')
Part 1: [48.234] 'pqrstuv' = (1048970, '000006136ef2ff3b291c85725f17325c')
Part 2: [286.531] 'abcdef' = (6742839, '000000072a1e4320d13deee9d934ae29')
Part 2: [250.078] 'pqrstuv' = (5714438, '000000c76bdbbb114044ada5ad14523b')
```

Waaaait a minute, what happened here?! Isn't parallelization supposed to automagically enhance the performance by 9001x fold??

What is actually being implemented in `day4_bad_parallel.py` is a naive _producer-consumer_ pattern. In this pattern, _producer_ creates job (which `i` to hash with and on what difficulty) for _consumer_ to consume (actually mine the hash with given job). For _producer_ to talk to _consumer_, they need to communicate via message. (more specifically, IPC)

Communication incurs overhead, and in this case the communication overhead severely negates any performance gain from we actually parallelizing it across multiple CPUs. You can read more about them here: https://distributed.readthedocs.io/en/latest/work-stealing.html

This means going back to the drawing board.
