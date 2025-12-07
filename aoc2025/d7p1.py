grid = [list(line.strip()) for line in open(0)]

rows = len(grid)
cols = len(grid[0])
for r in range(rows):
    for c in range(cols):
        if grid[r][c] == 'S':
            sr = r
            sc = c
            break
        else:
            continue
    break

q = [(sr, sc)]
split = set()
seen = set()
while q:
    r, c = q.pop(0)
    if (r, c) in seen:
        continue
    seen.add((r, c))
    if c < 0 or c > cols:
        continue
    for nr in range(r + 1, rows):
        if grid[nr][c] == '^':
            split.add((nr, c))
            q.append((nr, c-1))
            q.append((nr, c+1))
            break

print(len(split))