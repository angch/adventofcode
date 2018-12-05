use std::collections::HashSet;
use std::error::Error;
use std::io::{BufReader, Read};
use std::iter::FromIterator;

pub fn part1(input: &String) -> usize {
    let mut result = String::with_capacity(input.len());

    // simulate polymer reaction
    let stop = String::from("!");
    let mut chars = input.chars().chain(stop.chars()).peekable();
    let mut backtrack = false;

    while let (Some(_c), Some(_next)) = (chars.next(), chars.peek()) {
        let mut c = _c;
        let mut next = _next;

        if backtrack {
            // backtrack one character
            backtrack = false;
            c = result.pop().unwrap_or('!');

            if c == '!' {
                // unable to backtrack, already at beginning of string
                continue;
            }
        }

        // is this reactive?
        if c.to_ascii_lowercase() == next.to_ascii_lowercase()
            && ((c.is_ascii_lowercase() && next.is_ascii_uppercase())
                || (c.is_ascii_uppercase() && next.is_ascii_lowercase()))
        {
            // react! mark the next iteration for backtracking
            backtrack = true;
        } else {
            // not reactive, copy over to result
            result.push(c);
        }
    }

    // count the remaining units in polymer
    result.len()
}

pub fn part2(input: &String) -> usize {
    // find out all available units
    let mut units = HashSet::new();

    for c in input.chars() {
        units.insert(c.to_ascii_lowercase());
    }

    // solve for shortest polymer after deleting units
    units
        .into_iter()
        .map(|c| {
            let buf = String::from_iter(input.chars().filter_map(|x| {
                if x == c.to_ascii_lowercase() || x == c.to_ascii_uppercase() {
                    None
                } else {
                    Some(x)
                }
            }));
            let result = part1(&buf);
            result
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
