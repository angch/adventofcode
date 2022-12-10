use std::collections::BTreeMap;
use std::io::BufRead;

fn main() {
    let mut stacks1 = BTreeMap::new();
    let mut stdin = std::io::stdin().lock();
    let mut input = String::new();
    'stacks: while stdin.read_line(&mut input).is_ok() {
        for (n, part) in (1..).zip(input.as_bytes().chunks(4)) {
            if part[1] == b'1' {
                break 'stacks;
            } else if part[1] != b' ' {
                let c = part[1] as char;
                stacks1.entry(n).or_insert_with(Vec::new).insert(0, c);
            }
        }
        input.clear();
    }
    stdin.read_line(&mut input).unwrap();
    input.clear();
    let mut stacks2 = stacks1.clone();
    while stdin.read_line(&mut input).unwrap() != 0 {
        let mut parts = input.trim().split(' ').skip(1).step_by(2);
        let n: usize = parts.next().unwrap().parse().unwrap();
        let from: u32 = parts.next().unwrap().parse().unwrap();
        let to: u32 = parts.next().unwrap().parse().unwrap();
        let stack = stacks1.get_mut(&from).unwrap();
        let v: Vec<_> = stack.drain(stack.len() - n..).rev().collect();
        stacks1.entry(to).or_insert_with(Vec::new).extend(v);
        let stack = stacks2.get_mut(&from).unwrap();
        let v: Vec<_> = stack.drain(stack.len() - n..).collect();
        stacks2.entry(to).or_insert_with(Vec::new).extend(v);
        input.clear();
    }
    let part1: String = stacks1.values().map(|s| s.last().unwrap()).collect();
    let part2: String = stacks2.values().map(|s| s.last().unwrap()).collect();
    println!("{part1} {part2}");
}
