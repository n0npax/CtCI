def compression(s: str) -> bool:
    cnt, prev_char, letters_bank = 0, "", []
    for c in s:
        if c == prev_char:
            cnt += 1
            continue
        if prev_char:
            letters_bank.append(f"{prev_char}{cnt}")
        cnt, prev_char = 1, c

    letters_bank.append(f"{prev_char}{cnt}")
    compressed = "".join(letters_bank)
    if len(compressed) > len(s):
        return s
    return compressed
