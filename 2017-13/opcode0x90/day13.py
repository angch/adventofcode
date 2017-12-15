#!/usr/bin/env python3
import sys


def part1(data):
    # construct the firewall
    fw = {
        int(layer): int(depth)
        for layer, _, depth in (line.partition(': ')
                                for line in data.strip().splitlines())
    }

    # construct the modulo table
    modtable = {layer: (depth - 1) * 2 for layer, depth in fw.items()}

    # begin the simulation
    path = tuple(fw.keys())
    severity = 0

    for t in path:
        # did we get caught?
        if not (t % modtable[t]):
            # oh noes!
            severity += (t * fw[t])

    return severity


###############################################################################


def part2(data):
    # construct the firewall
    fw = {
        int(layer): int(depth)
        for layer, _, depth in (line.partition(': ')
                                for line in data.strip().splitlines())
    }

    # construct the modulo table
    modtable = {layer: (depth - 1) * 2 for layer, depth in fw.items()}

    # search for solution
    path = tuple(fw.keys())
    for t in range(sys.maxsize):
        # did we bypass firewall?
        if all((t + i) % modtable[i] for i in path):
            return t
    else:
        raise ValueError("Unable to solve problem after %d iterations!" % t)


###############################################################################

if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser(description='AoC 2017-13 Solution')
    parser.add_argument(
        '--input',
        default="input.txt",
        help='Input file to run the solution with.')
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

    with open(args.input) as f:
        data = f.read()

        if not args.no_part1:
            print("Part 1: %r" % part1(data))
        if not args.no_part2:
            print("Part 2: %r" % part2(data))
