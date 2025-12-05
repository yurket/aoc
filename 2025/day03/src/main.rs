use std::{
    cmp::{Reverse, max},
    fs,
};

fn read_file(filename: &str) -> Vec<Vec<i32>> {
    let contents = fs::read_to_string(filename).expect("Should have been able to read the file");
    let lines: Vec<Vec<i32>> = contents
        .lines()
        .map(|s| {
            s.chars()
                .map(|ch| ch.to_digit(10).unwrap() as i32)
                .collect()
        })
        .collect();
    lines
}

fn part1(lines: &[Vec<i32>]) -> u32 {
    let mut sum = 0;
    for line in lines {
        sum += get_max_joltage(line, 2) as u32;
    }

    sum
}

fn part2(lines: &[Vec<i32>]) -> i128 {
    let mut sum: i128 = 0;
    for line in lines {
        sum += get_max_joltage(line, 12);
    }
    sum
}

fn get_max_joltage(nums: &[i32], banks_num: i32) -> i128 {
    println!("input: {nums:?}");
    let arr: Vec<(i32, usize)> = nums.iter().enumerate().map(|(i, val)| (*val, i)).collect();

    let mut joltage: i128 = 0;

    let mut i: i32 = banks_num;
    let mut start = 0;
    let mut end: i32 = arr.len() as i32 - i;
    while i > 0 {
        let &(max_val, index) = arr[start..=end as usize]
            .iter()
            .max_by_key(|(v, index)| (v, Reverse(index)))
            .unwrap();
        println!("\tmax: {max_val}, idx: {index} [{start}..={end}]");

        start = max(start + 1, index + 1);
        i -= 1;
        joltage += (max_val as i128) * 10i128.pow(i as u32);
        end = arr.len() as i32 - i;
    }
    println!("\tjoltage: {joltage}");
    joltage
}

fn main() {
    let lines = read_file("input");
    // println!("input lines: {:?}", lines);

    let sol1 = part1(&lines);
    println!("Part 1 solution: {sol1}");

    let sol2 = part2(&lines);
    println!("Part 2 solution: {sol2}");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let lines = read_file("test_input");
        assert_eq!(part1(&lines), 357);
    }

    #[test]
    fn test_get_max_joltage() {
        let cases = vec![
            ("987654321111111", 987654321111_i128),
            ("811111111111119", 811111111119),
            ("234234234234278", 434234234278),
            ("818181911112111", 888911112111),
        ];

        for (line, expected) in cases {
            let nums: Vec<i32> = line
                .chars()
                .map(|ch| ch.to_digit(10).unwrap() as i32)
                .collect();
            assert_eq!(
                get_max_joltage(&nums, 12),
                expected,
                "Failed for line {line}"
            );
        }
    }

    #[test]
    fn test_part2() {
        let lines = read_file("test_input");
        assert_eq!(part2(&lines), 3121910778619);
    }
}
