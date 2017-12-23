#!/usr/bin/env python3
# %%
import logging
import string
import functools

###############################################################################


#
# list of operations, all mutations will be performed in-place
#
def op_spin(x, programs):
    buf = programs.copy()
    length = len(programs)

    for i in range(length):
        programs[(x + i) % length] = buf[i]


def op_exchange(a, b, programs):
    programs[a], programs[b] = programs[b], programs[a]


def op_partner(a, b, programs):
    i = programs.index(a)
    j = programs.index(b)
    programs[i], programs[j] = programs[j], programs[i]


def compile(data):
    """Parse the input into list of curried op_* functions to avoid cost
        of parsing over and over again."""
    instruction_map = {
        's': (op_spin, int),
        'x': (op_exchange, int),
        'p': (op_partner, ord)
    }

    for pos, instruction in enumerate(data.strip().split(',')):
        prefix = instruction[0]

        try:
            f, map_f = instruction_map[prefix]
        except KeyError:
            # invalid instruction
            raise ValueError("Invalid instruction '%s' at position '%d'." %
                             (instruction, pos))

        imm = instruction[1:].split('/')
        yield functools.partial(f, *map(map_f, imm))


###############################################################################


def part1(data, size=16):
    # compile the instructions
    instructions = list(compile(data))

    # execute the instructions
    programs = bytearray(string.ascii_letters[:size], 'ascii')
    for pos, op in enumerate(instructions):
        op(programs)
        logging.debug("[%d] %r = %s" % (pos, op, programs))

    return programs.decode('ascii')


def part2(data, size=16):
    # compile the instructions
    instructions = tuple(compile(data))

    # execute the instructions
    cache = {}
    cache_hit = 0
    cache_miss = 0
    cycle = set()

    programs = bytearray(string.ascii_letters[:size], 'ascii')
    i = 1000000000
    try:
        while i > 0:
            programs_ = bytes(programs)

            # is the result cached?
            result = cache.get(programs_)
            if result:
                # reuse result from cache, no need to compute again
                programs = result
                cache_hit += 1

                # check for cycle
                if programs_ not in cycle:
                    cycle.add(programs_)
                else:
                    # cycle detected, perform fast skipping
                    # skip all multiples of cycle length
                    i = i % len(cycle)
                    cycle = set()
                    logging.debug("Cycles detected, performing fast skipping!")
            else:
                cache_miss += 1
                cycle = set()

                # execute the instructions
                for op in instructions:
                    op(programs)

                # cache the result
                cache[programs_] = programs.copy()

            i -= 1
            logging.debug("[%d] Cache miss-hit: %d - %d" % (i, cache_miss,
                                                            cache_hit))
    except KeyboardInterrupt:
        pass

    return programs.decode('ascii')


###############################################################################

if __name__ == '__main__':
    import argparse
    import time

    parser = argparse.ArgumentParser(description='AoC 2017-16 Solution')
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
        '-s',
        '--size',
        type=int,
        default=16,
        help='Specify length of program.')
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
        result = part1(data, args.size)
        t = time.clock() - start
        print("Part 1: [%.3f] %r" % (t, result))
    if not args.no_part2:
        start = time.clock()
        result = part2(data, args.size)
        t = time.clock() - start
        print("Part 2: [%.3f] %r" % (t, result))
