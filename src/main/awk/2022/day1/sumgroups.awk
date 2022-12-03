#!/usr/bin/awk -f

# Advent of Code 2022 Day 1

# Takes as input a list of numbers (one number per line) and sums the values
# using empty lines as a group delimiter.
#
# e.g., sum the top 3 groups
# ./sumgroups.awk input.txt | sort -n | tail -3 | ./sumgroups.awk
length($0) == 0 {
    print(sum)
    sum = 0
}
length($0) > 0 {
    sum += $0
}
END {
    print(sum)
}