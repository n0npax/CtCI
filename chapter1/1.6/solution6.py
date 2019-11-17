def compression(s: str) -> bool:
    cnt, prev_char, compressed = 0, None, ""
    for c in s:
        if c == prev_char:
            cnt += 1
            continue
        if prev_char:
            compressed += f"{prev_char}{cnt}"
        cnt, prev_char = 1, c
    compressed += f"{prev_char}{cnt}"
    if len(compressed) > len(s):
        return s
    return compressed
