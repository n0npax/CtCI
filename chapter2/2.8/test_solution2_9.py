import pytest

from solution2_9 import detect_cycle


def test_detect_cycle():
    class Node:
        def __init__(self, val=None):
            self.next = None
            self.val = val

    head = Node()
    tail = head
    my_cycle = head
    for i in range(10):
        n = Node(i)
        tail.next = n
        tail = n
    tail.next = head

    new_head = Node()
    new_tail = new_head
    for _ in range(100, 110):
        n = Node(i)
        new_head.next = n
        new_tail = n
    new_tail.next = head

    detected_cycle = detect_cycle(new_head)
    assert detected_cycle.val == my_cycle.val
