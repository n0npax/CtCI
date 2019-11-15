import pytest

from solution import is_string_permutation


@pytest.mark.parametrize(
    "input_str1,input_str2,expected",
    [
        ("", "", True),
        ("abc", "abc", True),
        ("123", "123", True),
        ("aab", "ccd", False),
        ("113", "445", False),
    ],
)
def test_uniq_chars(input_str1: str, input_str2, expected: bool):
    assert is_string_permutation(input_str1, input_str2) == expected
