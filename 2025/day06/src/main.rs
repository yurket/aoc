use std::{fs, ops::Mul};

// type Op = fn(u128, u128) -> u128;
#[derive(Debug, Clone, Copy)]
enum Op {
    Add,
    Mul,
}

fn read_file(filename: &str) -> (Vec<Vec<u128>>, Vec<Op>) {
    let contents = fs::read_to_string(filename).expect("Should have been able to read the file");
    let all_lines: Vec<&str> = contents.lines().collect();

    let nums = all_lines[..all_lines.len() - 1]
        .iter()
        .map(|s| {
            s.split_whitespace()
                .map(|w| w.parse::<u128>().unwrap())
                .collect()
        })
        .collect();

    let ops: Vec<Op> = all_lines
        .last()
        .unwrap()
        .split_whitespace()
        .map(|w| match w {
            "+" => Op::Add,
            "*" => Op::Mul,
            _ => panic!("Unknown op"),
        })
        .collect();

    return (nums, ops);
}

fn part1(filename: &str) -> u128 {
    let (nums, ops) = read_file(filename);

    let mut results: Vec<u128> = ops
        .iter()
        .map(|op| match op {
            Op::Add => 0,
            Op::Mul => 1,
            _ => panic!("Unkonwn op"),
        })
        .collect();

    for i in 0..ops.len() {
        let op = ops[i];
        for nums_row in &nums {
            match op {
                Op::Add => results[i] = results[i] + nums_row[i],
                Op::Mul => results[i] = results[i] * nums_row[i],
                _ => panic!("Unkonwn op"),
            }
        }
    }

    println!("results: {results:?}");
    results.iter().sum()
}

// fn part2() -> u128 {
//     0
// }

fn main() {
    let sol1 = part1("input");
    println!("Part 1 solution: {sol1}");

    // let sol2 = part2(&ranges);
    // println!("Part 2 solution: {sol2}");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(part1("test_input"), 4277556);
    }

    // #[test]
    // fn test_part2() {
    //     let (ranges, _) = read_file("test_input");
    //     assert_eq!(part2(&ranges), 14);
    // }
}
