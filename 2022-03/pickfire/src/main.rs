use std::collections::HashSet;
use std::io::Read;

fn find_same(parts: &[&str]) -> u32 {
    let same = parts
        .iter()
        .map(|part| part.bytes().collect::<HashSet<_>>())
        .reduce(|acc, part| acc.intersection(&part).cloned().collect::<HashSet<_>>())
        .unwrap();
    match same.into_iter().next().unwrap() {
        b @ b'a'..=b'z' => (b - b'a') as u32 + 1,
        b @ b'A'..=b'Z' => (b - b'A') as u32 + 27,
        _ => unreachable!(),
    }
}

fn main() {
    let mut buffer = String::new();
    std::io::stdin().read_to_string(&mut buffer).unwrap();

    let lines: Vec<_> = buffer.trim_end().split('\n').collect();
    let part1: u32 = lines
        .iter()
        .map(|line| {
            let parts = line.split_at(line.len() / 2);
            find_same(&[parts.0, parts.1])
        })
        .sum();
    let part2: u32 = lines.chunks(3).map(find_same).sum();
    println!("{} {}", part1, part2);
}
