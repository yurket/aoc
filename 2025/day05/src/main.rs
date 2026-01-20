use std::{cmp::max, fs};

type Range = (u128, u128);
type Id = u128;

fn read_file(filename: &str) -> (Vec<Range>, Vec<Id>) {
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
    let mut fresh_count = 0;
    for id in ids {
        for (start, end) in ranges {
            if start <= id && id <= end {
                fresh_count += 1;
                break;
            }
        }
    }
    fresh_count
}

fn merge_ranges(orig_ranges: &[Range]) -> Vec<Range> {
    let mut ranges = orig_ranges.to_vec();
    if ranges.len() < 2 {
        return ranges;
    }

    ranges.sort();
    let mut i = 0;
    let mut len = ranges.len();
    loop {
        if i + 1 > len - 1 {
            break;
        }
        let current = &ranges[i];
        let next = &ranges[i + 1];
        if next.0 >= current.0 && next.0 <= current.1 {
            ranges[i].1 = max(current.1, next.1);
            ranges.remove(i + 1);
            len -= 1;
            continue;
        }
        i += 1;
    }

    ranges
}

fn part2(ranges: &[Range]) -> u128 {
    let merged_ranges = merge_ranges(ranges);

    merged_ranges
        .iter()
        .map(|(start, end)| end - start + 1)
        .sum()
}

fn main() {
    let (ranges, ids) = read_file("input");
    println!("input ranges: {:?}", ranges);
    println!("input ids: {:?}", ids);

    let sol1 = part1(&ranges, &ids);
    println!("Part 1 solution: {sol1}");

    let sol2 = part2(&ranges);
    println!("Part 2 solution: {sol2}");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let (ranges, ids) = read_file("test_input");
        assert_eq!(part1(&ranges, &ids), 3);
    }

    #[test]
    fn test_part2() {
        let (ranges, _) = read_file("test_input");
        assert_eq!(part2(&ranges), 14);
    }
}
