from Levenshtein import distance


def main():
    f = open("input")

    result = []
    output = []

    for line in f:
        result.append(line)

    for line in result:
        for check in result:
            if line == check:
                continue

            if distance(line, check) == 1:
                output.append(line)

    temp = []
    for i in range(len(output[0])):
        if output[0][i] == output[1][i]:
            temp.append(output[0][i])
    print("".join(temp))


if __name__ == "__main__":
    main()
