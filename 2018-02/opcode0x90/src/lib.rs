extern crate itertools;

use std::collections::HashMap;
use std::fs::File;
use std::io::{BufRead, BufReader};

use itertools::Itertools;

pub fn part1(input: &Vec<String>) -> u32 {
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

    let mut two = 0;
    let mut three = 0;

    for line in input.into_iter() {
        // compute char frequency map
        let fm = frequency_map(line);

        // check if frequency map contains exactly two/three chars
        let mut x = 0;
        let mut y = 0;

        for v in fm.values() {
            match *v {
                2 => x = 1,
                3 => y = 1,
                _ => (),
            }
            // early termination if both condition is already met
            if x > 0 && y > 0 {
                break;
            }
        }

        two += x;
        three += y;
    }

    two * three
}

pub fn part2(input: &Vec<String>) -> String {
    input
        .into_iter()
        .tuple_combinations()
        .find_map(|(a, b)| {
            // find both common and distinct chars
            let mut same = String::with_capacity(a.len());
            let mut diff = 0;

            for (a, b) in a.chars().zip(b.chars()) {
                if a == b {
                    same.push(a);
                } else {
                    diff += 1;

                    // early skipping if diff > 1
                    if diff > 1 {
                        break;
                    }
                }
            }

            match diff {
                1 => Some(same),
                _ => None,
            }
        }).expect("there is no candidate with distance of one!")
}

pub fn get_input() -> Result<Vec<String>, Box<std::error::Error>> {
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
