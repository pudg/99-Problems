def bfs(root, target):
    if root == None:
        return
    queue = []
    queue.append(root)
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