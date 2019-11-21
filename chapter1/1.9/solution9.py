from typing import List


def is_rotation(s1: str, s2: str) -> bool:
    if len(s1) != len(s2):
        return False
    return is_substr(s1, s2 + s2)


def is_substr(s1: str, s2: str):
    return s1 in s2
