from functools import cache

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

# recurse with memoization
@cache
def paths(r, c):
    for nr in range(r+1, len(grid)):
        if nr == rows-1:
            return 1

        if grid[nr][c] == '^':
            left, right = 0, 0
            if c > 0:
                left = paths(nr, c-1)
            if c < cols-1:
                right = paths(nr, c+1)
            return left + right

    return 0

print(paths(sr, sc))
