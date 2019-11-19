import math
from typing import List


def rotate(img: List[List[List[bytes]]]) -> List[List[List[bytes]]]:
    n = len(img)
    n2 = math.ceil(n / 2)
    n3 = math.floor(n / 2)
    for i in range(n2):
        for j in range(n3):
            a = img[i][j]
            b = img[n - 1 - j][i]
            c = img[n - 1 - i][n - 1 - j]
            d = img[j][n - 1 - i]

            img[i][j] = b
            img[n - 1 - j][i] = c
            img[n - 1 - i][n - 1 - j] = d
            img[j][n - 1 - i] = a
    return img
