#!/usr/bin/env python3
# %%
import logging
import itertools


def part1(data):
    x = y = 0
    visited = {(0, 0)}

    for d in data.strip():
        dx, dy = {'^': (0, 1), 'v': (0, -1), '>': (1, 0), '<': (-1, 0)}[d]

        x += dx
        y += dy
        visited.add((x, y))

    return len(visited)


###############################################################################


def part2(data):
    # https://docs.python.org/3/library/itertools.html#itertools-recipes
    def grouper(iterable, n, fillvalue=None):
        "Collect data into fixed-length chunks or blocks"
        # grouper('ABCDEFG', 3, 'x') --> ABC DEF Gxx"
        args = [iter(iterable)] * n
        return itertools.zip_longest(*args, fillvalue=fillvalue)

    # sanity check
    if len(data.strip()) % 2:
        # reject odd inputs
        return -1

    ax = ay = bx = by = 0
    visited = {(0, 0)}
    direction = {'^': (0, 1), 'v': (0, -1), '>': (1, 0), '<': (-1, 0)}

    for a, b in grouper(data.strip(), 2):
        dx, dy = direction[a]
        ax += dx
        ay += dy
        visited.add((ax, ay))

        dx, dy = direction[b]
        bx += dx
        by += dy
        visited.add((bx, by))

    return len(visited)


###############################################################################

if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser(description='AoC 2015-03 Solution')
    parser.add_argument(
        'INPUT',
        nargs='?',
        default="input.txt",
        help='Input file to run the solution with.')
    parser.add_argument(
        '-v', '--verbose', action="store_true", help='Turn on verbose logging.')
    parser.add_argument(
        '-1',
        '--no-part1',
        action="store_true",
        help='Exclude Part 1 solution from run.')
    parser.add_argument(
        '-2',
        '--no-part2',
        action="store_true",
        help='Exclude Part 2 solution from run.')
    args = parser.parse_args()

    if args.verbose:
        logging.basicConfig(level=logging.DEBUG)

    with open(args.INPUT) as f:
        data = f.read()

    if not args.no_part1:
        for i, line in enumerate(data.splitlines(), start=1):
            print("Part 1: line %d = %r" % (i, part1(line)))
    if not args.no_part2:
        for i, line in enumerate(data.splitlines(), start=1):
            print("Part 2: line %d = %r" % (i, part2(line)))
