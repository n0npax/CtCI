from collections import Counter


def is_string_permutation(s1: str, s2: str) -> bool:
    return Counter(s1) == Counter(s2)
