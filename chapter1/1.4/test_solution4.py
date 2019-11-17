from itertools import permutations

import pytest

from solution4 import custom_permutations, palindrome_permutations


@pytest.mark.parametrize(
    "input_str,expected",
    [
        ("", set([""])),
        ("aa", set(["aa"])),
        ("aba", set(["aab", "aba", "baa"])),
        (
            "aabbc",
            set(
                [
                    "acbba",
                    "abcab",
                    "cbbaa",
                    "babca",
                    "aacbb",
                    "bacab",
                    "ababc",
                    "baabc",
                    "aabcb",
                    "bacba",
                    "abcba",
                    "caabb",
                    "cbaba",
                    "babac",
                    "bbaca",
                    "bcaba",
                    "abbac",
                    "acbab",
                    "bbaac",
                    "cbaab",
                    "bcaab",
                    "bbcaa",
                    "abacb",
                    "cabba",
                    "aabbc",
                    "acabb",
                    "abbca",
                    "baacb",
                    "cabab",
                    "bcbaa",
                ]
            ),
        ),
        ("abc", set()),
    ],
)
def test_palindrome_permutations(input_str: str, expected: tuple):
    assert palindrome_permutations(input_str) == expected


@pytest.mark.parametrize("input_str", ["abc", "abad", "llama", "abcde"])
def test_custom_permutations(input_str: str):
    custom_permutations_set = set("".join(x) for x in custom_permutations(input_str))
    standard_permutations_set = set("".join(x) for x in permutations(input_str))
    assert standard_permutations_set == custom_permutations_set
