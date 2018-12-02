extern crate itertools;
extern crate time;

use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::iter::FromIterator;

use itertools::Itertools;
use time::PreciseTime;

fn part1(input: &Vec<String>) -> u32 {
    fn frequency_map(s: &str) -> HashMap<char, u32> {
        // compute frequency count for each characters in given string
        let mut dict = HashMap::new();

        for c in s.chars() {
            // insert char into dict if not exist, otherwise increment count
            let i = dict.entry(c).or_insert(0);
            *i += 1
        }
        dict
    }

    let (x, y) = input.iter().fold((0, 0), |(a, b), line| {
        // compute char frequency map
        let fm = frequency_map(line);

        // check if frequency map contains exactly two/three chars
        let contains_two = fm.values().any(|v| *v == 2);
        let contains_three = fm.values().any(|v| *v == 3);

        (
            if contains_two { a + 1 } else { a },
            if contains_three { b + 1 } else { b },
        )
    });

    x * y
}

fn part2(input: &Vec<String>) -> String {
    /// compute distance between two strings
    fn distance(a: &str, b: &str) -> usize {
        let it = a.chars().zip(b.chars());

        // count number of instances where char is different from both strings
        // given the same iterator position
        it.filter(|(a, b)| a != b).count()
    }

    fn common_chars(a: &str, b: &str) -> String {
        // find common chars between two given strings
        let chars = a
            .chars()
            .zip(b.chars())
            .filter(|(a, b)| a == b)
            .map(|(a, _)| a);
        String::from_iter(chars)
    }

    // locate pairs of string where distance between them is exactly 1
    let mut candidates = input
        .into_iter()
        .tuple_combinations()
        .filter(|(a, b)| distance(a, b) == 1)
        .collect::<Vec<_>>();

    // HACK: there should not be more than one candidate
    assert_eq!(1, candidates.len());

    // find common chars
    let (a, b) = candidates
        .pop()
        .expect("There is no candidate with distance of one!");
    common_chars(a, b)
}

fn main() -> Result<(), Box<std::error::Error>> {
    // read data from input.txt
    let input = get_input()?;

    let (part1, runtime) = time(|| part1(&input));
    println!("[{}s] part1: {}", runtime, part1);

    let (part2, runtime) = time(|| part2(&input));
    println!("[{}s] part2: {}", runtime, part2);

    Ok(())
}

// Run function and return result with seconds duration
fn time<F, T>(f: F) -> (T, f64)
where
    F: FnOnce() -> T,
{
    let start = PreciseTime::now();
    let res = f();
    let end = PreciseTime::now();

    let runtime_nanos = start
        .to(end)
        .num_nanoseconds()
        .expect("Benchmark iter took greater than 2^63 nanoseconds");
    let runtime_secs = runtime_nanos as f64 / 1_000_000_000.0;
    (res, runtime_secs)
}

fn get_input() -> Result<Vec<String>, Box<std::error::Error>> {
    // read data from input.txt
    let f = File::open("input.txt").expect("input.txt not found!");
    let input = BufReader::new(f).lines().collect::<Result<Vec<_>, _>>()?;

    Ok(input)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(
            12,
            part1(
                &vec!["abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"]
                    .into_iter()
                    .map(|s| String::from(s))
                    .collect()
            )
        )
    }

    #[test]
    fn test_part2() {
        assert_eq!(
            "fgij",
            part2(
                &vec!["abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"]
                    .into_iter()
                    .map(|s| String::from(s))
                    .collect()
            )
        )
    }
}
