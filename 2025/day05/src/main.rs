use std::fs;

type Range = (u128, u128);

fn read_file(filename: &str) -> (Vec<Range>, Vec<u128>) {
    let contents = fs::read_to_string(filename).expect("Should have been able to read the file");
    let (ranges_str, ids_str) = contents.split_once("\n\n").unwrap();

    let ranges: Vec<Range> = ranges_str
        .lines()
        .map(|s| {
            let (s, e) = s.split_once('-').unwrap();
            (s.parse().unwrap(), e.parse().unwrap())
        })
        .collect();

    let ids: Vec<u128> = ids_str.lines().map(|s| s.parse().unwrap()).collect();

    (ranges, ids)
}

fn part1(ranges: &[Range], ids: &[u128]) -> i32 {
    0
}

// fn part2(lines: &mut [Vec<char>]) -> i32 {
//     0
// }

fn main() {
    let (ranges, ids) = read_file("input");
    println!("input ranges: {:?}", ranges);
    println!("input ids: {:?}", ids);

    let sol1 = part1(&ranges, &ids);
    println!("Part 1 solution: {sol1}");

    // let sol2 = part2(&mut lines);
    // println!("Part 2 solution: {sol2}");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let (ranges, ids) = read_file("test_input");
        assert_eq!(part1(&ranges, &ids), 3);
    }

    // #[test]
    // fn test_part2() {
    //     let mut lines = read_file("test_input");
    //     assert_eq!(part2(&mut lines), 43);
    // }
}
