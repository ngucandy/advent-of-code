#!/usr/bin/env awk -f

# Advent of Code 2022 Day 8 part 2

function score(x, y) {
    return viewUp(x, y) * viewDown(x, y) * viewLeft(x, y) * viewRight(x, y)
}

function viewUp(x, y) {
    if (x == 1) {
        return 0
    }
    for (i = x-1; i >= 1; i--) {
        if (matrix[x, y] <= matrix[i, y]) {
            return x-i
        }
    }
    return x-1
}

function viewDown(x, y) {
    if (x == height) {
        return 0
    }
    for (i = x+1; i <= height; i++) {
        if (matrix[x, y] <= matrix[i, y]) {
            return i-x
        }
    }
    return height-x
}

function viewLeft(x, y) {
    if (y == 1) {
        return 0
    }
    for (i = y-1; i >= 1; i--) {
        if (matrix[x, y] <= matrix[x, i]) {
            return y-i
        }
    }
    return y-1
}

function viewRight(x, y) {
    if (y == width) {
        return 0
    }
    for (i = y+1; i <= width; i++) {
        if (matrix[x, y] <= matrix[x, i]) {
            return i-y
        }
    }
    return width-y
}

{
    split($0, heights, "")
    height = NR
    width = length(heights)
    for (y in heights) {
        matrix[NR, y] = heights[y]
    }
}

END {
    printf("Matrix size: %dx%d\n", height, width)
    for (x = 1; x <= height; x++) {
        for (y = 1; y <= width; y++) {
            sscore = score(x, y)
            printf("%d", sscore)
            max = sscore > max ? sscore : max
        }
        printf("\n")
    }
    printf("Max scenic score: %d\n", max)
}