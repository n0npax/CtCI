import pytest

from solution7 import rotate


@pytest.mark.parametrize(
    "input,expected",
    [
        ([], []),
        ([[1, 2], [3, 4]], [[3, 1], [4, 2]]),
        ([[1, 2, 3], [4, 5, 6], [7, 8, 9]], [[7, 4, 1], [8, 5, 2], [9, 6, 3]]),
        (
            [[1, 2, 3, 4], [5, 6, 7, 8], [9, "A", "B", "C"], ["D", "E", "F", "0"]],
            [["D", 9, 5, 1], ["E", "A", 6, 2], ["F", "B", 7, 3], ["0", "C", 8, 4]],
        ),
    ],
)
def test_rotate(input: list, expected: list):
    assert rotate(input) == expected
