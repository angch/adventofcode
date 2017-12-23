#!/usr/bin/env python3
# %%
import logging
import functools
import collections

###############################################################################


#
# list of operations, each operation mutates cpu state in-place
#
def op_set(x, y, cpu, id_):
    try:
        cpu[id_][x] = int(y)
    except ValueError:
        cpu[id_][x] = cpu[id_][y]
    cpu[id_]['_pc'] += 1


def op_add(x, y, cpu, id_):
    try:
        cpu[id_][x] += int(y)
    except ValueError:
        cpu[id_][x] += cpu[id_][y]
    cpu[id_]['_pc'] += 1


def op_mul(x, y, cpu, id_):
    try:
        cpu[id_][x] *= int(y)
    except ValueError:
        cpu[id_][x] *= cpu[id_][y]
    cpu[id_]['_pc'] += 1


def op_mod(x, y, cpu, id_):
    try:
        cpu[id_][x] %= int(y)
    except ValueError:
        cpu[id_][x] %= cpu[id_][y]
    cpu[id_]['_pc'] += 1


def op_snd(x, cpu, id_):
    # remember the frequency
    try:
        value = int(x)
    except ValueError:
        value = cpu[id_][x]

    cpu[id_]['_snd'] = value
    cpu[id_]['_pc'] += 1

    if cpu[id_]['_multicore']:
        # send value to another program
        # HACK: this code assumes that there is only 2 valid program... but
        #       since you cannot specify which program to send values to,
        #       therefore it is still technically correct since there is only
        #       2 available program.
        cpu[1 - id_]['_channel'].append(value)

        # increment number of times program has sent a value
        cpu[id_]['_snd_count'] += 1


def op_rcv(x, cpu, id_):
    value = cpu[id_].get(x)
    if value:
        # recover the frequency
        cpu[id_]['_rcv'] = cpu[id_]['_snd']

    # only allow receive when multicore mode is enabled
    if not cpu[id_]['_multicore']:
        cpu[id_]['_pc'] += 1
    else:
        try:
            # do we have anything to receive?
            value = cpu[id_]['_channel'].popleft()
        except IndexError:
            # wait for value and signal potential deadlock
            cpu[id_]['_rcv_wait'] = 1
        else:
            # store into register
            cpu[id_][x] = value

            # reset deadlock event
            cpu[id_]['_rcv_wait'] = 0

            # advance program counter
            cpu[id_]['_pc'] += 1


def op_jgz(x, y, cpu, id_):
    try:
        value = int(x)
    except ValueError:
        value = cpu[id_][x]

    if value > 0:
        # jump
        try:
            cpu[id_]['_pc'] += int(y)
        except ValueError:
            cpu[id_]['_pc'] += cpu[id_][y]
    else:
        cpu[id_]['_pc'] += 1


def compile(data):
    """Parse the input into list of curried op_* functions to avoid cost
       of parsing over and over again."""
    instruction_map = {
        'snd': op_snd,
        'set': op_set,
        'add': op_add,
        'mul': op_mul,
        'mod': op_mod,
        'rcv': op_rcv,
        'jgz': op_jgz,
    }

    for pos, instruction in enumerate(data.strip().splitlines()):
        token = instruction.split()
        operand = token[0]

        try:
            f = instruction_map[operand]
        except KeyError:
            # invalid instruction
            raise ValueError("Invalid instruction '%s' at position '%d'." %
                             (instruction, pos))

        yield functools.partial(f, *token[1:])


###############################################################################


def part1(data):
    # compile the input into list of bytecodes
    bytecodes = list(compile(data))
    # from pprint import pprint
    # pprint(bytecodes)

    # execute the program
    cpu = {0: collections.defaultdict(lambda: 0)}
    for i in range(len(bytecodes)**2):
        try:
            f = bytecodes[cpu[0]['_pc']]
        except IndexError:
            # segmentation fault, halt!
            break
        else:
            # execute the bytecode
            logging.debug("[%d] %r" % (i, f))
            f(cpu, 0)
            logging.debug("[%d] cpu = %r" % (i, cpu))

            # did we recover anything?
            rcv = cpu[0].get('_rcv')
            if rcv:
                return rcv

    # program ended without recovering any frequency
    raise ValueError("No frequency recovered!")


def part2(data):
    # compile the input into list of bytecodes
    bytecodes = list(compile(data))

    # initialize the cpu
    cpu = {
        0: collections.defaultdict(lambda: 0),
        1: collections.defaultdict(lambda: 0)
    }
    cpu[1]['_channel'] = collections.deque()
    cpu[0]['_channel'] = collections.deque()
    cpu[0]['_multicore'] = True
    cpu[1]['_multicore'] = True
    cpu[0]['p'] = 0
    cpu[1]['p'] = 1

    # execute the program
    for i in range(len(bytecodes)**4):
        for id_ in range(len(cpu)):
            if cpu[id_]['_hlt']:
                # core is halted, dont execute anything!
                continue

            try:
                f = bytecodes[cpu[id_]['_pc']]
            except IndexError:
                # segmentation fault, halt!
                cpu[id_]['_hlt'] = True
                continue
            else:
                # execute the bytecode
                logging.debug("[%d][%d] %r" % (i, id_, f))
                f(cpu, id_)
                logging.debug("[%d][%d] cpu = %r" % (i, id_, cpu[id_]))

        # did deadlock occur or both cores halted?
        if (cpu[0]['_rcv_wait'] or cpu[0]['_hlt']) and (cpu[1]['_rcv_wait']
                                                        or cpu[1]['_hlt']):
            break
    else:
        # program ended without recovering any frequency
        raise ValueError("End of bytecode!")

    # program ended without recovering any frequency
    from pprint import pformat
    logging.debug("Final cpu state = %s" % pformat(cpu))
    return cpu[1]['_snd_count']


###############################################################################

if __name__ == '__main__':
    import argparse
    import time

    parser = argparse.ArgumentParser(description='AoC 2017-18 Solution')
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
