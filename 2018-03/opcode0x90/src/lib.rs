extern crate regex;

use std::collections::{HashMap, HashSet};
use std::error::Error;
use std::io::{BufRead, BufReader, Read};

use regex::Regex;

#[derive(Debug)]
pub struct Claim {
    id: usize,
    x: usize,
    y: usize,
    width: usize,
    height: usize,
}

#[derive(Debug)]
pub struct Mark {
    claims: HashSet<usize>,
    count: usize,
}

type Fabric = HashMap<(usize, usize), Mark>;

impl Claim {
    fn right(&self) -> usize {
        self.x + self.width
    }
    fn bottom(&self) -> usize {
        self.y + self.height
    }
    /// mark the claim on fabric
    fn mark(&self, fabric: &mut Fabric) -> () {
        for x in self.x..self.right() {
            for y in self.y..self.bottom() {
                let coord = (x, y);

                // this is mine now!
                let mark = fabric.entry(coord).or_insert(Mark {
                    claims: HashSet::new(),
                    count: 0,
                });

                // or is it?
                mark.claims.insert(self.id);
                mark.count += 1;
            }
        }
    }
}

/// I got lazy and just went with the naive bitmap solution instead
pub fn part1(input: &Vec<Claim>) -> usize {
    let mut fabric: Fabric = HashMap::new();

    // process all the claims
    for claim in input {
        claim.mark(&mut fabric);
    }

    // count areas with more than 1 claims
    fabric.values().filter(|mark| mark.count > 1).count()
}

pub fn part2(input: &Vec<Claim>) -> usize {
    let mut fabric: Fabric = HashMap::new();

    // process all the claims
    for claim in input {
        claim.mark(&mut fabric);
    }

    // locate areas that doesn't overlap and extract the claims
    let mut claims: HashSet<usize> = HashSet::new();
    let mut overlapped: HashSet<usize> = HashSet::new();

    for mark in fabric.values() {
        match mark.count {
            1 => claims.extend(mark.claims.clone()),
            _ => overlapped.extend(mark.claims.clone()),
        }
    }

    // extract non-overlapping claims
    let mut id: Vec<&usize> = claims.difference(&overlapped).collect();

    // HACK: there should be only one non-overlapping claims
    assert_eq!(1, id.len());

    *id.pop()
        .expect("unable to find any non-overlapping claims!")
}

pub fn get_input(f: impl Read) -> Result<Vec<Claim>, Box<Error>> {
    // read data from input.txt
    let input = BufReader::new(f).lines().collect::<Result<Vec<_>, _>>()?;

    // parse the input into Claim
    let re = Regex::new(r"^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$")?;
    let claims = input
        .iter()
        .filter_map(|line| {
            // attempt to parse the input line
            re.captures(line.as_str())
        }).map(|parsed| {
            let try_parse = |n| -> Result<usize, Box<Error>> {
                Ok(parsed
                    .get(n)
                    .ok_or_else(|| "malformed input")?
                    .as_str()
                    .parse::<usize>()?)
            };

            // extract regexp captured group into Claim
            Ok(Claim {
                id: try_parse(1)?,
                x: try_parse(2)?,
                y: try_parse(3)?,
                width: try_parse(4)?,
                height: try_parse(5)?,
            })
        }).collect::<Result<Vec<_>, Box<Error>>>()?;

    Ok(claims)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let data = r#"
#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
        "#;
        let input = get_input(data.as_bytes()).unwrap();
        assert_eq!(4, part1(&input));
    }

    #[test]
    fn test_part2() {
        let data = r#"
#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2
        "#;
        let input = get_input(data.as_bytes()).unwrap();
        assert_eq!(3, part2(&input));
    }
}
