def main():
    combined_rect = set()
    overlapped_rect = set()

    f = open("input")
    for line in f:
        id_, x, y, size_x, size_y = parse(line)
        for coord in generate_coord(x, y, size_x, size_y):
            if coord in combined_rect:
                overlapped_rect.add(coord)
            combined_rect.add(coord)

    print(len(overlapped_rect))


def parse(line):
    id_, _, start_point, size = line.split()
    x, y = start_point.split(",")
    y = y[:-1]
    size_x, size_y = size.split("x")
    return id_, int(x), int(y), int(size_x), int(size_y)


def generate_coord(x, y, size_x, size_y):
    for x_ in range(x, x + size_x):
        for y_ in range(y, y + size_y):
            yield (x_, y_)


if __name__ == "__main__":
    main()
