extern crate rayon;

use std::collections::HashSet;
use std::error::Error;
use std::io::{BufReader, Read};
use std::iter::FromIterator;

use rayon::prelude::*;

pub fn part1(input: &String) -> usize {
    // simulate polymer reaction
    let result = input
        .chars()
        .fold(String::with_capacity(input.len()), |mut buf, c| {
            let tail = buf.chars().last().unwrap_or('!');

            // is this reactive?
            if c.to_ascii_lowercase() == tail.to_ascii_lowercase()
                && ((c.is_ascii_lowercase() && tail.is_ascii_uppercase())
                    || (c.is_ascii_uppercase() && tail.is_ascii_lowercase()))
            {
                // reactive! drop the last char from buffer
                buf.pop();
            } else {
                // not reactive, append the char to end of buffer
                buf.push(c);
            }
            // println!("[{}, {}] {}", tail, c, buf);
            buf
        });

    // count the remaining units in polymer
    result.len()
}

pub fn part2(input: &String) -> usize {
    // find out all available units
    let units: HashSet<char> = HashSet::from_iter(input.chars());

    // solve for shortest polymer after deleting units
    units
        .par_iter()
        .map(|c| {
            // construct a new string with given unit deleted
            let buf = String::from_iter(
                input.matches(|x| !(x == c.to_ascii_lowercase() || x == c.to_ascii_uppercase())),
            );
            // run the simulation
            part1(&buf)
        }).min()
        .expect("there is no solution!")
}

pub fn get_input(f: impl Read) -> Result<String, Box<Error>> {
    // read data from input.txt
    let mut buf = String::new();
    BufReader::new(f).read_to_string(&mut buf)?;

    // remove whitespaces
    let input = String::from(buf.trim());
    Ok(input)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(10, part1(&String::from("dabAcCaCBAcCcaDA")));
    }

    #[test]
    fn test_part2() {
        assert_eq!(4, part2(&String::from("dabAcCaCBAcCcaDA")));
    }
}
