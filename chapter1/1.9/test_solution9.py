import pytest

from solution9 import is_rotation


@pytest.mark.parametrize(
    "input_str1,input_str2,expected",
    [("llama", "lamal", True), ("abcdef", "cdefab", True), ("bar", "foo", False)],
)
def test_is_rotation(input_str1: str, input_str2: str, expected: str):
    assert is_rotation(input_str1, input_str2) == expected
