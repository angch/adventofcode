2018-02
=======

Solution for [Advent of Code](https://adventofcode.com/2018) day 1.

Compiling and Running
---------------------

1. Install [Rust](https://www.rust-lang.org/en-US/install.html).
2. `cargo run --release`

Result
------

```
part1: 7533
part2: mphcuasvrnjzzkbgdtqeoylva
```

Benchmark
---------

```
$ cargo bench
part1                   time:   [255.55 us 257.54 us 259.73 us]
                        change: [-3.4495% -0.5399% +2.6067%] (p = 0.74 > 0.05)
                        No change in performance detected.
Found 7 outliers among 100 measurements (7.00%)
  5 (5.00%) high mild
  2 (2.00%) high severe

part2                   time:   [1.0042 ms 1.0146 ms 1.0274 ms]
                        change: [-54.652% -53.442% -52.062%] (p = 0.00 < 0.05)
                        Performance has improved.
Found 13 outliers among 100 measurements (13.00%)
  6 (6.00%) high mild
  7 (7.00%) high severe
```
