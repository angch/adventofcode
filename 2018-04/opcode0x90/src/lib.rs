extern crate regex;

use std::collections::{BTreeMap, HashMap};
use std::error::Error;
use std::io::{BufRead, BufReader, Read};

use regex::Regex;

type Minute = usize;
type GuardID = usize;

#[derive(Debug)]
enum EventType {
    Error = 0,
    BeginShift,
    FallAsleep,
    WakesUp,
}

#[derive(Debug, Clone, Eq, PartialEq, Ord, PartialOrd, Hash)]
pub struct Timestamp {
    year: usize,
    month: usize,
    day: usize,
    hour: usize,
    minute: Minute,
}

#[derive(Debug)]
pub struct Event {
    timestamp: Timestamp,
    event_type: EventType,
    guard_id: Option<GuardID>,
}

/// items in BTreeMap is always ordered, so sorting comes for free!
type Log = BTreeMap<Timestamp, Event>;

type SleepCount = usize;
type SleepMap = HashMap<Minute, SleepCount>;
type SleepSession = HashMap<GuardID, SleepMap>;

fn get_sleep_session(input: &Log) -> SleepSession {
    let mut session: SleepSession = HashMap::new();
    let mut id = 0;
    let mut sleep_start: Minute = 0;

    for event in input.values() {
        match event.event_type {
            EventType::BeginShift => id = event.guard_id.expect("guard ID is missing!"),
            EventType::FallAsleep => sleep_start = event.timestamp.minute,
            EventType::WakesUp => {
                let sleep_end = event.timestamp.minute;

                // HACK: sleep should only happen between 00:00 - 00:59
                assert!(sleep_start < sleep_end);

                // mark the sleep duration
                let sleep = session.entry(id).or_insert(SleepMap::new());

                for m in sleep_start..sleep_end {
                    let count = sleep.entry(m).or_insert(0);
                    *count += 1;
                }
            }
            _ => panic!("unknown event type encountered!"),
        }
    }
    session
}

pub fn part1(input: &Log) -> usize {
    // parse for sleep session
    let session = get_sleep_session(input);

    // who has been sleeping on the job the most?!
    let (id, sleep) = session
        .into_iter()
        .max_by_key(|(_, sleep)| -> Minute { sleep.values().sum() })
        .expect("nobody has been sleeping on the job!");

    // locate the which minute it has been sleeping the most
    let (m, _) = sleep
        .into_iter()
        .max_by_key(|(_, count)| *count)
        .expect("hmm, it was not found guilty sleeping on the job.");

    id * m
}

pub fn part2(input: &Log) -> usize {
    // parse for sleep session
    let session = get_sleep_session(input);

    // find which guard has been asleep the most on same minute, and which minute it
    let (id, sleep) = session
        .into_iter()
        .max_by_key(|(_, sleep)| -> Minute { *sleep.values().max().unwrap_or(&0) })
        .expect("nobody has been sleeping on the job!");

    // which minute is that?
    let (m, _) = sleep
        .into_iter()
        .max_by_key(|(_, count)| *count)
        .expect("hmm, it was not found guilty sleeping on the job.");

    id * m
}

pub fn get_input(f: impl Read) -> Result<Log, Box<Error>> {
    // read data from input.txt
    let input = BufReader::new(f).lines();

    // parse input into events
    let mut log: Log = Log::new();
    let re = Regex::new(r"^\[(\d+)-(\d+)-(\d+)\s+(\d+):(\d+)\]\s+(.+)$")?;
    let re_guard = Regex::new(r"Guard #(\d+) begins shift")?;

    for line in input {
        if let Some(parsed) = re.captures(line?.as_str()) {
            let try_parse = |n| -> Result<usize, Box<Error>> {
                Ok(parsed
                    .get(n)
                    .ok_or_else(|| "malformed input")?
                    .as_str()
                    .parse::<usize>()?)
            };

            let timestamp = Timestamp {
                year: try_parse(1)?,
                month: try_parse(2)?,
                day: try_parse(3)?,
                hour: try_parse(4)?,
                minute: try_parse(5)?,
            };
            let mut event = Event {
                timestamp: timestamp.clone(),
                event_type: EventType::Error,
                guard_id: None,
            };

            // what kind of event is this?
            let msg = parsed
                .get(6)
                .ok_or_else(|| "malformed input")?
                .as_str()
                .trim();

            if let Some(guard_event) = re_guard.captures(msg) {
                // a new challenger has arrived!
                event.event_type = EventType::BeginShift;
                event.guard_id = Some(
                    guard_event
                        .get(1)
                        .ok_or_else(|| "malformed input: unable to parse guard id!")?
                        .as_str()
                        .parse()?,
                );
            } else {
                event.event_type = match msg {
                    // somebody falls asleep on the job!
                    "falls asleep" => EventType::FallAsleep,
                    "wakes up" => EventType::WakesUp,
                    _ => panic!("malformed input: unrecognized event type"),
                };
            }

            // insert event into log
            log.insert(timestamp, event);
        }
    }

    Ok(log)
}

#[cfg(test)]
mod tests {
    use super::*;

    const DATA: &str = r#"
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up
        "#;

    #[test]
    fn test_part1() {
        let input = get_input(DATA.as_bytes()).unwrap();
        assert_eq!(240, part1(&input));
    }

    #[test]
    fn test_part2() {
        let input = get_input(DATA.as_bytes()).unwrap();
        assert_eq!(4455, part2(&input));
    }
}
