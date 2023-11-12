def grid(grid):
    rearranged = [''.join(sorted(list(row))) for row in grid]
    
    for i in range(len(rearranged[0])):
        col = [rearranged[j][i] for j in range(len(rearranged))]
        asc_col = sorted(col)
        
        if ''.join(asc_col) != ''.join(col):
            return "NO"
        
    return "YES"
    
class Node:
    def __init__(self, data=None, nxt=None):
        self.data = data
        self.nxt = nxt
    
def merge_lists(l1, l2):
    merged = Node()
    helper = merged

    while l1 and l2:
        if l1.data < l2.data:
            helper.nxt = Node(l1.data)
            l1 = l1.nxt
            helper = helper.nxt
        elif l2.data < l1.data:
            helper.nxt = Node(l2.data)
            l2 = l2.nxt
            helper = helper.nxt
        else:
            helper.nxt = Node(l1.data)
            helper = helper.nxt
            helper.nxt = Node(l2.data)
            helper = helper.nxt
            l1, l2 = l1.nxt, l2.nxt

    if l1:
        while l1:
            helper.nxt = Node(l1.data)
            l1 = l1.nxt
            helper = helper.nxt
    if l2:
        while l2:
            helper.nxt = Node(l2.data)
            l2 = l2.nxt
            helper = helper.nxt

    return merged.nxt


class Queue:
    def __init__(self):
        self.stack = []
        self.rstack = []

    def enqueue(self, val):
        self.rstack.insert(0, val)
        self.stack = self.rstack[::-1]
        return
    def dequeue(self):
        val = self.stack.pop(0)
        self.rstack.pop()
        return val
    def peek(self):
        print(self.stack[0])
        return


def stack_queue():
    queries = input('queries> ')
    i = 0
    q = Queue()

    while i < int(queries):
        operation = input('operation> ')
        operation = operation.split(' ')
        if '1' in operation:
            q.enqueue(int(operation[1]))
        elif '2' in operation:
            q.dequeue()
        else:
            q.peek()
        i += 1

def balanced_parens(s):
    if len(s) % 2 != 0:
        return "NO"
    
    stack = []
    obrackets = set(('(', '{', '['))
    s_len = len(s)

    for i in range(s_len):
        if s[i] in obrackets:
            stack.append(s[i])
        else:
            if len(stack) == 0 and i < s_len:
                return "NO"
            top = stack.pop()
            if s[i] == ')' and top != '(':
                return "NO"
            elif s[i] == '}' and top != '{':
                return "NO"
            elif s[i] == ']' and top != '[':
                return "NO"
            
    if len(stack) == 0:
        return "YES"
    return "NO"

from collections import defaultdict

def plus_minus(nums):
    counts = defaultdict(int)

    for num in nums:
        if num > 0:
            counts[1] += 1
        elif num < 0:
            counts[-1] += 1
        else:
            counts[0] += 1

    freqs = list(map(lambda x: x / len(nums), counts.values()))
    return freqs

def min_max_sum(nums):
    nums = sorted(nums)
    min = sum(nums[:4])
    max = sum(nums[-4:])
    return min, max

def convert_time(t):
    result = ""
    if t[-2:] == "AM" and t[:2] == "12":
        result = "00" + t[2:len(t)-2]
    elif t[-2:] == "AM":
        result = t[:len(t)-2]
    elif t[-2:] == "PM" and t[:2] == "12":
        result = t[:len(t)-2]
    else:
        result = str(12 + int(t[:2])) + t[2:len(t)-2]

    return result

def sparse_arrays(vals, query):
    freqs = defaultdict(int)
    for val in vals:
        freqs[val] += 1

    result = []
    for q in query:
        if q in freqs:
            result.append(freqs[q])
        else:
            result.append(0)
    return result

def lonely_int(nums):
    freqs = defaultdict(int)
    result = None
    for num in nums:
        freqs[num] += 1

    for key in freqs.keys():
        if freqs[key] == 1:
            result = key
            break
    return result

def diag_diff(matrix):
    i, td, bu = 0, 0, 0
    for j in range(len(matrix)-1, -1, -1):
        td += matrix[i][i]
        bu += matrix[j][i]
        i += 1
    return abs(td-bu)
