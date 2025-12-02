use std::fs;

fn read_file(filename: &str) -> Vec<String> {
    let contents = fs::read_to_string(filename).expect("Should have been able to read the file");
    let lines: Vec<String> = contents.lines().map(|s| s.to_string()).collect();
    lines
}

fn part1(lines: &[String]) -> i32 {
    let mut dial_points_to_zero_num = 0;
    let mut dial = 50;
    for line in lines {
        let direction = &line[0..1];
        let distance = line[1..].parse::<i32>().unwrap();

        match direction {
            "R" => dial += distance,
            "L" => dial -= distance,
            _ => panic!("Unknown direction {direction} in line {line}"),
        }

        dial %= 100;
        if dial == 0 {
            dial_points_to_zero_num += 1
        }
    }
    println!("Final position is {dial}");

    dial_points_to_zero_num
}

fn main() {
    let lines = read_file("input");
    // println!("input lines: {:?}", lines);

    let sol1 = part1(&lines);
    println!("Part 1 solution: {sol1}");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let lines = read_file("test_input");
        assert_eq!(part1(&lines), 3);
    }
}
