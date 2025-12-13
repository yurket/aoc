use std::fs;

fn print_lines(lines: &[Vec<char>]) {
    for line in lines {
        for c in line {
            print!("{c} ");
        }
        println!();
    }
}

fn read_file(filename: &str) -> Vec<Vec<char>> {
    let contents = fs::read_to_string(filename).expect("Should have been able to read the file");
    let lines: Vec<Vec<char>> = contents.lines().map(|s| s.chars().collect()).collect();

    lines
}

fn get_nearby_rolls_num(field: &[Vec<char>], i: isize, j: isize) -> i32 {
    let mut rolls = 0;

    for ii in -1..=1 {
        for jj in -1..=1 {
            if (i + ii < 0 || i + ii >= field.len() as isize)
                || (j + jj < 0 || j + jj >= field[0].len() as isize)
                || (ii == 0 && jj == 0)
            {
                continue;
            }
            if field[(i + ii) as usize][(j + jj) as usize] == '@' {
                rolls += 1;
            }
        }
    }

    println!("Found {rolls} rolls at pos [{i}][{j}]");
    rolls
}

fn part1(lines: &[Vec<char>]) -> i32 {
    print_lines(&lines);

    let mut count = 0;
    for i in 0..lines.len() {
        for j in 0..lines[0].len() {
            if lines[i][j] != '@' {
                continue;
            }
            if get_nearby_rolls_num(lines, i as isize, j as isize) < 4 {
                count += 1;
            }
        }
    }

    count
}

fn part2(lines: &mut [Vec<char>]) -> i32 {
    print_lines(&lines);

    let mut count = 0;
    loop {
        let mut new_portion = 0;
        for i in 0..lines.len() {
            for j in 0..lines[0].len() {
                if lines[i][j] != '@' {
                    continue;
                }
                if get_nearby_rolls_num(lines, i as isize, j as isize) < 4 {
                    new_portion += 1;
                    lines[i][j] = 'x';
                }
            }
        }
        if new_portion == 0 {
            break;
        }
        count += new_portion;
    }
    count
}

fn main() {
    let mut lines = read_file("input");
    // println!("input lines: {:?}", lines);

    let sol1 = part1(&lines);
    println!("Part 1 solution: {sol1}");

    let sol2 = part2(&mut lines);
    println!("Part 2 solution: {sol2}");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let lines = read_file("test_input");
        assert_eq!(part1(&lines), 13);
    }

    #[test]
    fn test_part2() {
        let mut lines = read_file("test_input");
        assert_eq!(part2(&mut lines), 43);
    }
}
