#!/usr/bin/env python3
# %%
import logging
import itertools
import functools
import operator


def part1(data, length=256):
    rope = list(range(length))
    tape = [int(x) for x in data.split(',')]
    skip = 0
    i = 0

    for length in tape:
        rope_ = itertools.cycle(enumerate(rope))

        # select the sublist to be reversed
        logging.debug("[%d,%d] length: %r" % (i, skip, length))
        sublist = list(itertools.islice(rope_, i, i + length))
        logging.debug("[%d,%d] sublist: %r" % (i, skip, sublist))

        # unzip the sublist and reverse the index
        # then re-assign values back to the rope
        index, values = zip(*sublist)
        for j, value in zip(reversed(index), values):
            rope[j] = value

        # advance the index
        logging.debug("[%d,%d] a: %r" % (i, skip, rope))
        i += (length + skip)
        skip += 1

    # return the result of multiplying the first two numbers in the list
    return rope[0] * rope[1]


###############################################################################


def part2(data):
    # https://docs.python.org/3/library/itertools.html#itertools-recipes
    def grouper(iterable, n, fillvalue=None):
        "Collect data into fixed-length chunks or blocks"
        # grouper('ABCDEFG', 3, 'x') --> ABC DEF Gxx"
        args = [iter(iterable)] * n
        return itertools.zip_longest(*args, fillvalue=fillvalue)

    rope = list(range(256))
    tape = [ord(x) for x in data] + [17, 31, 73, 47, 23]
    skip = 0
    i = 0

    for _ in range(64):
        for length in tape:
            rope_ = itertools.cycle(enumerate(rope))

            # select the sublist to be reversed
            sublist = tuple(itertools.islice(rope_, i, i + length))

            # unzip the sublist and reverse the index
            # then re-assign values back to the rope
            index, values = zip(*sublist)
            for j, value in zip(reversed(index), values):
                rope[j] = value

            # advance the index
            logging.debug("[%d,%d] a: %r" % (i, skip, rope))
            i += (length + skip)
            skip += 1

    # return the hexadecimal representation of Knot Hash
    return ''.join(
        hex(c)[2:]
        for c in (functools.reduce(operator.xor, chunk)
                  for chunk in grouper(rope, 16)))


###############################################################################

if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser(description='AoC 2017-10 Solution')
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

    for line in data.splitlines():
        length, _, data_ = line.partition(':')
        if not data_:
            data_ = length
            length = 256

        if not args.no_part1:
            print("Part 1: %r = %r" % (data_, part1(data_, length=int(length))))
        if not args.no_part2:
            print("Part 2: %r = %r" % (data_, part2(data_)))
