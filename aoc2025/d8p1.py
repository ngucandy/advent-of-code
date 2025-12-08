import math
from itertools import combinations

# parse input to create a tuple of x,y,z coordinates for each jbox
jboxes = [tuple(map(int, line.strip().split(','))) for line in open(0)]
pairs = list(combinations(jboxes, 2))

def distance(a, b):
    x1, y1, z1 = a
    x2, y2, z2 = b
    return math.sqrt((x1-x2)**2 + (y1-y2)**2 + (z1-z2)**2)

distances = {}
for pair in pairs:
    d = distance(pair[0], pair[1])
    distances[d] = pair

# each jbox starts in its own circuit
circuits = [{jbox} for jbox in jboxes]

for d in sorted(distances.keys())[:1000]:
    a, b = distances[d]

    # find circuit containing jbox a
    for i, circuit in enumerate(circuits):
        if a in circuit:
            break

    # check if jbox b is already part of same circuit
    if b in circuit:
        continue

    ca = circuits.pop(i)

    # find circuit containing jbox b
    for i, circuit in enumerate(circuits):
        if b in circuit:
            break
    # merge circuits for a and b
    circuit |= ca

# compute size of all circuits
sizes = list(map(len, circuits))
# multiple largest 3 circuit sizes
product = 1
for size in sorted(sizes, reverse=True)[:3]:
    product *= size
print(product)