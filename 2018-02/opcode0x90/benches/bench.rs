#[macro_use]
extern crate criterion;
extern crate aoc;

use aoc::*;
use criterion::Criterion;

fn bench_part1(c: &mut Criterion) {
    c.bench_function("part1", |b| {
        let input = get_input().unwrap();
        b.iter(|| part1(&input))
    });
}

fn bench_part2(c: &mut Criterion) {
    c.bench_function("part2", |b| {
        let input = get_input().unwrap();
        b.iter(|| part2(&input))
    });
}

criterion_group!(benches, bench_part1, bench_part2);
criterion_main!(benches);
