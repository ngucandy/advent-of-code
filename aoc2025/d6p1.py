lines = [line.strip().split() for line in open(0)]
nums = [list(map(int, elements)) for elements in lines[:-1]]
ops = lines[-1]
grand_total = 0
for i, op in enumerate(ops):
    if op == '+':
        total = 0
        for num in nums:
            total += num[i]
    else:
        total = 1
        for num in nums:
            total *= num[i]
    grand_total += total
print(grand_total)