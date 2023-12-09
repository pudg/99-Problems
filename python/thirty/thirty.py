def bfs(root, target):
    if root == None:
        return
    queue = [root]
    while queue:
        curr = queue.pop()
        if curr.left != None:
            queue.append(curr.left)
        if curr.right != None:
            queue.append(curr.right)
        queue = queue[1:]
        if curr.val == target:
            return True
        
    return False

def dfs(root, target):
    if root is None:
        return False
    
    stack = [root]
    while stack:
        curr = stack.pop(0)
        if curr.val == target:
            return True
        if curr.left is not None:
            stack.append(curr.left)
        if curr.right is not None:
            stack.append(curr.right)
    return False

def min_height(root):
    if root is None:
        return 0
    
    depth = 1
    queue = [root]

    while queue:
        for _ in range(len(queue)):
            curr = queue.pop(0)
            if curr.left is None and curr.right is None:
                return depth
            if curr.left is not None:
                queue.append(curr.left)
            if curr.right is not None:
                queue.append(curr.right)
        depth += 1
    return depth

class Node():
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def add_two_numbers(l1, l2):
    carry = 0
    dummy = Node()
    helper = dummy
    while l1 or l2 or carry != 0:
        l1_val = l1.val if l1 else 0
        l2_val = l2.val if l2 else 0
        curr_sum = l1_val + l2_val + carry
        val = curr_sum % 10
        carry = carry // 10
        helper.next = Node(val)
        helper = helper.next
        l1 = l1.next if l1 else None
        l2 = l2.next if l2 else None

    return dummy.next

def path_sum(root, target_sum):
    if root is None:
        return False
    stack = [(root, root.val)]

    while stack:
        curr = stack.pop(0)
        if curr[0].left is None and curr[0].right is None:
            if curr[1] == target_sum:
                return True
        else:
            if curr[0].left is not None:
                stack.append((curr[0].left, curr[1] + curr[0].left.val))
            if curr[0].right is not None:
                stack.append((curr[0].right, curr[1] + curr[0].right.val))

    return False

def preorder_traversal(root):
    if root is None:
        return []
    
    values, stack = [], [root]
    while stack:
        curr = stack.pop()
        values.append(curr.val)

        if curr.right is not None:
            stack.append(curr.right)
        if curr.left is not None:
            stack.append(curr.left)
        
    return values

def postorder_traversal(root):
    if root is None:
        return []
    
    values, stack = [], [root]
    while stack:
        curr = stack.pop()
        values.append(curr.val)
        if curr.left is not None:
            stack.append(curr.left)
        if curr.right is not None:
            stack.append(curr.right)
            
    return values[::-1]

def ll_cycle(head):
    if head is None:
        return False
    visited = {}
    curr = head

    while curr:
        if curr in visited:
            return True
        visited[curr] = True
        curr = curr.next

    return False

def ll_intersection(headA, headB):
    if headA is None or headB is None:
        return None
    
    memory = set()
    helperA = headA
    while helperA:
        memory.add(helperA)
        helperA = helperA.next

    helperB = headB
    while helperB:
        if helperB in memory:
            return helperB
        helperB = helperB.next

    return None


def longest_substring(s):
    if s == "":
        return 0
    
    memory = {}
    max_len, lhs = 0, 0

    for rhs in range(len(s)):
        if s[rhs] in memory and memory[s[rhs]] >= lhs:
            lhs = memory[s[rhs]] + 1
        else:
            max_len = max(max_len, rhs - lhs + 1)
        memory[s[rhs]] = rhs
    
    return max_len