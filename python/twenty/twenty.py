def count_sort(nums):
    zeros = [0] * (max(nums) + 1)
    for num in nums:
        zeros[num] += 1

    for (i, val) in enumerate(zeros):
        if val != 0:
            print((str(i)+'\n')*val)



def pangram(s):
    from collections import defaultdict
    s = s.lower()
    chars = defaultdict(int)
    for c in s:
        if c != ' ':
            chars[c] += 1

    return len(chars.keys()) == 26