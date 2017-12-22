#!/usr/bin/env python3
# %%
import logging


def part1(data):
    a, b = data
    score = 0

    for _ in range(40000000):
        a = (a * 16807) % 2147483647
        b = (b * 48271) % 2147483647

        if (a & 0xffff) == (b & 0xffff):
            score += 1

    return score


def part2(data):
    def gen(x, factor, multiples):
        while True:
            x = (x * factor) % 2147483647
            if not (x % multiples):
                yield x & 0xffff

    a, b = data

    a = gen(a, 16807, 4)
    b = gen(b, 48271, 8)
    return sum(next(a) == next(b) for _ in range(5000000))


def part2_numba(data):
    a, b = data
    score = 0

    for _ in range(5000000):
        while True:
            a = (a * 16807) % 2147483647
            if not (a % 4):
                break
        while True:
            b = (b * 48271) % 2147483647
            if not (b % 8):
                break

        if (a & 0xffff) == (b & 0xffff):
            score += 1

    return score


###############################################################################

if __name__ == '__main__':
    import argparse
    import time

    parser = argparse.ArgumentParser(description='AoC 2017-15 Solution')
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
        a, b = [int(line.split()[-1]) for line in data.splitlines()]
        data_ = (a, b)

    if not args.no_part1:
        start = time.clock()
        result = part1(data_)
        t = time.clock() - start
        print("Part 1: [%.3f] %r = %r" % (t, data_, result))
    if not args.no_part2:
        start = time.clock()
        result = part2(data_)
        t = time.clock() - start
        print("Part 2: [%.3f] %r = %r" % (t, data_, result))

        try:
            import numba
            part2_numba = numba.jit(
                part2_numba, nopython=True, nogil=True, parallel=True)

            start = time.clock()
            result = part2_numba(data_)
            t = time.clock() - start
            print("Part 2 - Numba: [%.3f] %r = %r" % (t, data_, result))
        except ImportError:
            pass
