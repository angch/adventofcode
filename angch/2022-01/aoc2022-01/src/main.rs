use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    assert!(day1("../test.txt") == (24000, 41000));
    day1("../input.txt");
}

fn day1(filename: &str) -> (i32, i32) {
    println!("Day 1: {}", filename);
    let mut elfs = Vec::new();
    let mut cal = 0;
    io::BufReader::new(File::open(filename).unwrap())
        .lines()
        .map(|line| line.unwrap().parse::<i32>())
        .for_each(|line| match line {
            Err(_x) => {
                elfs.push(-cal);
                cal = 0;
            }
            Ok(x) => {
                cal += x;
            }
        });
    if cal < 0 {
        elfs.push(-cal);
    }
    elfs.sort();
    let part1 = -elfs[0];
    // fancy way to say -elfs[0] - elfs[1] - elfs[2]
    let part2 = -elfs.iter().take(3).sum::<i32>();
    println!("Part 1: {}", part1);
    println!("Part 2: {}", part2);
    (part1, part2)
}
