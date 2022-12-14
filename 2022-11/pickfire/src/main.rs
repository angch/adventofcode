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
    items1: Vec<usize>,
    items2: Vec<usize>,
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
            items1: items.clone(),
            items2: items,
            operation,
            test,
            test_pass,
            test_fail,
        });
    }
    let mut inspects1 = vec![0; monkeys.len()];
    let mut inspects2 = vec![0; monkeys.len()];
    let mut part1 = 0usize;
    for n in 0..10000 {
        for i in 0..monkeys.len() {
            for item in std::mem::take(&mut monkeys[i].items1) {
                if n > 19 {
                    continue;
                }
                let val = |v| match v {
                    Value::Constant(n) => n,
                    Value::Variable => item,
                };
                let item = match monkeys[i].operation {
                    Op::Add(x, y) => val(x) + val(y),
                    Op::Mul(x, y) => val(x) * val(y),
                } / 3;
                inspects1[i] += 1;
                let next_monkey = if item % monkeys[i].test == 0 {
                    monkeys[i].test_pass
                } else {
                    monkeys[i].test_fail
                };
                monkeys[next_monkey].items1.push(item);
            }
            for item in std::mem::take(&mut monkeys[i].items2) {
                let val = |v| match v {
                    Value::Constant(n) => n,
                    Value::Variable => item,
                };
                let item = match monkeys[i].operation {
                    Op::Add(x, y) => val(x) + val(y),
                    Op::Mul(x, y) => {
                        let (z, over) = val(x).overflowing_mul(val(y));
                        if over {
                            println!("over");
                        }
                        z
                    }
                };
                inspects2[i] += 1;
                let next_monkey = if item % monkeys[i].test == 0 {
                    monkeys[i].test_pass
                } else {
                    monkeys[i].test_fail
                };
                monkeys[next_monkey].items2.push(item);
            }
        }
        if n == 19 {
            inspects1.sort_unstable();
            part1 = inspects1[inspects1.len() - 2..].iter().product();
        }
        if n == 0 || n == 19 || (n > 999 && (n + 1) % 1000 == 0) {
            println!("{n} {inspects2:?}");
        }
    }
    inspects2.sort_unstable();
    let part2: usize = inspects2[inspects2.len() - 2..].iter().product();
    println!("{part1} {part2}");
}
