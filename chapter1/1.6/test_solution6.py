import pytest

from solution6 import compression


@pytest.mark.parametrize(
    "input_str,expected",
    [
        ("", ""),
        ("aaa", "a3"),
        ("abc", "abc"),
        ("aaabbb", "a3b3"),
        ("aaaAAAaaa", "a3A3a3"),
        ("abcdefg", "abcdefg"),
    ],
)
def test_compression(input_str: str, expected: str):
    assert compression(input_str) == expected
