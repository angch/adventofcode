fn main() {
    let mut part1 = 0;
    let mut x = 1;
    let mut v = 0;
    let mut left = 0;
    let mut lines = std::io::stdin().lines();
    for cycle in 1.. {
        if left == 0 {
            let line = match lines.next() {
                Some(line) => line.unwrap(),
                None => break,
            };
            (left, v) = match &line[0..4] {
                "noop" => (1, 0),
                "addx" => (2, line[5..].parse().unwrap()),
                _ => unreachable!(),
            };
        }
        if (cycle - 20) % 40 == 0 {
            part1 += cycle * x;
        }
        if (x - 1..=x + 1).contains(&((cycle - 1) % 40)) {
            print!("#");
        } else {
            print!(".");
        }
        if cycle % 40 == 0 {
            println!();
        }
        left -= 1;
        if left == 0 {
            x += v;
        }
    }
    println!("{part1}");
}
