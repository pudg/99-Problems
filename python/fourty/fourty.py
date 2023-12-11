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

def reverse_ll(head):
    prev, curr = None, head

    while curr:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt
    return prev

class ListNode():
    def __init__(self, val=None, nxt=None):
        self.val = val
        self.next = nxt

def merge_lists(list1, list2):
    dummy = ListNode()
    helper = dummy

    while list1 or list2:
        if list1 is None:
            helper.next = ListNode(list2.val)
            helper, list2 = helper.next, list2.next
        elif list2 is None:
            helper.next = ListNode(list1.val)
            helper, list1 = helper.next, list1.next
        else:
            if list1.val <= list2.val:
                helper.next = ListNode(list1.val)
                helper, list1 = helper.next, list1.next
            else:
                helper.next = ListNode(list2.val)
                helper, list2 = helper.next, list2.next
    return dummy.next