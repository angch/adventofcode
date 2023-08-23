use std::collections::HashSet;
use std::io::Read;

fn see<'a>(
    sees: impl Iterator<Item = usize> + 'a,
    trees: &'a [(usize, u8)],
    see_from: Option<u8>,
) -> impl Iterator<Item = usize> + 'a {
    let mut last = 0;
    let mut blocked = false;
    sees.filter_map(move |index| {
        let (pos, height) = trees[index];
        if see_from.map_or_else(|| last < height, |_| !blocked) {
            last = height;
            blocked = see_from.map_or(false, |from| from <= height);
            Some(pos)
        } else {
            None
        }
    })
}

fn main() {
    let mut trees = String::new();
    std::io::stdin().read_to_string(&mut trees).unwrap();
    let len = trees.bytes().position(|c| c == b'\n').unwrap();
    let trees: Vec<_> = trees.replace('\n', "").bytes().enumerate().collect();
    let visible: HashSet<_> = (0..len)
        .flat_map(|n| {
            see(n * len..(n + 1) * len, &trees, None)
                .chain(see((n * len..(n + 1) * len).rev(), &trees, None))
                .chain(see((n..len * len).step_by(len), &trees, None))
                .chain(see((n..len * len).step_by(len).rev(), &trees, None))
        })
        .collect();
    let part1 = visible.len();
    let part2 = (0..len * len)
        .map(|i| {
            let (y, x) = (i / len, i % len);
            let from = Some(trees[i].1);
            let u = see((x..y * len).step_by(len).rev(), &trees, from).count();
            let l = see((y * len..y * len + x).rev(), &trees, from).count();
            let d = see((y * len + len + x..len * len).step_by(len), &trees, from).count();
            let r = see(y * len + x + 1..y * len + len, &trees, from).count();
            u * l * d * r
        })
        .max()
        .unwrap();
    println!("{part1} {part2}");
}
