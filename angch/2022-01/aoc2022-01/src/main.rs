use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    day1("../test.txt");
    day1("../input.txt");
}
fn day1(filename: &str) {
    println!("Day 1: {}", filename);
    if let Ok(lines) = read_lines(filename) {
        // Consumes the iterator, returns an (Optional) String
        let mut elfs = Vec::new();
        let mut cal: i32 = 0;
        for line in lines {
            if let Ok(t) = line {
                if t == "" {
                    // append to elfs
                    elfs.push(-cal);
                    cal = 0;
                    continue;
                }
                let cal2 = t.parse::<i32>().unwrap();

                // println!("{}", cal2);
                cal = cal + cal2;
            }
        }
        elfs.sort();
        // println!("{:?}", elfs);
        println!("Part 1: {}", -elfs[0]);
        println!("Part 2: {}", -elfs[0] - elfs[1] - elfs[2]);
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
where
    P: AsRef<Path>,
{
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
