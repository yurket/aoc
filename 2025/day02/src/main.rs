use std::fs;

type Interval = (i64, i64);

fn read_file(filename: &str) -> Vec<Interval> {
    let contents = fs::read_to_string(filename).expect("Should have been able to read the file");

    contents
        .split(',')
        .map(|s| {
            let (start, end) = s.trim().split_once('-').unwrap();
            (start.parse().unwrap(), end.parse().unwrap())
        })
        .collect()
}

fn get_first_half(num: i64, is_end: bool) -> i64 {
    let s = num.to_string();
    let mut len = s.len();
    if len == 1 {
        len = 2;
    }

    let mid = (len + is_end as usize) / 2;
    s[..mid].parse().unwrap()
}

fn write_twice(num: i64) -> i64 {
    format!("{0}{0}", num).parse().unwrap()
}

fn part1(intervals: &[Interval]) -> i64 {
    let mut sum: i64 = 0;
    for (start, end) in intervals {
        println!("Interval {start} - {end}");
        let start_half = get_first_half(*start, false);
        let end_half = get_first_half(*end, true);
        println!("\t{start_half}, {end_half}");

        for n in start_half..=end_half {
            let wrong_id = write_twice(n);
            if wrong_id >= *start && wrong_id <= *end {
                println!("\tWrong id: {wrong_id}");
                sum += wrong_id;
            }
        }
    }
    sum
}

fn part2(intervals: &[Interval]) -> i64 {
    0
}

fn main() {
    let intervals = read_file("input");
    // println!("input lines: {:?}", lines);

    let sol1 = part1(&intervals);
    println!("Part 1 solution: {sol1}");

    let sol2 = part2(&intervals);
    println!("Part 2 solution: {sol2}");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_first_half() {
        let cases = vec![
            (1, 1, false),
            (21, 2, false),
            (100, 1, false),
            (100, 10, true),
            (1221, 12, false),
        ];

        for (input, expected, is_end) in cases {
            assert_eq!(
                get_first_half(input, is_end),
                expected,
                "Failed for input: {}",
                input
            );
        }
    }

    #[test]
    fn test_part1() {
        let intervals = read_file("test_input");
        assert_eq!(part1(&intervals), 1227775554);
    }

    // #[test]
    // fn test_part2() {
    //     let intervals = read_file("test_input");
    //     assert_eq!(part2(&intervals), 6);
    // }
}
