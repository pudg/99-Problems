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