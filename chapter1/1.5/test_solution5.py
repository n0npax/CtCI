import pytest

from solution5 import one_away


@pytest.mark.parametrize(
    "input_str1,input_str2,expected",
    [
        ("", "", True),
        ("aa", "aa", True),
        ("aba", "aaa", True),
        ("aaba", "aaa", True),
        ("aaa", "aaba", True),
        ("abc", "aaa", False),
        ("aaaaa", "aaa", False),
    ],
)
def test_one_away(input_str1: str, input_str2: str, expected: tuple):
    assert one_away(input_str1, input_str2) == expected
