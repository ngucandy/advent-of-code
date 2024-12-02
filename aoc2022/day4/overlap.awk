#!/usr/bin/awk -f

# Advent of Code 2022 Day 4 Part 2

# $ src/main/awk/2022/day4/overlap.awk src/main/resources/2022/day4.txt | src/main/awk/2022/day1/sumgroups.awk                                                                                                 (main✱)
# 914
BEGIN {
    FS = "[,-]"
}
{
    if (($1 >= $3 && $1 <= $4) ||
        ($2 >= $3 && $2 <= $4) ||
        ($3 >= $1 && $3 <= $2) ||
        ($4 >= $1 && $4 <= $2)) {
        print(1)
    }
}