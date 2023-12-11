def good_pairs(nums):
    from collections import defaultdict
    counts = defaultdict(int)

    for num in nums:
        counts[num] += 1

    pairs = 0
    for key in counts.keys():
        if counts[key] > 1:
            pairs += int(counts[key] * (counts[key] - 1) / 2)
    return pairs

def unique_occurrences(nums):
    from collections import defaultdict
    counts = defaultdict(int)

    for num in nums:
        counts[num] += 1

    unique_counts = set()
    for val in counts.values():
        if val in unique_counts:
            return False
        unique_counts.add(val)
    return True
