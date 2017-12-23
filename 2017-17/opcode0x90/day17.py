#!/usr/bin/env python3
# %%
import logging
import collections

###############################################################################


def part1(data):
    steps = int(data)
    spinlock = collections.deque([0])

    for i in range(1, 2017 + 1):
        # spin to win
        spinlock.rotate(-steps)
        logging.debug("[%d] rotate: %r" % (i, spinlock))

        # insert at position
        spinlock.append(i)
        logging.debug("[%d] insert: %r" % (i, spinlock))

    return spinlock[0]


def part2(data):
    steps = int(data)
    spinlock = collections.deque([0])

    for i in range(1, 50000000 + 1):
        spinlock.rotate(-steps)
        spinlock.append(i)

    return spinlock[spinlock.index(0) + 1]


###############################################################################

if __name__ == '__main__':
    import argparse
    import time

    parser = argparse.ArgumentParser(description='AoC 2017-17 Solution')
    parser.add_argument(
        'INPUT',
        nargs='?',
        default="input.txt",
        help='Input file to run the solution with.')
    parser.add_argument(
        '-v',
        '--verbose',
        action="store_true",
        help='Turn on verbose logging.')
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
        start = time.clock()
        result = part1(data)
        t = time.clock() - start
        print("Part 1: [%.3f] %r" % (t, result))
    if not args.no_part2:
        start = time.clock()
        result = part2(data)
        t = time.clock() - start
        print("Part 2: [%.3f] %r" % (t, result))
