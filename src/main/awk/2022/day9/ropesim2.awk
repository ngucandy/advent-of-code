#!/usr/bin/env awk -f

# Advent of Code 2022 Day 9 part 2

function abs(n) {
    return n < 0 ? -n : n
}

# moves the tail closer to x,y
function moveTail(x, y, tailnum) {
    if (TX[tailnum] != x) {
        TX[tailnum] += x > TX[tailnum] ? 1 : -1
    }

    if (TY[tailnum] != y) {
        TY[tailnum] += y > TY[tailnum] ? 1 : -1
    }

    # track the visisted positions of the last tail
    if (tailnum == TAILS) {
        TRACK[TX[tailnum], TY[tailnum]] = 1
    }
    else {
        # move the next tail if the current tail is 2 positions away in either direction
        if (abs(TX[tailnum] - TX[tailnum+1]) > 1 || abs(TY[tailnum] - TY[tailnum+1]) > 1) {
            # recursion!!
            moveTail(TX[tailnum], TY[tailnum], tailnum+1)
        }
    }
}

# move the head left or right
function moveHeadHoriz(direction, times) {
    for (i = 0; i < times; i++) {
        HX += direction
        if (HX < 0) {
            print("Starting position too low: " START)
            exit 1
        }
        # move the first tail if the head is now 2 positions away
        if (abs(HX - TX[1]) > 1) {
            moveTail(HX, HY, 1)
        }
    }
}

# move the head up or down
function moveHeadVert(direction, times) {
    for (i = 0; i < times; i++) {
        HY += direction
        if (HY < 0) {
            print("Starting position too low: " START)
            exit 1
        }
        # move the first tail if the head is now 2 positions away
        if (abs(HY - TY[1]) > 1) {
            moveTail(HX, HY, 1)
        }
    }
}

BEGIN {
    # starting positions
    START = 300
    TAILS = 9
    HX = START
    HY = HX
    for (i = 1; i <= TAILS; i++ ) {
        TX[i] = HX
        TY[i] = HY
    }
    # only track the last tail
    TRACK[TX[TAILS], TY[TAILS]] = 1
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
    print("Positions visited by last tail: " visited)
}