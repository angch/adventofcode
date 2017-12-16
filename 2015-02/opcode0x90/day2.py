#!/usr/bin/env python3
# %%
import logging
import itertools


def part1(data):
    orders = [(int(d) for d in line.split('x')) for line in data.splitlines()
              if line.strip()]

    total = 0
    for l, w, h in orders:
        area = (2 * l * w) + (2 * w * h) + (2 * h * l)
        slack = min(a * b for a, b in itertools.combinations((l, w, h), 2))

        total += (area + slack)

    return total


###############################################################################


def part2(data):
    orders = [(int(d) for d in line.split('x')) for line in data.splitlines()
              if line.strip()]

    total = 0
    for l, w, h in orders:
        bow = (l * w * h)
        ribbon = min(
            (2 * a + 2 * b) for a, b in itertools.combinations((l, w, h), 2))

        total += (bow + ribbon)

    return total


###############################################################################

if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser(description='AoC 2015-02 Solution')
    parser.add_argument(
        'INPUT',
        nargs='?',
        default="input.txt",
        help='Input file to run the solution with.')
    parser.add_argument(
        '-v', '--verbose', action="store_true", help='Turn on verbose logging.')
    parser.add_argument(
        '-l', '--line', action="store_true", help='Parse input line by line.')
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

    if not args.line:
        if not args.no_part1:
            print("Part 1: %r" % part1(data))
        if not args.no_part2:
            print("Part 2: %r" % part2(data))
    else:
        if not args.no_part1:
            for line in data.splitlines():
                print("Part 1: %r = %r" % (line, part1(line)))
        if not args.no_part2:
            for line in data.splitlines():
                print("Part 2: %r = %r" % (line, part2(line)))
