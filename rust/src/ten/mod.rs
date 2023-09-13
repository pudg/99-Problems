use std::collections::HashMap;
use regex::Regex;
use std::iter::zip;

// Use two pointers. One slow, one fast. Perform swaps to the slow index.
pub fn remove_dups(nums: &mut Vec<i32>) -> usize {
    let mut slow = 1;

    for fast in 1..nums.len() {
        if nums[fast] != nums[fast - 1] {
            nums[slow] = nums[fast];
            slow += 1;
        }
    }
    slow
}

// Use a num -> index mapping, to find solution within a single pass.
pub fn two_sum(nums: &mut Vec<i32>, target: i32) -> Vec<i32> {
    let mut locations: HashMap<i32, i32> = HashMap::new();

    for num_idx in 0..nums.len() {
        let delta = target - nums[num_idx];
        if locations.contains_key(&delta) {
            return vec![num_idx as i32, locations[&delta]]
        } else {
            locations.insert(nums[num_idx], num_idx as i32);
        }
    }
    vec![]
}

//Use regex crate to remove all non-alphanumeric characters from phrase.
// Use two pointers to check for character equality over the string.
pub fn valid_palind(phrase: String) -> bool {
    let re = Regex::new(r"[^a-zA-Z0-9]").unwrap();

    let result = re.replace_all(&phrase, "")
    .to_lowercase();
    let pairs = zip(result.chars(),result.chars().rev());

    for pair in pairs {
        if pair.0 != pair.1 {
            return false;
        }
    }

    true
}
// Use Newtons method to approximate the square root.
pub fn newtons_sqroot(num: i32, tol: f32) -> f32 {
    let mut x: f32 = num as f32;
    let mut root: f32 = 0.5 * (x + (num as f32 / x));

    loop {
        if !((x - root).abs() < tol) {
            x = root;
            root = 0.5 * (x + (num as f32 / x));
        } else {
            return root;
        }
    }
}

// Use a hashmap to memorize previously seen integers.
pub fn contains_dups(nums: Vec<i32>) -> bool {
    let mut memory: HashMap<i32, i32> = HashMap::new();

    for num in nums {
        if memory.contains_key(&num) {
            return true;
        }
        memory.insert(num, 1);
    }

    false
}

// Use a hashmap to memorize previously seen integers.
pub fn contains_dups_v2(nums: Vec<i32>) -> bool {
    let mut memory: HashMap<i32, i32> = HashMap::new();

    for i in 0..nums.len() {
        if memory.contains_key(&nums[i]) {
            return true;
        }
        memory.insert(nums[i], 1);
    }
    false
}