#!/usr/bin/awk -f

# Advent of Code 2022 Day 3 Part 2

# $ src/main/awk/2022/day3/badge.awk src/main/resources/2022/day3.txt | src/main/awk/2022/day1/sumgroups.awk                                                                                                   (main✱)
# 2639
BEGIN {
    ITEMS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    FS = ""
}
{
    sacknum = NR % 3
    for (i = 1; i <= NF; i++) {
        priority = index(ITEMS, $i)
        sack[sacknum"-"priority] = 1
    }
}
NR % 3 == 0 {
    for (i = 1; i <= NF; i++) {
        priority = index(ITEMS, $i)
        sack["3-"priority] = 1
    }
    for (i = 1; i <= length(ITEMS); i++) {
        if (sack["1-"i] + sack["2-"i] + sack["3-"i] == 3) {
            print(i)
        }
    }
    delete sack
}