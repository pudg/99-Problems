use std::collections::HashMap;
use regex::Regex;
use std::iter::zip;
use std::cmp;

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

pub fn group_anagrams(words: Vec<String>) -> Vec<Vec<String>> {
    let mut anagrams: HashMap<String, Vec<String>> = HashMap::new();

    for word in words {
        let mut chars: Vec<char> = word.chars().collect();
        chars.sort_by(|a, b| a.cmp(b));
        let sorted_chars = String::from_iter(chars);
        if anagrams.contains_key(&sorted_chars) {
            anagrams.get_mut(&sorted_chars).map(|val| val.push(sorted_chars));
        } else {
            anagrams.insert(sorted_chars, vec![word]);
        }
    }    
    
    anagrams.values().cloned().collect()
}

pub fn top_k_freq(nums: Vec<i32>, k: i32) -> Vec<i32> {
    let mut freqs: HashMap<i32, i32> = HashMap::new();

    nums.into_iter()
    .for_each(|num| *freqs.entry(num).or_insert(0) += 1);

    let mut result: Vec<(i32, i32)> = freqs.into_iter().collect();
    result.sort_by(|a, b| b.1.cmp(&a.1));

    result.iter().take(k as usize).map(|x| x.0).collect()
}

pub fn longest_substr_len(s: String) -> i32 {
    let mut longest = 0;
    let mut lhs = 0;
    let mut i = 0;
    let mut seen_chars: HashMap<char, i32> = HashMap::new();
    
    for c in s.chars() {
        if seen_chars.contains_key(&c) && lhs <= seen_chars[&c] {
            lhs = seen_chars[&c] + 1;
        } else {
            longest = cmp::max(longest, i - lhs + 1);
        }
        seen_chars.entry(c).and_modify(|idx|{ *idx = i}).or_insert(i);
        i += 1;
    }

    longest
}

pub fn rotate(c: char, k: u32, base: u32) -> char {
    let mut encoded = c as u32;
    encoded += k;
    encoded -= base;
    encoded = encoded % 26;
    encoded += base;
    char::from_u32(encoded).unwrap_or(c)
}

pub fn cypher(message: String, k: u32) -> String {
    let mut encoded = String::new();
    for m in message.chars() {
        if m.is_alphanumeric() && m.is_uppercase() {
            encoded.push(rotate(m, k, 65));
        } else if m.is_alphanumeric() && m.is_lowercase() {
            encoded.push(rotate(m, k, 97));
        } else {
            encoded.push(m);
        }
    }
    encoded
}