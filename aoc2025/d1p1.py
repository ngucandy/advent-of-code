rotations = [line.strip() for line in open(0)]
dial = 50
count = 0
for rotation in rotations:
    direction = rotation[:1]
    distance = int(rotation[1:])
    if direction == 'L': distance = -distance
    dial = (dial + distance) % 100
    if dial == 0: count += 1
print(count)