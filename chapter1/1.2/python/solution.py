from collections import Counter


def is_permutation(s1: str, s2: str) -> bool:
    return Counter(s1) == Counter(s2)
