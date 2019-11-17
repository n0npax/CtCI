from collections import Counter
from itertools import permutations
from typing import Generator


def palindrome_permutations(s: str) -> set:
    # returns empty set if word is not palindrome permutation
    # if word is palidrom permutation returns all possible permutations

    # word is palindrome permutation if all letter in counter are even. Just 1 letter can be odd
    c = Counter(s)
    odd_cnt = 0
    for _, v in c.items():
        if v % 2:
            odd_cnt += 1
    if odd_cnt > 1:
        return set()  # not a palindrome permutation
    return set("".join(x) for x in permutations(s))


# assuming using permutations from standard lib is not allowed
def custom_permutations(s: str) -> Generator[str, None, None]:
    if len(s) == 1:
        yield s
    for i in range(len(s)):
        sl = list(s)
        if sl[i] != sl[0]:
            sl[i], sl[0] = sl[0], sl[i]
        sl = "".join(sl)
        for ys in custom_permutations(sl[1:]):
            yield sl[0] + ys
