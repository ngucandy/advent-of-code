top, bottom = open(0).read().split("\n\n")
ranges = [tuple(map(int, range.split("-"))) for range in top.splitlines()]
nums = list(map(int, bottom.splitlines()))

count = 0
for num in nums:
    for s, e in ranges:
        if s <= num <= e:
            count += 1
            break

print(count)