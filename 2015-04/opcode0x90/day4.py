#!/usr/bin/env python3
import logging
import time
import sys


def adventcoin(key, i=0, difficulty=5):
    import hashlib
    digest = hashlib.md5(("%s%d" % (key, i)).encode('utf8')).hexdigest()
    return (digest.startswith('0' * difficulty), i, digest)


###############################################################################


def adventofcode_4(data, difficulty=5):
    for i in range(sys.maxsize):
        bingo, _, digest = adventcoin(data, i, difficulty)
        if bingo:
            return (i, digest)
    else:
        raise TimeoutError


###############################################################################

if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser(description='AoC 2015-04 Solution')
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

    if not args.no_part1:
        for line in data.splitlines():
            start = time.clock()
            result = adventofcode_4(line, difficulty=5)
            t = time.clock() - start
            print("Part 1: [%.3f] %r = %r" % (t, line, result))
    if not args.no_part2:
        for line in data.splitlines():
            start = time.clock()
            result = adventofcode_4(line, difficulty=6)
            t = time.clock() - start
            print("Part 2: [%.3f] %r = %r" % (t, line, result))
