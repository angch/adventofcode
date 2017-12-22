#!/usr/bin/env python3
# %%
import logging
import itertools
import functools
import operator


def knot(data):
    # https://docs.python.org/3/library/itertools.html#itertools-recipes
    def grouper(iterable, n, fillvalue=None):
        "Collect data into fixed-length chunks or blocks"
        # grouper('ABCDEFG', 3, 'x') --> ABC DEF Gxx"
        args = [iter(iterable)] * n
        return itertools.zip_longest(*args, fillvalue=fillvalue)

    rope = list(range(256))
    ropelen = len(rope)
    tape = [ord(x) for x in data] + [17, 31, 73, 47, 23]
    skip = 0
    i = 0

    for _ in range(64):
        for length in tape:
            # select the sublist to be reversed
            index = []
            values = []
            for j in (x % ropelen for x in range(i, i + length)):
                index.append(j)
                values.append(rope[j])

            # reverse the index, then re-assign values back to the rope
            for j, value in zip(reversed(index), values):
                rope[j] = value

            # advance the index
            i += (length + skip)
            skip += 1

    # return the hexadecimal representation of Knot Hash
    return ''.join(
        format(c, '08b')
        for c in (functools.reduce(operator.xor, chunk)
                  for chunk in grouper(rope, 16)))


###############################################################################


def part1(data):
    # compute the knot hashes, then count the number of ones
    hash_ = (knot('%s-%d' % (data, i)) for i in range(128))
    return sum(h.count('1') for h in hash_)


def part2(data):
    def floodfill(map_, x, y, find='1', replace='x', count=0):
        """Flood fill is performed recursively by scanning in order of
            left, up, right, down, iterating depth-first."""
        # pre-scan, ensure given coord has valid hit
        if map_[y][x] == ord(find):
            map_[y][x] = ord(replace)
            count += 1

            # pass 1: scan left
            logging.debug("  - [%d][%d] scanning left" % (x, y))
            for i in range(x - 1, -1, -1):
                logging.debug("    - [%d][%d] = %s" % (x, y, map_[y][i]))
                if map_[y][i] == ord(find):
                    map_, count = floodfill(map_, i, y, find, replace, count)
                else:
                    break

            # pass 2: scan up
            logging.debug("  - [%d][%d] scanning up" % (x, y))
            for i in range(y - 1, -1, -1):
                logging.debug("    - [%d][%d] = %s" % (x, y, map_[i][x]))
                if map_[i][x] == ord(find):
                    map_, count = floodfill(map_, x, i, find, replace, count)

                else:
                    break

            # pass 3: scan right
            logging.debug("  - [%d][%d] scanning right" % (x, y))
            for i in range(x + 1, len(map_[y])):
                logging.debug("    - [%d][%d] = %s" % (x, y, map_[y][i]))
                if map_[y][i] == ord(find):
                    map_, count = floodfill(map_, i, y, find, replace, count)

                else:
                    break

            # pass 4: scan down
            logging.debug("  - [%d][%d] scanning down" % (x, y))
            for i in range(y + 1, len(map_)):
                logging.debug("    - [%d][%d] = %s" % (x, y, map_[i][x]))
                if map_[i][x] == ord(find):
                    map_, count = floodfill(map_, x, i, find, replace, count)

                else:
                    break

        logging.debug("=======================")
        return (map_, count)

    # generate the 128x128 map of knot hashes
    knot_map = [
        bytearray(knot('%s-%d' % (data, i)), 'ascii') for i in range(128)
    ]

    # perform flood-fill algorithm to find all regions
    count = 0

    # scan the entire map for regions
    for x, y in ((x, y) for y in range(128) for x in range(128)):
        knot_map, c = floodfill(knot_map, x, y)
        if c:
            count += 1

    return count


###############################################################################

if __name__ == '__main__':
    import argparse
    import time

    parser = argparse.ArgumentParser(description='AoC 2015-14 Solution')
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

    if not args.no_part1:
        for line in data.splitlines():
            start = time.clock()
            result = part1(line)
            t = time.clock() - start
            print("Part 1: [%.3f] %r = %r" % (t, line, result))
    if not args.no_part2:
        for line in data.splitlines():
            start = time.clock()
            result = part2(line)
            t = time.clock() - start
            print("Part 2: [%.3f] %r = %r" % (t, line, result))
