use std::collections::HashSet;

fn main() {
    let line = std::io::stdin().lines().next().unwrap().unwrap();
    let uniq = |n| {
        let uniq = |s: &[u8]| s.iter().collect::<HashSet<_>>().len() == s.len();
        line.as_bytes().windows(n).position(uniq).unwrap() + n
    };
    let part1 = uniq(4);
    let part2 = uniq(14);
    println!("{part1} {part2}");
}
