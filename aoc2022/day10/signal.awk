#!/usr/bin/env awk -f

# Advent of Code 2022 Day 10 part 1

# Takes the output from register.awk and computes the sum of signal strenghts at
# cycles 20, 60, 100, 140, 180 and 220.  The signal strength at a given cycle is
# calculated as the cycle number * the X register value.
#
# e.g.,
# $ src/main/awk/2022/day10/register.awk input.txt | src/main/awk/2022/day10/signal.awk | src/main/awk/2022/day1/sumgroups.awk
NR == 20 || (NR - 20) % 40 == 0 {
    print(NR * $0)
}