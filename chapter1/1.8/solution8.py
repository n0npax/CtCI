from typing import List


def zero_fun(matrix: List[List[int]]) -> List[List[int]]:
    rows_to_zero, cols_to_zero = set(), set()
    for i, col in enumerate(matrix):
        for j, val in enumerate(col):
            if not val:
                cols_to_zero.add(i)
                rows_to_zero.add(j)
    if rows_to_zero:
        zero_col = [0] * len(matrix[0])
        for col_index in cols_to_zero:
            matrix[col_index] = zero_col
        cols_to_zero = set(range(len(zero_col))) - cols_to_zero
        for row_index in rows_to_zero:
            for col_index in cols_to_zero:
                matrix[col_index][row_index] = 0
    return matrix
