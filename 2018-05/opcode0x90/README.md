2018-05
=======

Solution for [Advent of Code](https://adventofcode.com/2018) day 5.

Compiling and Running
---------------------

1. Install [Rust](https://www.rust-lang.org/en-US/install.html).
2. `cargo run --release`

Result
------

```sh
part1: 10878
part2: 6874
```

Benchmark
---------

```
$ cargo bench
part1                   time:   [15.580 ms 15.905 ms 16.306 ms]
Found 12 outliers among 100 measurements (12.00%)
  6 (6.00%) high mild
  6 (6.00%) high severe

part2                   time:   [443.25 ms 445.49 ms 448.13 ms]
Found 12 outliers among 100 measurements (12.00%)
  5 (5.00%) high mild
  7 (7.00%) high severe
```

Take #2

```
$ cargo bench
part1                   time:   [475.34 us 477.72 us 480.59 us]
                        change: [-97.023% -96.961% -96.891%] (p = 0.00 < 0.05)
                        Performance has improved.
Found 10 outliers among 100 measurements (10.00%)
  5 (5.00%) high mild
  5 (5.00%) high severe

part2                   time:   [17.547 ms 17.701 ms 17.873 ms]
                        change: [-96.139% -96.079% -96.021%] (p = 0.00 < 0.05)
                        Performance has improved.
Found 2 outliers among 100 measurements (2.00%)
  1 (1.00%) high mild
  1 (1.00%) high severe
```

Take #3: after folding the laundry

```
$ cargo bench
part1                   time:   [388.15 us 389.95 us 392.13 us]
                        change: [-20.603% -19.023% -17.400%] (p = 0.00 < 0.05)
                        Performance has improved.
Found 7 outliers among 100 measurements (7.00%)
  2 (2.00%) high mild
  5 (5.00%) high severe

part2                   time:   [20.330 ms 20.351 ms 20.371 ms]
                        change: [-11.001% -10.373% -9.7182%] (p = 0.00 < 0.05)
                        Performance has improved.
Found 11 outliers among 100 measurements (11.00%)
  3 (3.00%) low mild
  1 (1.00%) high mild
  7 (7.00%) high severe
```

Test #4: powered by [rayon](https://github.com/rayon-rs/rayon).

Ignore part1 as result is dubious since no code changes was made.

```
$ cargo bench
part1                   time:   [373.70 us 375.29 us 377.21 us]
                        change: [-6.2159% -4.3964% -3.0351%] (p = 0.00 < 0.05)
                        Performance has improved.
Found 8 outliers among 100 measurements (8.00%)
  1 (1.00%) low mild
  5 (5.00%) high mild
  2 (2.00%) high severe

part2                   time:   [10.105 ms 10.140 ms 10.174 ms]
                        change: [-51.269% -50.769% -50.236%] (p = 0.00 < 0.05)
                        Performance has improved.
Found 2 outliers among 100 measurements (2.00%)
```
