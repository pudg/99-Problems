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

def sales_match(nums):
    from collections import defaultdict
    count = 0
    freqs = defaultdict(int)

    for num in nums:
        freqs[num] += 1
        if freqs[num] % 2 == 0:
            count += 1

    return count
    
def book_page(num, page):
    lhs = page // 2
    rhs = (num // 2) - (page // 2)
    return min(lhs, rhs)

def min_unfair(arr, k):
    arr.sort()
    global_min = arr[-1]
    for i in range(len(arr) - k + 1):
        diff = arr[i + k - 1] - arr[i]
        global_min = min(global_min, diff)

    return global_min

def dynamic_arr(n, queries):
    arr = [[] for _ in range(n)]
    answers = []
    lastAnswer = 0
    
    for query in queries:
        if query[0] == 1:
            idx = ((query[1] ^ lastAnswer) % n)
            arr[idx].append(query[2])
        else:
            idx = ((query[1] ^ lastAnswer) % n)
            lastAnswer = arr[idx][query[2] % len(arr[idx])]
            answers.append(lastAnswer)
            
    return answers

import math
def nearest_power2(n):
    return 2 ** (math.floor(math.log(n, 2)))

def counter_game(n):
    move = True
    while n != 1:
        nearest = nearest_power2(n)
        if n != nearest:
            n = n - nearest
        else:
            n = n // 2
        if n == 1:
            break
        move = not move
    return "Player1" if move else "Player2"