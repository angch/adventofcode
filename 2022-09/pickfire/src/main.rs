use std::collections::HashSet;

fn main() {
    let mut visited1 = HashSet::new();
    let mut visited2 = HashSet::new();
    let mut knots1 = vec![(0isize, 0isize); 2];
    let mut knots2 = vec![(0isize, 0isize); 10];
    for line in std::io::stdin().lines() {
        let line = line.unwrap();
        let (dx, dy) = match &line[0..1] {
            "U" => (0, 1),
            "D" => (0, -1),
            "L" => (-1, 0),
            "R" => (1, 0),
            _ => unreachable!(),
        };
        let count: usize = line[2..].parse().unwrap();
        let step = |knots: &mut Vec<(isize, isize)>, visited: &mut HashSet<(isize, isize)>| {
            for _ in 0..count {
                knots[0].0 += dx;
                knots[0].1 += dy;
                for (h, t) in (0..).zip(1..knots.len()) {
                    if (knots[h].0 - knots[t].0).pow(2) + (knots[h].1 - knots[t].1).pow(2) > 2 {
                        knots[t].0 += (knots[h].0 - knots[t].0).signum();
                        knots[t].1 += (knots[h].1 - knots[t].1).signum();
                    }
                }
                visited.insert(*knots.last().unwrap());
            }
        };
        step(&mut knots1, &mut visited1);
        step(&mut knots2, &mut visited2);
    }
    let part1 = visited1.len();
    let part2 = visited2.len();
    println!("{part1} {part2}");
}
