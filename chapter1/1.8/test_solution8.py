import pytest

from solution8 import zero_fun


@pytest.mark.parametrize(
    "input_matrix,expected",
    [
        ([], []),
        ([[1, 2], [3, 4]], [[1, 2], [3, 4]]),
        (
            [[1, 2, 3, 4], [5, 6, 7, 8], [9, 4, 4, 4], [4, 4, 4, 0]],
            [[1, 2, 3, 0], [5, 6, 7, 0], [9, 4, 4, 0], [0, 0, 0, 0]],
        ),
        (
            [[1, 3, 4], [1, 5, 8], [2, 4, 4], [1, 7, 0]],
            [[1, 3, 0], [1, 5, 0], [2, 4, 0], [0, 0, 0]],
        ),
    ],
)
def test_compression(input_matrix: list, expected: list):
    assert zero_fun(input_matrix) == expected
