#!/usr/bin/env python3
import logging
import time
import sys
import functools
import multiprocessing


def adventcoin(key, i=0, difficulty=5):
    import hashlib
    digest = hashlib.md5(("%s%d" % (key, i)).encode('utf8')).hexdigest()
    return (digest.startswith('0' * difficulty), i, digest)


###############################################################################


def adventofcode_4(data, difficulty=5):
    # spawn as many processes as available cpu cores
    with multiprocessing.Pool(processes=multiprocessing.cpu_count()) as pool:
        f = functools.partial(adventcoin, data, difficulty=difficulty)

        # naive implementation of producer-consumer design
        for bingo, i, digest in pool.imap_unordered(f, range(sys.maxsize)):
            if bingo:
                return (i, digest)


###############################################################################

if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser(
        description='AoC 2015-04 Bad Parallel Example')
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
