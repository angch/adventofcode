use std::collections::HashSet;
use std::error::Error;
use std::io::{BufReader, Read};
use std::iter::FromIterator;
use std::mem;

pub fn part1(input: &String) -> usize {
    // make a front and back buffer
    let mut front = String::from(input.trim());
    let mut back = String::with_capacity(input.len());
    let mut dirty = true;

    while dirty {
        // reset dirty flag
        dirty = false;

        // simulate polymer reaction
        let mut prev: char = '!';
        for c in front.drain(..) {
            // is this reactive?
            if prev.to_ascii_lowercase() == c.to_ascii_lowercase()
                && ((prev.is_ascii_lowercase() && c.is_ascii_uppercase())
                    || (prev.is_ascii_uppercase() && c.is_ascii_lowercase()))
            {
                // react! delete previous char in back buffer
                back.pop();

                // mark the dirty flag and invalidate prev char
                dirty |= true;
                prev = '!';
            } else {
                // push the char into back buffer
                back.push(c);

                // the character will remember this...
                prev = c;
            }
        }

        // swap the front and back buffer
        mem::swap(&mut front, &mut back);
    }

    // count the remaining units in polymer
    front.len()
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
