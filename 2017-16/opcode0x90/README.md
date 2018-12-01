Benchmarks
----------

```
$ ./day16.py
Part 1: [0.045] 'bijankplfgmeodhc'
Part 2: [0.317] 'bpjahknliomefdgc'
```

For more intuition on how the part 2 fast skipping works, turn on verbose logging.

```
$ ./day16.py -v1
DEBUG:root:[999999999] Cache miss-hit: 1 - 0
DEBUG:root:[999999998] Cache miss-hit: 2 - 0
DEBUG:root:[999999997] Cache miss-hit: 3 - 0
DEBUG:root:[999999996] Cache miss-hit: 4 - 0
DEBUG:root:[999999995] Cache miss-hit: 5 - 0
DEBUG:root:[999999994] Cache miss-hit: 6 - 0
DEBUG:root:[999999993] Cache miss-hit: 7 - 0
DEBUG:root:[999999992] Cache miss-hit: 8 - 0
DEBUG:root:[999999991] Cache miss-hit: 9 - 0
DEBUG:root:[999999990] Cache miss-hit: 10 - 0
DEBUG:root:[999999989] Cache miss-hit: 11 - 0
DEBUG:root:[999999988] Cache miss-hit: 12 - 0
DEBUG:root:[999999987] Cache miss-hit: 13 - 0
DEBUG:root:[999999986] Cache miss-hit: 14 - 0
DEBUG:root:[999999985] Cache miss-hit: 15 - 0
DEBUG:root:[999999984] Cache miss-hit: 16 - 0
DEBUG:root:[999999983] Cache miss-hit: 17 - 0
DEBUG:root:[999999982] Cache miss-hit: 18 - 0
DEBUG:root:[999999981] Cache miss-hit: 19 - 0
DEBUG:root:[999999980] Cache miss-hit: 20 - 0
DEBUG:root:[999999979] Cache miss-hit: 21 - 0
DEBUG:root:[999999978] Cache miss-hit: 22 - 0
DEBUG:root:[999999977] Cache miss-hit: 23 - 0
DEBUG:root:[999999976] Cache miss-hit: 24 - 0
DEBUG:root:[999999975] Cache miss-hit: 25 - 0
DEBUG:root:[999999974] Cache miss-hit: 26 - 0
DEBUG:root:[999999973] Cache miss-hit: 27 - 0
DEBUG:root:[999999972] Cache miss-hit: 28 - 0
DEBUG:root:[999999971] Cache miss-hit: 29 - 0
DEBUG:root:[999999970] Cache miss-hit: 30 - 0
DEBUG:root:[999999969] Cache miss-hit: 31 - 0
DEBUG:root:[999999968] Cache miss-hit: 32 - 0
DEBUG:root:[999999967] Cache miss-hit: 33 - 0
DEBUG:root:[999999966] Cache miss-hit: 34 - 0
DEBUG:root:[999999965] Cache miss-hit: 35 - 0
DEBUG:root:[999999964] Cache miss-hit: 36 - 0
DEBUG:root:[999999963] Cache miss-hit: 36 - 1
DEBUG:root:[999999962] Cache miss-hit: 36 - 2
DEBUG:root:[999999961] Cache miss-hit: 36 - 3
DEBUG:root:[999999960] Cache miss-hit: 36 - 4
DEBUG:root:[999999959] Cache miss-hit: 36 - 5
DEBUG:root:[999999958] Cache miss-hit: 36 - 6
DEBUG:root:[999999957] Cache miss-hit: 36 - 7
DEBUG:root:[999999956] Cache miss-hit: 36 - 8
DEBUG:root:[999999955] Cache miss-hit: 36 - 9
DEBUG:root:[999999954] Cache miss-hit: 36 - 10
DEBUG:root:[999999953] Cache miss-hit: 36 - 11
DEBUG:root:[999999952] Cache miss-hit: 36 - 12
DEBUG:root:[999999951] Cache miss-hit: 36 - 13
DEBUG:root:[999999950] Cache miss-hit: 36 - 14
DEBUG:root:[999999949] Cache miss-hit: 36 - 15
DEBUG:root:[999999948] Cache miss-hit: 36 - 16
DEBUG:root:[999999947] Cache miss-hit: 36 - 17
DEBUG:root:[999999946] Cache miss-hit: 36 - 18
DEBUG:root:[999999945] Cache miss-hit: 36 - 19
DEBUG:root:[999999944] Cache miss-hit: 36 - 20
DEBUG:root:[999999943] Cache miss-hit: 36 - 21
DEBUG:root:[999999942] Cache miss-hit: 36 - 22
DEBUG:root:[999999941] Cache miss-hit: 36 - 23
DEBUG:root:[999999940] Cache miss-hit: 36 - 24
DEBUG:root:[999999939] Cache miss-hit: 36 - 25
DEBUG:root:[999999938] Cache miss-hit: 36 - 26
DEBUG:root:[999999937] Cache miss-hit: 36 - 27
DEBUG:root:[999999936] Cache miss-hit: 36 - 28
DEBUG:root:[999999935] Cache miss-hit: 36 - 29
DEBUG:root:[999999934] Cache miss-hit: 36 - 30
DEBUG:root:[999999933] Cache miss-hit: 36 - 31
DEBUG:root:[999999932] Cache miss-hit: 36 - 32
DEBUG:root:[999999931] Cache miss-hit: 36 - 33
DEBUG:root:[999999930] Cache miss-hit: 36 - 34
DEBUG:root:[999999929] Cache miss-hit: 36 - 35
DEBUG:root:[999999928] Cache miss-hit: 36 - 36
DEBUG:root:Cycles detected, performing fast skipping!
DEBUG:root:[27] Cache miss-hit: 36 - 37
DEBUG:root:[26] Cache miss-hit: 36 - 38
DEBUG:root:[25] Cache miss-hit: 36 - 39
DEBUG:root:[24] Cache miss-hit: 36 - 40
DEBUG:root:[23] Cache miss-hit: 36 - 41
DEBUG:root:[22] Cache miss-hit: 36 - 42
DEBUG:root:[21] Cache miss-hit: 36 - 43
DEBUG:root:[20] Cache miss-hit: 36 - 44
DEBUG:root:[19] Cache miss-hit: 36 - 45
DEBUG:root:[18] Cache miss-hit: 36 - 46
DEBUG:root:[17] Cache miss-hit: 36 - 47
DEBUG:root:[16] Cache miss-hit: 36 - 48
DEBUG:root:[15] Cache miss-hit: 36 - 49
DEBUG:root:[14] Cache miss-hit: 36 - 50
DEBUG:root:[13] Cache miss-hit: 36 - 51
DEBUG:root:[12] Cache miss-hit: 36 - 52
DEBUG:root:[11] Cache miss-hit: 36 - 53
DEBUG:root:[10] Cache miss-hit: 36 - 54
DEBUG:root:[9] Cache miss-hit: 36 - 55
DEBUG:root:[8] Cache miss-hit: 36 - 56
DEBUG:root:[7] Cache miss-hit: 36 - 57
DEBUG:root:[6] Cache miss-hit: 36 - 58
DEBUG:root:[5] Cache miss-hit: 36 - 59
DEBUG:root:[4] Cache miss-hit: 36 - 60
DEBUG:root:[3] Cache miss-hit: 36 - 61
DEBUG:root:[2] Cache miss-hit: 36 - 62
DEBUG:root:[1] Cache miss-hit: 36 - 63
DEBUG:root:[0] Cache miss-hit: 36 - 64
Part 2: [0.356] 'bpjahknliomefdgc'
```
