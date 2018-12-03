extern crate aoc;

use aoc::*;

fn main() -> Result<(), Box<std::error::Error>> {
    // read data from input.txt
    let input = get_input()?;

    let part1 = part1(&input);
    println!("part1: {}", part1);

    let part2 = part2(&input);
    println!("part2: {}", part2);

    Ok(())
}
