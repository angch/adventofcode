use std::collections::BinaryHeap;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut heap = BinaryHeap::new();
    let mut current = 0;
    for line in std::io::stdin().lines() {
        let line = line?;
        if line.is_empty() {
            heap.push(current);
            current = 0;
        } else {
            current += line.parse::<u32>()?;
        }
    }
    let count1 = heap.peek().unwrap();
    let count2: u32 = heap.iter().take(3).sum();
    println!("{count1} {count2}");

    Ok(())
}
