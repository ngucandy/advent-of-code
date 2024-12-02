#!/usr/bin/env awk -f

# Advent of Code 2022 Day 9 part 1

function abs(n) {
    return n < 0 ? -n : n
}

# moves the tail to position x,y
function moveTail(x, y) {
    TX = x
    TY = y
    TRACK[x, y] = 1
}

function moveHeadHoriz(direction, times) {
    for (i = 0; i < times; i++) {
        prevHX = HX
        HX += direction
        if (HX < 0) {
            print("Starting position too low: " START)
            exit 1
        }
        if (abs(HX - TX) > 1) {
            moveTail(prevHX, HY)
        }
    }
}

function moveHeadVert(direction, times) {
    for (i = 0; i < times; i++) {
        prevHY = HY
        HY += direction
        if (HY < 0) {
            print("Starting position too low: " START)
            exit 1
        }
        if (abs(HY - TY) > 1) {
            moveTail(HX, prevHY)
        }
    }
}

BEGIN {
    # starting positions
    START = 300
    HX = START
    HY = HX
    TX = HX
    TY = HY
    TRACK[TX, TY] = 1
}

# move right
/^R/ {
    moveHeadHoriz(1, $2)
}

# move left
/^L/ {
    moveHeadHoriz(-1, $2)
}

# move up
/^U/ {
    moveHeadVert(1, $2)
}

# move down
/^D/ {
    moveHeadVert(-1, $2)
}

END {
    for (pos in TRACK) {
        visited += TRACK[pos]
    }
    print("Positions visited by tail: " visited)
}