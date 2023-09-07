use std::collections::HashMap;

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

// Use a hash map to store index locations of values.
// Finds solution in a single pass through the slice of nums.
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