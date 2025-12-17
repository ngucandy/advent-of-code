blocks = open(0).read().split('\n\n')
total = 0
for line in blocks[-1].splitlines():
    region, counts = line.split(': ')
    width, length = map(int, region.split('x'))
    presents = sum(list(map(int, counts.split())))

    # try to fit presents in region with no overlap
    fitwidth = width // 3
    fitlength = length // 3
    fitmax = fitwidth * fitlength
    if presents <= fitmax:
        total += 1

print(total)