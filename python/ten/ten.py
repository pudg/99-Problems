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
    l1 = Node(2)
    l2 = Node(1)
    l1.nxt = Node(3)
    l2.nxt = Node(4)
    l2.nxt.nxt = Node(5)

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

