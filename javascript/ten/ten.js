// Use two pointers. One slow, one fast. Perform swaps to the slow index.
export const removeDups = (nums) => {
    let slow = 1;

    for (let fast = 1; fast < nums.length; fast++) {
        if (nums[fast] != nums[ fast - 1]) {
            nums[slow] = nums[fast];
            slow++;
        }
    }
    return slow;
}

// Use a num -> index mapping, to find solution within a single pass.
export const twoSum = (nums, target) => {
    let locations = {};

    for (let idx = 0; idx < nums.length; idx++) {
        let delta = target - nums[idx];
        if (delta in locations) {
            return [idx, locations[delta]];
        } else {
            locations[nums[idx]] = idx;
        }
    }
    return [];
}