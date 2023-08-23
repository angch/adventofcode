use std::collections::VecDeque;

fn main() {
    let lines = std::io::stdin().lines().map(Result::unwrap);
    let map: Vec<Vec<_>> = lines
        .map(|line| line.chars().into_iter().collect())
        .collect();
    let mut s = (0, 0);
    let mut e = (0, 0);
    let mut aa = Vec::new();
    for y in 0..map.len() {
        for x in 0..map[0].len() {
            match map[y][x] {
                'S' => s = (x, y),
                'E' => e = (x, y),
                'a' => aa.push((x, y)),
                _ => {}
            }
        }
    }
    let steps1 = steps(map.clone(), vec![s], e);
    let steps2 = steps(map, aa, e);
    println!("{steps1} {steps2}");
}

fn steps(map: Vec<Vec<char>>, ss: Vec<(usize, usize)>, e: (usize, usize)) -> usize {
    let r = map.len();
    let c = map[0].len();
    let mut parents = vec![vec![None; c]; r];
    for s in &ss {
        parents[s.1][s.0] = Some(*s);
    }
    let mut queue = VecDeque::from(ss);
    'outer: while let Some((x, y)) = queue.pop_front() {
        for (dx, dy) in [(-1, 0), (1, 0), (0, -1), (0, 1)] {
            let (nx, ny) = {
                let (nx, ny) = (x as isize + dx, y as isize + dy);
                if nx < 0 || nx as usize >= c || ny < 0 || ny as usize >= r {
                    continue;
                }
                (nx as usize, ny as usize)
            };
            let next_char = (map[y][x] as u8 + 1) as char;
            if parents[ny][nx].is_some() {
                continue;
            } else if ('a'..=next_char).contains(&map[ny][nx]) || map[y][x] == 'S' {
                parents[ny][nx] = Some((x, y));
                queue.push_back((nx, ny));
            } else if map[ny][nx] == 'E' && map[y][x] == 'z' {
                parents[ny][nx] = Some((x, y));
                break 'outer;
            }
        }
    }
    let (mut px, mut py) = e;
    let mut steps = 0;
    while (px, py) != parents[py][px].unwrap() {
        (px, py) = parents[py][px].unwrap();
        steps += 1;
    }
    steps
}
