use std::cell::RefCell;
use std::collections::HashMap;
use std::rc::Rc;

#[derive(Debug)]
enum File {
    Dir(Rc<RefCell<HashMap<String, File>>>),
    File(usize),
}

fn dir_sizes(dir: &Rc<RefCell<HashMap<String, File>>>, dirs: &mut Vec<usize>) -> usize {
    let mut sum = 0;
    for file in (&*dir.borrow()).values() {
        match file {
            File::Dir(dir) => sum += dir_sizes(dir, dirs),
            File::File(size) => sum += size,
        }
    }
    dirs.push(sum);
    sum
}

fn main() {
    let tree = Rc::new(RefCell::new(HashMap::new()));
    let mut cwds: Vec<Rc<RefCell<HashMap<String, File>>>> = vec![tree.clone()];
    let mut lines = std::io::stdin().lines().skip(1).peekable();
    while let Some(Ok(line)) = lines.next() {
        match &line.as_bytes()[2..4] {
            b"cd" => {
                if &line.as_bytes()[5..] == b".." {
                    cwds.pop();
                } else {
                    let next = &line[5..];
                    let dir = match &cwds.last().unwrap().borrow()[next] {
                        File::Dir(dir) => dir.clone(),
                        File::File(_) => unreachable!(),
                    };
                    cwds.push(dir);
                }
            }
            b"ls" => {
                while let Some(Ok(line)) = lines.next_if(|l| !l.as_ref().unwrap().starts_with('$'))
                {
                    let mut parts = line.split(' ');
                    let size = parts.next().unwrap();
                    let name = parts.next().unwrap();
                    let file = if size == "dir" {
                        File::Dir(Rc::new(RefCell::new(HashMap::new())))
                    } else {
                        File::File(size.parse().unwrap())
                    };
                    let cwd = cwds.last_mut().unwrap();
                    cwd.borrow_mut().insert(name.to_owned(), file);
                }
            }
            _ => unreachable!(),
        };
    }
    let mut dirs = Vec::new();
    let root_free = 70000000 - dir_sizes(&tree, &mut dirs);
    let need_free = 30000000 - root_free;
    dirs.sort_unstable();
    let part1 = dirs.iter().filter(|&size| *size <= 100000).sum::<usize>();
    let part2 = dirs.iter().find(|&size| *size >= need_free).unwrap();
    println!("{part1} {part2}");
}
