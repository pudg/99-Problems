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