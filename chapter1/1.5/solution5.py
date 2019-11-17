def one_away(s1: str, s2: str) -> bool:
    l1, l2 = len(s1), len(s2)
    if abs(l1 - l2) > 1:
        return False
    if s1 == s2:
        return True
    if l1 > l2:
        s1, s2 = s2, s1
    for i, c in enumerate(s1):
        if c == s2[i]:
            continue
        for i, c in enumerate(reversed(s1[i + 1 :])):
            if c == s2[i]:
                continue
            return False
        return True
    raise Exception("Shouldn't happen")
