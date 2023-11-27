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