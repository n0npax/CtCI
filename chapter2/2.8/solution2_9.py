def detect_cycle(head):
    a, b = head, head
    while True:
        a = a.next.next
        b = b.next
        if a == b:
            break

    c = head
    while True:
        if a == c:
            return a
        c, a = c.next, a.next
