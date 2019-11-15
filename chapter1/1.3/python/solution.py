from collections import Counter
# bytearray is printable and can substiture cstring for question purposes
# Assumption: sufficent spaces at the end means enough or more. if there is more spaces then needed, spaces at the end are not stripped nor repleaced
def urlify(s: bytearray, l: int) -> bytearray:
    last_index = len(s) - 1
    replace_switch = False
    space = ord(" ")
    escaped_space = [ord("%"), ord("2"), ord("0")]
    spaces_slots = 0
    for i in reversed(range(len(s))):
        if s[i] == space and not replace_switch:
            continue
        # adjust shift if there is more spaces at the EOL then sufficent
        if not replace_switch:
            replace_switch = True
            spaces_slots = (last_index - i) // 2
            last_index -= 2 * (spaces_slots - Counter(s[:l])[space])
        if s[i] == space:
            s[last_index - 2 : last_index + 1] = escaped_space
            last_index -= 3
        else:
            s[last_index] = s[i]
            last_index -= 1
    return s
