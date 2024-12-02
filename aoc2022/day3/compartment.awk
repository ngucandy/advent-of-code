#!/usr/bin/awk -f

# Advent of Code 2022 Day 3 Part 1

# $ src/main/awk/2022/day3/rucksack.awk src/main/resources/2022/day3.txt | src/main/awk/2022/day1/sumgroups.awk                                                                                                (main✱)
# 7811
BEGIN {
    ITEMS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    FS = ""
}
{
    mid = NF / 2
    for (i = 1; i <= mid; i++) {
        priority = index(ITEMS, $i)
        comp1[priority] = 1
    }
    for (i = mid + 1; i <= NF; i++) {
        priority = index(ITEMS, $i)
        comp2[priority] = 1
    }
    for (i = 1; i <= length(ITEMS); i++) {
        if (comp1[i] + comp2[i] == 2) {
            print(i)
        }
    }
    delete comp1
    delete comp2
}