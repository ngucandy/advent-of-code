#!/usr/bin/env awk -f

# Advent of Code 2022 Day 10 part 1

# Takes as input a series of instructions noop and addx instructions (one per line)
# and outputs the X register value during the execution cycles of the instruction.
# A noop instruction takes one cycle to execute and addx takes two.
BEGIN {
    X = 1
}

/^noop/ {
    print(X)
}

/^addx/ {
    print(X)
    print(X)
    X += $2
}
