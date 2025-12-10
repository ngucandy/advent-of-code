top = open(0).read().split("\n\n")[0]
inputs = [tuple(map(int, range.split("-"))) for range in top.splitlines()]
ranges = []
# for each input range, check for overlap with existing ranges
while inputs:
    s, e = inputs.pop(0)
    for i in range(0, len(ranges)):
        rs, re = ranges[i]
        if s > re or e < rs:
            # input range is completely outside existing range
            continue
        if rs <= s and re >= e:
            # input range is fully contained in existing range
            break
        if s <= rs and e >= re:
            # input range fully encloses existing range
            ranges.pop(i)
            inputs.append((s, e))
            break
        if e > re:
            # input range overlaps past the end of the existing range
            ranges.pop(i)
            inputs.append((rs, e))
            break
        if s < rs:
            # input range overlaps before the existing range
            ranges.pop(i)
            inputs.append((s, re))
            break
    else:
        # input range does not overlap with any existing range
        ranges.append((s, e))

print(sum([e-s+1 for s, e in ranges]))
