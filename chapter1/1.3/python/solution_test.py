import pytest

from solution import urlify


@pytest.mark.parametrize(
    "input_str,input_len,expected_str",
    [
        ("", 0, ""),
        ("Mr John Smith    ", 13, "Mr%20John%20Smith"),
        ("Mr Jack Smith        ", 13, "Mr%20Jack%20Smith    "),
        ("aa bb  ", 5, "aa%20bb"),
        ("cc aa bb    ", 8, "cc%20aa%20bb"),
    ],
)
def test_uniq_chars_bytearray(input_str: str, input_len: int, expected_str: str):
    assert urlify(bytearray(input_str, "utf-8"), input_len) == bytearray(
        expected_str, "utf-8"
    )
