import pytest

from itertools import permutations

from solution import palindrome_permutations, custom_permutations


@pytest.mark.parametrize(
    "input_str,expected",
    [
        ("", set([""])),
        ("aa", set(["aa"])),
        ("aba", set(["aab", "aba", "baa"])),
        ("abc", set()),
    ],
)
def test_palindrome_permutations(input_str: str, expected: tuple):
    assert palindrome_permutations(input_str) == expected


@pytest.mark.parametrize("input_str", [("abc",), ("abad",), ("llama",)])
def test_custom_permutations(input_str: str):
    custom_permutations_set = set("".join(x) for x in custom_permutations(input_str))
    standard_permutations_set = set("".join(x) for x in permutations(input_str))
    assert standard_permutations_set == custom_permutations_set
