#!/usr/bin/env python3
# %%


def part1(data):
    def parse_data(data):
        for line in data.splitlines():
            if not line:
                continue

            left, _, right = line.partition('->')
            if not right:
                left = line
            node, _, weight = left.partition(' ')
            yield (node, int(weight.strip()[1:-1]),
                   {x.strip()
                    for x in right.split(',') if x.strip()})

    # pass 1: get all the weights
    nodes = {
        key: {
            'weight': weight,
            'childs': childs
        }
        for key, weight, childs in parse_data(data)
    }

    # pass 2: construct the tree
    not_root = set()
    for k, v in nodes.items():
        childs = v['childs']
        if not childs:
            not_root.add(k)
        else:
            not_root.update(childs)
            v['childs'] = {c: nodes[c] for c in childs}

    # find the root node by eliminating all nodes marked not_root
    root = (nodes.keys() - not_root).pop()
    return (root, nodes)


###############################################################################


def part2(data):
    def check_weight(node, key):
        weight = node['weight']

        childs = node['childs']
        if childs:
            for _ in range(2):
                child_weight = {
                    k: check_weight(child, k)
                    for k, child in childs.items()
                }

                # are they balanced?
                if len(set(child_weight.values())) > 1:
                    # unbalanced node detected!
                    # print("[%s] child_weight: %r" % (key, child_weight))

                    max_ = max(child_weight.values())
                    for k, v in child_weight.items():
                        if v == max_:
                            unb_child = k
                            break
                    else:
                        # this is not supposed to happen
                        raise RuntimeError(
                            "Failed to locate unbalanced child node!")

                    w = childs[unb_child]['weight'] - (
                        max_ - min(child_weight.values()))
                    print(
                        "Unbalanced node '%s' found! Its weight should be: %d" %
                        (unb_child, w))

                    # fix the weight and try again
                    childs[unb_child]['weight'] = w
                    continue

                # add to node weight
                weight += sum(child_weight.values())
                break
            else:
                # this is not supposed to happen
                raise ValueError("Unable to balance the tree!")

        return weight

    # re-use the output from part 1
    root, adj = part1(data)
    # from pprint import pprint
    # pprint(adj[root])

    return check_weight(adj[root], root)


###############################################################################

if __name__ == '__main__':
    import argparse

    parser = argparse.ArgumentParser(description='AoC 2017-07 Solution')
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
        help='Exclude Part 1 solution from run.')
    args = parser.parse_args()

    with open(args.input) as f:
        data = f.read()

        if not args.no_part1:
            solution, _ = part1(data)
            print("Part 1: %r" % solution)
        if not args.no_part2:
            part2(data)
            # print("Part 2: %r" % part2(data))
