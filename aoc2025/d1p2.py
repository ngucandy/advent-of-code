rotations = [line.strip() for line in open(0)]
dial = 50
count = 0
for rotation in rotations:
    direction = rotation[:1]
    distance = int(rotation[1:])
    if distance > 100:
        count += distance // 100
        distance %= 100
    if direction == 'L':
        if 0 < dial <= distance: count += 1
        distance = -distance
    else:
        if distance >= 100-dial: count += 1
    dial = (dial + distance) % 100
print(count)