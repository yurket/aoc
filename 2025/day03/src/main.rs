use std::{cmp::Reverse, fs, i32};

fn read_file(filename: &str) -> Vec<String> {
    let contents = fs::read_to_string(filename).expect("Should have been able to read the file");
    let lines: Vec<String> = contents.lines().map(|s| s.to_string()).collect();
    lines
}

fn get_largest_num_starting_from(line: &str, start: usize) -> u32 {
    return line[start..]
        .chars()
        .map(|ch| ch.to_digit(10).unwrap())
        .max()
        .unwrap();
}

fn part1(lines: &[String]) -> u32 {
    let mut sum = 0;
    for line in lines {
        println!("line: {line}");
        let mut inventory: Vec<(u32, usize)> = line
            .chars()
            .enumerate()
            .map(|(i, ch)| (ch.to_digit(10).unwrap(), i))
            .collect();

        inventory.sort_by_key(|(key, _)| Reverse(*key));
        println!("\t {inventory:?}");

        let mut joltage = 0;
        if inventory[0].1 == line.len() - 1 {
            joltage = inventory[1].0 * 10 + inventory[0].0;
        } else {
            joltage = inventory[0].0 * 10 + get_largest_num_starting_from(line, inventory[0].1 + 1 );
        }
        println!("joltage: {joltage}");

        sum += joltage;
    }

    sum
}

// fn part2(lines: &[String]) -> i32 {

// }

fn main() {
    let lines = read_file("input");
    // println!("input lines: {:?}", lines);

    let sol1 = part1(&lines);
    println!("Part 1 solution: {sol1}");

    // let sol2 = part2(&lines);
    // println!("Part 2 solution: {sol2}");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let lines = read_file("test_input");
        assert_eq!(part1(&lines), 357);
    }

    // #[test]
    // fn test_part2() {
    //     let lines = read_file("test_input");
    //     assert_eq!(part2(&lines), 6);
    // }
}
