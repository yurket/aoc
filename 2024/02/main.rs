use std::fs::read_to_string;

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename)
        .unwrap() // panic on possible file-reading errors
        .lines() // split the string into an iterator of string slices
        .map(String::from) // make each slice into a string
        .collect() // gather them together into a vector
}

fn is_safe(line: &Vec<i32>) -> bool {
    let mut prev = line[0];
    let mut sign = 0;
    for num in line[1..].iter() {
        let current = num;
        let diff = current - prev;
        if diff.abs() < 1 || diff.abs() > 3 {
            return false;
        }

        if sign == 0 {
            sign = diff;
        }
        if sign * diff < 0 {
            return false;
        }

        prev = *current;
    }
    true
}

fn str_to_vec(line: &String) -> Vec<i32> {
    line.split(" ").map(|s| s.parse::<i32>().unwrap()).collect()
}

fn part1(lines: &Vec<String>) -> i32 {
    let safe_reports = lines.iter().fold(0, |acc, line: &String| {
        acc + is_safe(&str_to_vec(line)) as i32
    });

    safe_reports
}

fn part2(lines: &Vec<String>) -> i32 {
    let mut safe_reports = 0;
    for line in lines {
        let int_line = str_to_vec(&line);
        for (i, _) in int_line.iter().enumerate() {
            let mut int_line_copy = int_line.to_vec();
            int_line_copy.remove(i);
            if is_safe(&int_line_copy) {
                safe_reports += 1;
                break;
            }
        }
    }
    safe_reports
}

fn main() {
    // let lines = read_lines("test_input");
    let lines = read_lines("input");

    println!("Part1 safe reports: {}", part1(&lines));
    println!("Part2 safe reports: {}", part2(&lines));
}
