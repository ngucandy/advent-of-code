from functools import cache

devices = {k: v.strip().split() for k, v in (line.strip().split(':') for line in open(0))}

@cache
def count_paths(src, dst):
    if src == dst:
        return 1

    if src not in devices:
        return 0

    count = 0
    for output in devices[src]:
        count += count_paths(output, dst)
    return count

paths_dac_fft = count_paths("svr", "dac") * count_paths("dac", "fft") * count_paths("fft", "out")
paths_fft_dac = count_paths("svr", "fft") * count_paths("fft", "dac") * count_paths("dac", "out")
print(paths_dac_fft + paths_fft_dac)