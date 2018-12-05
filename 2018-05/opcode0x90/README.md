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
