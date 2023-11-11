use std::collections::{VecDeque, HashSet, HashMap};

pub fn valid_parens(parens: String) -> bool {
    let mut stack: VecDeque<char> = VecDeque::new();
    let open_parens: HashSet<char> = HashSet::from(['(', '{', '[']);
    for p in parens.chars() {
        if open_parens.contains(&p) {
            stack.push_back(p.clone());
        } else {
            if stack.len() == 0 {
                return false;
            }
            let open_paren = stack.pop_back().unwrap();
            if p == ')' && open_paren != '(' {
                return false;
            } else if p == '}' && open_paren != '{' {
                return false;
            } else if p == ']' && open_paren != '[' {
                return false;
            }
        }
    }

    stack.len() == 0
}

pub fn plus_minus(nums: Vec<i32>) {
    let mut counts: HashMap<char, f32> = HashMap::new();
    let nums_len = nums.len() as f32;
    for num in nums {
        if num > 0 {
            counts.entry('p').and_modify(|p| { *p = *p + 1.0}).or_insert(1.0);
        } else if num < 0 {
            counts.entry('n').and_modify(|n| { *n = *n + 1.0}).or_insert(1.0);
        } else {
            counts.entry('z').and_modify(|z| { *z = *z + 1.0}).or_insert(1.0);
        }
    }
    let values: Vec<f32> = counts.values().map(|x| {x / nums_len}).collect();

    println!("{:?}", values);
}

pub fn min_max_sum(nums: Vec<i32>) -> Vec<i32> {
    let mut sums: Vec<i32> = Vec::new();
    let mut sorted = nums.clone();
    sorted.sort();
    sums.push(sorted[0..=3].into_iter().sum());
    sums.push(sorted[sorted.len()-4..sorted.len()].into_iter().sum());
    sums
}