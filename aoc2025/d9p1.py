from itertools import combinations

tiles = [tuple(map(int, line.strip().split(','))) for line in open(0)]
pairs = list(combinations(tiles, 2))
areas = [(abs(p1[0]-p2[0])+1) * (abs(p1[1]-p2[1])+1) for p1, p2 in pairs]

print(max(areas))