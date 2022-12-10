#!/usr/bin/env awk -f

# Advent of Code 2022 Day 10 part 2

# Redners the output register.awk.
#
# e.g.,
# $ src/main/awk/2022/day10/register.awk input.txt | src/main/awk/2022/day10/render.awk

# start new row every 40 cycles
(NR - 1) % 40 == 0 {
    printf("\n")
}

# for every cycle
{
    # the position in the row the crt is currently rendering
    crtPosition = (NR - 1) % 40
    if (crtPosition >= ($0 - 1) && crtPosition <= ($0 + 1)) {
        printf("#")
    }
    else {
        printf(".")
    }
}

END {
    printf("\n")
}
