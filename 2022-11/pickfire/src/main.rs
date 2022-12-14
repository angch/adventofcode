#[derive(Debug, Clone, Copy)]
enum Value {
    Constant(usize),
    Variable, // old
}

#[derive(Debug, Clone, Copy)]
enum Op {
    Add(Value, Value),
    Mul(Value, Value),
}

#[derive(Debug)]
struct Monkey {
    items: [Vec<usize>; 2],
    operation: Op,
    test: usize,
    test_pass: usize,
    test_fail: usize,
}

fn main() {
    let mut monkeys = Vec::new();
    let mut lines = std::io::stdin().lines().map(Result::unwrap);
    while lines.next().is_some() {
        let items = lines.next().unwrap();
        let items = items.split(": ").nth(1).unwrap();
        let items: Vec<_> = items.split(", ").map(|n| n.parse().unwrap()).collect();
        let operation = lines.next().unwrap();
        let mut operation = operation.split("new = ").nth(1).unwrap().split(' ');
        let parse_op = |s| match s {
            "old" => Value::Variable,
            n => Value::Constant(n.parse().unwrap()),
        };
        let value = parse_op(operation.next().unwrap());
        let operation = match operation.next().unwrap() {
            "+" => Op::Add(value, parse_op(operation.next().unwrap())),
            "*" => Op::Mul(value, parse_op(operation.next().unwrap())),
            _ => unreachable!(),
        };
        let test = lines.next().unwrap();
        let test = test.split("by ").nth(1).unwrap().parse().unwrap();
        let test_pass = lines.next().unwrap();
        let test_pass = test_pass.split("monkey ").nth(1).unwrap().parse().unwrap();
        let test_fail = lines.next().unwrap();
        let test_fail = test_fail.split("monkey ").nth(1).unwrap().parse().unwrap();
        lines.next();
        monkeys.push(Monkey {
            items: [items.clone(), items],
            operation,
            test,
            test_pass,
            test_fail,
        });
    }
    let mut inspects = vec![vec![0; monkeys.len()]; 2];
    let mut part1 = 0usize;
    let lcm: usize = monkeys.iter().map(|m| m.test).product();
    for n in 0..10000 {
        for i in 0..monkeys.len() {
            let item_value = |item, operation| {
                let val = |v| match v {
                    Value::Constant(n) => n,
                    Value::Variable => item,
                };
                match operation {
                    Op::Add(x, y) => val(x) + val(y),
                    Op::Mul(x, y) => val(x) * val(y),
                }
            };
            let next_monkey = |item, monkey: &Monkey| {
                if item % monkey.test == 0 {
                    monkey.test_pass
                } else {
                    monkey.test_fail
                }
            };
            for mut item in std::mem::take(&mut monkeys[i].items[0]) {
                if n > 19 {
                    continue;
                }
                item = item_value(item, monkeys[i].operation) / 3;
                inspects[0][i] += 1;
                let next_monkey = next_monkey(item, &monkeys[i]);
                monkeys[next_monkey].items[0].push(item);
            }
            for mut item in std::mem::take(&mut monkeys[i].items[1]) {
                item = item_value(item, monkeys[i].operation);
                inspects[1][i] += 1;
                let next_monkey = next_monkey(item, &monkeys[i]);
                monkeys[next_monkey].items[1].push(item);
            }
        }
        for monkey in &mut monkeys {
            for item in &mut monkey.items[1] {
                *item %= lcm;
            }
        }
        if n == 19 {
            inspects[0].sort_unstable();
            part1 = inspects[0][inspects[0].len() - 2..].iter().product();
        }
    }
    inspects[1].sort_unstable();
    let part2: usize = inspects[1][inspects[1].len() - 2..].iter().product();
    println!("{part1} {part2}");
}
