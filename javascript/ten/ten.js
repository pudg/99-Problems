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

// Use two pointers to check for character equality over the string.
export const validPalind = (phrase) => {
    phrase = phrase.replace(/[^a-zA-Z0-9]/gi, '').toLowerCase();
    let lhs = 0;
    let rhs = phrase.length - 1;
    while (lhs < rhs) {
        if (phrase[lhs] != phrase[rhs]) {
            return false;
        }
        lhs += 1;
        rhs -= 1;
    }
}


// Use Newtons method to approximate the square root.
export const newtonSquareRoot = (num, tol) => {
    let x = num;

    let root = 0.5 * (x + (num / x));

    while (!(Math.abs(x - root) < tol)) {
        x = root;
        root = 0.5 * (x + (num / x));
    }
    return root
}

export const containsDups = (nums) => {
    let memory = {};

    for (let num of nums) {
        if (num in memory) {
            return true;
        }
        memory[num] = 1;
    }

    return false;
}

export const validAnagram = (s, t) => {
    s = s.split('').sort();
    t = t.split('').sort();
    let s_chars = {};
    for (let ch of s) {
        if (ch in s_chars) {
            s_chars[ch] += 1;
        } else {
            s_chars[ch] = 1;
        }
    }

    for (let ch of t) {
        if (ch in s_chars && s_chars[ch] != 0) {
            s_chars[ch] -= 1;
        } else {
            return false;
        }
    }
    return true;
}