//! 1 rock
//! 2 paper
//! 3 scissors
//!
//! part 1
//! choice score:
//!
//!     op (1 2 3), me (1 2 3), expected score = (!op - 2 + me) % 3
//!     draw 1 1 1 01 01 01
//!     win  1 2 2 01 10 10
//!     lose 1 3 0 01 11 00
//!     lose 2 1 0 10 01 00
//!     draw 2 2 1 10 10 01
//!     win  2 3 2 10 11 10
//!     win  3 1 2 11 01 10
//!     lose 3 2 0 11 10 00
//!     draw 3 3 1 11 11 01
//!     expected score (0 1 2) -> round score (0 3 6) = score * 3
//!
//! round score:
//!
//!     given by input
//!     input (0 1 2) -> choice (1 2 3) = me + 1
//!
//! part 2
//! round score:
//!
//!     given by input
//!     input score (0 1 2) -> round score (0 3 6) = score * 3
//!
//! choice score:
//!
//!     me (0 1 2), op (1 2 3), expected choice (1 2 3) = (me + op + 4) % 3 + 1
//!     lose 0 1 3 00 01 11
//!     lose 0 2 1 00 10 01
//!     lose 0 3 2 00 11 10
//!     draw 1 1 1 01 01 01
//!     draw 1 2 2 01 10 10
//!     draw 1 3 3 01 11 11
//!     win  2 1 2 10 01 10
//!     win  2 2 3 10 10 11
//!     win  2 3 1 10 11 01

fn main() {
    let mut score1 = 0;
    let mut score2 = 0;
    for line in std::io::stdin().lines() {
        let line = line.unwrap();
        let bytes = line.as_bytes();
        let op = bytes[0] as u32 & 3;
        let me = bytes[2] as u32 & 3;
        score1 += /* round */ (!op - 1 + me) % 3 * 3 + /* choice */ me + 1;
        score2 += /* round */ me * 3 + /* choice */ (me + op + 4) % 3 + 1;
    }
    println!("{} {}", score1, score2);
}

// old code
// #[derive(PartialEq, Clone, Copy)]
// enum Choice {
//     Rock,
//     Paper,
//     Scissors,
// }
//
// fn choice_score(choice: Choice) -> u32 {
//     match choice {
//         Choice::Rock => 1,
//         Choice::Paper => 2,
//         Choice::Scissors => 3,
//     }
// }
//
// fn round_score(op: Choice, me: Choice) -> u32 {
//     if win(op) == me {
//         6
//     } else if op == me {
//         3
//     } else {
//         0
//     }
// }
//
// fn win(choice: Choice) -> Choice {
//     match choice {
//         Choice::Rock => Choice::Paper,
//         Choice::Paper => Choice::Scissors,
//         Choice::Scissors => Choice::Rock,
//     }
// }
//
// fn lose(choice: Choice) -> Choice {
//     match choice {
//         Choice::Paper => Choice::Rock,
//         Choice::Scissors => Choice::Paper,
//         Choice::Rock => Choice::Scissors,
//     }
// }
//
// fn main() -> Result<(), Box<dyn std::error::Error>> {
//     let mut score1 = 0;
//     let mut score2 = 0;
//     for line in std::io::stdin().lines() {
//         let line = line?;
//         let mut parts = line.split(' ');
//         let op = match parts.next().unwrap() {
//             "A" => Choice::Rock,
//             "B" => Choice::Paper,
//             "C" => Choice::Scissors,
//             _ => unreachable!(),
//         };
//         let (me1, me2) = match parts.next().unwrap() {
//             "X" => (Choice::Rock, lose(op)),
//             "Y" => (Choice::Paper, op),
//             "Z" => (Choice::Scissors, win(op)),
//             _ => unreachable!(),
//         };
//         score1 += round_score(op, me1) + choice_score(me1);
//         score2 += round_score(op, me2) + choice_score(me2);
//     }
//     println!("{} {}", score1, score2);
//     Ok(())
// }
