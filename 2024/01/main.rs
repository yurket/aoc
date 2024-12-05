use std::fs::read_to_string;
use std::collections::HashMap;


fn read_lines(filename: &str) -> Vec<String> {
    read_to_string(filename) 
        .unwrap()  // panic on possible file-reading errors
        .lines()  // split the string into an iterator of string slices
        .map(String::from)  // make each slice into a string
        .collect()  // gather them together into a vector
}

fn part1(left: &Vec<i32>, right: &Vec<i32>) -> i32 {
    std::iter::zip(left, right)
        .map(|(l,r)| (l-r).abs())
        .sum()
}

fn part2(left: &Vec<i32>, right: &Vec<i32>) -> i32 {
    let mut counter = HashMap::new();
    for r in right {
        *counter.entry(r).or_insert(0) += 1;
    }
    // println!("{:?}", counter);

    let mut distances: i32 = 0;
    for l in left {
        if counter.contains_key(&l) {
            distances += l * counter[l]
        }
    }
    
    distances
}


fn main() {
    let lines = read_lines("input.txt");

    let mut left = Vec::new();
    let mut right = Vec::new();
    for line in lines {
        let mut two_nums = line.split_whitespace();

        let l = two_nums.next().unwrap().parse::<i32>().unwrap();
        left.push(l);
        let r = two_nums.next().unwrap().parse::<i32>().unwrap();
        right.push(r);
    }

    left.sort();
    right.sort();
    // println!("{:?}", left);
    // println!("{:?}", right);


    println!("part1 sum of distances: {}", part1(&left, &right));    
    println!("part2 sum of distances: {}", part2(&left, &right));    
}
