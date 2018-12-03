use std::collections::HashSet;
use std::fs::File;
use std::io::{BufRead, BufReader};

pub fn part1(input: &Vec<i32>) -> i32 {
    // compute the sum
    input.iter().sum()
}

pub fn part2(input: &Vec<i32>) -> i32 {
    let mut delta = HashSet::new();
    let mut acc = 0;

    for x in input.iter().cycle() {
        // attempt to store each accumulator into HashSet
        // any duplicate means we found the solution
        if !delta.insert(acc) {
            break;
        }
        acc += x;
    }
    acc
}

pub fn get_input() -> Result<Vec<i32>, Box<std::error::Error>> {
    // read data from input.txt
    let f = File::open("input.txt").expect("input.txt not found!");
    let input = BufReader::new(f).lines().flatten();

    // parse the input into integers
    let parsed = input
        .into_iter()
        .map(|line| line.parse::<i32>())
        .collect::<Result<Vec<_>, _>>()?;

    Ok(parsed)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(3, part1(&vec![1, -2, 3, 1]));
        assert_eq!(3, part1(&vec![1, 1, 1]));
        assert_eq!(0, part1(&vec![1, 1, -2]));
        assert_eq!(-6, part1(&vec![-1, -2, -3]));
    }

    #[test]
    fn test_part2() {
        assert_eq!(2, part2(&vec![1, -2, 3, 1]));
        assert_eq!(0, part2(&vec![1, -1]));
        assert_eq!(10, part2(&vec![3, 3, 4, -2, -4]));
        assert_eq!(5, part2(&vec![-6, 3, 8, 5, -6]));
        assert_eq!(14, part2(&vec![7, 7, -2, -7, -4]));
    }
}
