from collections import deque

devices = {k: v.strip().split() for k, v in (line.strip().split(':') for line in open(0))}

count = 0
q = deque([["you"]])
while q:
    path = q.popleft()
    device = path[-1]
    if device == "out":
        count += 1
        continue
    for output in devices[device]:
        # skip if output has already been seen on this path to avoid loops
        if output in path:
            continue
        q.append(path + [output])

print(count)
