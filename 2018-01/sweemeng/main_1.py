def main():
    current = 0
    input_file = open("./input_file")
    for item in input_file:

        if len(item) == 1:
            break
        ops = item[0]
        value = item[1:]
        value = int(value)
        if ops == "+":
            current = current + value
        else:
            current = current - value
    print(current)


if __name__ == "__main__":
    main()
