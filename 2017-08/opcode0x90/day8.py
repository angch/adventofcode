#!/usr/bin/env python3
import logging
# %%
import operator
from collections import defaultdict


def part1(data):
    program = (line.split() for line in data.strip().splitlines())
    register = defaultdict(lambda: 0)

    for instruction in program:
        # decode the instruction
        reg, op, imm, if_, cmp_reg, cmp_op, cmp_imm = instruction
        assert if_ == 'if'

        # did the comparison succeed?
        cmp_map = {
            '>': operator.gt,
            '<': operator.lt,
            '>=': operator.ge,
            '==': operator.eq,
            '<=': operator.le,
            '!=': operator.ne
        }
        if cmp_map[cmp_op](register[cmp_reg], int(cmp_imm)):
            # execute the instruction
            register[reg] += int(imm) if op == 'inc' else -int(imm)

    logging.debug("cpu: %r" % register)w
    return max(register.values())


###############################################################################


def part2(data):
    program = (line.split() for line in data.strip().splitlines())
    register = defaultdict(lambda: 0)
    max_ = 0

    for instruction in program:
        # decode the instruction
        reg, op, imm, if_, cmp_reg, cmp_op, cmp_imm = instruction
        assert if_ == 'if'

        # did the comparison succeed?
        cmp_map = {
            '>': operator.gt,
            '<': operator.lt,
            '>=': operator.ge,
            '==': operator.eq,
            '<=': operator.le,
            '!=': operator.ne
        }
        if cmp_map[cmp_op](register[cmp_reg], int(cmp_imm)):
            # execute the instruction
            value = int(imm) if op == 'inc' else -int(imm)
            register[reg] += value
            max_ = max([max_, register[reg]])

    logging.debug("cpu: %r" % register)
    return max_


###############################################################################

if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser(description='AoC 2017-08 Solution')
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
        print("Part 1: %r" % part1(data))
    if not args.no_part2:
        print("Part 2: %r" % part2(data))
