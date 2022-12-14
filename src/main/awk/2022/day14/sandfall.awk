#!/usr/bin/env awk -f

# Advent of Code 2022 Day 14 part 1
BEGIN {
    FS = " -> "
    MINX = 500
    MAXX = 500
    MINY = 0
    MAXY = 0
    SOURCEX = 500
    SOURCEY = 0
    CAVE[SOURCEX, SOURCEY] = "+"
}

{
    for (i = 2; i <= NF; ++i) {
        split($(i - 1), xy1, ",")
        split($i, xy2, ",")
        for (x = min(xy1[1], xy2[1]); x <= max(xy1[1], xy2[1]); ++x) {
            for (y = min(xy1[2], xy2[2]); y <= max(xy1[2], xy2[2]); ++y) {
                CAVE[x, y] = "#"
            }
        }

        MINX = min(xy2[1], min(xy1[1], MINX))
        MAXX = max(xy2[1], max(xy1[1], MAXX))
        MINY = min(xy2[2], min(xy1[2], MINY))
        MAXY = max(xy2[2], max(xy1[2], MAXY))
    }
}

END {
    for (i = 1; MAXY > dropSand(); ++i) {
    }
    print(i - 1)
    printCave()
}

function min(x, y) {
    return x < y ? x : y
}

function max(x, y) {
    return x > y ? x : y
}

function dropSand() {
    x = SOURCEX
    for (y = SOURCEY; y < MAXY; ++y) {
        if (CAVE[x, y + 1] == "") {
            continue
        }
        if (CAVE[x - 1, y + 1] == "") {
            x -= 1
            continue
        }
        if (CAVE[x + 1, y + 1] == "") {
            x += 1
            continue
        }
        CAVE[x, y] = "o"
        return y
    }
    return MAXY
}

function printCave() {
    for (y = MINY; y <= MAXY; ++y) {
        for (x = MINX; x <= MAXX; ++x) {
            printf("%s", CAVE[x, y] == "" ? "." : CAVE[x, y])
        }
        printf("\n")
    }
}