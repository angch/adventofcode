def main():
    f = open("input")
    count = {"two": 0, "three": 0}
    for line in f:
        wc = {}
        for letter in line:
            if letter in wc:
                wc[letter] += 1
            else:
                wc[letter] = 1
        if 2 in wc.values():
            count["two"] += 1

        if 3 in wc.values():
            count["three"] += 1

    return count["two"] * count["three"]


if __name__ == "__main__":
    print(main())
