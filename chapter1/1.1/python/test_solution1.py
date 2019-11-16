import pytest

from solution1 import uniq_chars


@pytest.mark.parametrize(
    "input_str,expected",
    [("", True), ("abc", True), ("123", True), ("aab", False), ("113", False)],
)
def test_uniq_chars(input_str: str, expected: bool):
    assert uniq_chars(input_str) == expected
