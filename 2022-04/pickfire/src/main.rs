fn main() {
    let mut part1 = 0;
    let mut part2 = 0;
    for line in std::io::stdin().lines() {
        let line = line.unwrap();
        let mut parts = line.split(',');
        let x = parts.next().unwrap();
        let y = parts.next().unwrap();
        let inner = |sections: &str| {
            let mut sections = sections.split('-');
            let l: u32 = sections.next().unwrap().parse().unwrap();
            let r: u32 = sections.next().unwrap().parse().unwrap();
            (l, r)
        };
        let (xl, xr) = inner(x);
        let (yl, yr) = inner(y);
        if (xl <= yl && yr <= xr) || (yl <= xl && xr <= yr) {
            part1 += 1;
            part2 += 1;
        } else if (xl <= yl && yl <= xr) || (xl <= yr && yr <= xr) {
            part2 += 1;
        }
    }
    println!("{} {}", part1, part2);
}
