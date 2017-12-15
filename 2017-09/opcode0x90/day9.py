#!/usr/bin/env python3
import logging

# %%


def part1(data):
    it = iter(data.strip())
    depth = 0
    stack = []

    for c in it:
        if c == '!':
            # skip the next char
            next(it)
        elif c == '{':
            depth += 1
            stack.append(depth)
        elif c == '}':
            depth -= 1
        elif c == '<':
            # begin special loop to consume the iterator until we find the next valid '>' char
            for d in it:
                if d == '!':
                    # skip the next char
                    next(it)
                elif d == '>':
                    break
            else:
                # this is not supposed to happen
                raise ValueError("EOF reached while finding the end of <>")

    return sum(stack)


def part2(data):
    it = iter(data.strip())
    count = 0

    for c in it:
        if c == '!':
            # skip the next char
            next(it)
        # elif c == '{':
        #     depth += 1
        #     stack.append(depth)
        # elif c == '}':
        #     depth -= 1
        elif c == '<':
            # begin special loop to consume the iterator until we find the next valid '>' char
            for d in it:
                if d == '!':
                    # skip the next char
                    next(it)
                    continue
                elif d == '>':
                    break
                count += 1
            else:
                # this is not supposed to happen
                raise ValueError("EOF reached while finding the end of <>")

    return count


###############################################################################

if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser(description='AoC 2017-09 Solution')
    parser.add_argument(
        '--input',
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

    with open(args.input) as f:
        data = f.read()

    for line in data.splitlines():
        if not args.no_part1:
            print("Part 1: %r" % part1(line))
        if not args.no_part2:
            print("Part 2: %r" % part2(line))
