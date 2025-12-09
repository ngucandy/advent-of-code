# only strip newlines; other whitespace is important
grid = [line.rstrip('\n') for line in open(0)]
nums = []
total = 0
for c in reversed(range(len(grid[0]))):
    num = ''
    op = ''
    for r in range(0, len(grid)):
        if grid[r][c] == '+' or grid[r][c] == '*':
            op = grid[r][c]
            break
        num += grid[r][c]
    if len(num.strip()) == 0:
        continue
    nums.append(int(num))
    if len(op) > 0:
        if op == '+':
            total += sum(nums)
        else:
            product = 1
            for num in nums:
                product *= num
            total += product
        nums = []
        op = ''

print(total)