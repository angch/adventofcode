from io import IOBase


def counter(input_list):
    seen = set()
    current = 0
    seen.add(current)
    found = False

    while not found:
        # Should really check for file,
        input_list.seek(0)

        for i in input_list:

            if len(i) == 1:
                break
            ops = i[0]
            value = i[1:]
            value = int(value)

            if ops == "-":
                current -= value
            else:
                current += value
            if current not in seen:
                seen.add(current)
            else:
                found = True
                break

    return current


if __name__ == "__main__":
    f = open("input_file")
    result = counter(f)
    print(result)

