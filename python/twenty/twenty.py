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

def array_perm(A, B, k):
    A.sort()
    B.sort(reverse=True)
    print(k)
    for pair in zip(A, B):
        if pair[0] + pair[1] < k:
            return 'NO'
        
    return 'YES'

def subarray_division(nums, day, month):
    result = 0

    for i in range(len(nums)):
        slice = nums[i:i+month]
        if sum(slice) == day:
            result += 1

    return result
    