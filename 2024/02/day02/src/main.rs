use std::fs::read_to_string;

fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename)
        .unwrap()  // panic on possible file-reading errors
        .lines()  // split the string into an iterator of string slices
        .map(String::from)  // make each slice into a string
        .collect()  // gather them together into a vector
}



fn is_safe(line: &String) -> bool{
    let mut line_iter = line.split(" ");
    let mut prev = line_iter.next().unwrap().parse::<i32>().unwrap();
    let mut sign = 0;
    for num_s in line_iter {
        let current = num_s.parse::<i32>().unwrap();
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

        prev = current;
    }
    true
}


fn is_safe2(line: &Vec<i32>) -> bool{
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
fn part1(lines: Vec<String>) -> i32{
    let safe_reports = lines.iter()
        .fold(0, |acc, line: &String|{acc + is_safe(line) as i32});

    safe_reports
}



// fn part2(lines: Vec<String>) -> i32{
    
// }


fn main() {
    let test_lines = read_lines("test_input");
    let lines = read_lines("input");

    let safe_reports = part1(lines);
    println!("Part1 safe reports: {}", safe_reports);



}

