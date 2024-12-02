#!/usr/bin/env awk -f

# Advent of Code 2022 Day 8 part 1

function isVisible(x, y) {
    return isVisibleUp(x, y) || isVisibleDown(x, y) || isVisibleLeft(x, y) || isVisibleRight(x, y)
}

function isVisibleUp(x, y) {
    if (x == 1) {
        return 1
    }
    for (i = x-1; i >= 1; i--) {
        if (matrix[x, y] <= matrix[i, y]) {
            return 0
        }
    }
    return 1
}

function isVisibleDown(x, y) {
    if (x == height) {
        return 1
    }
    for (i = x+1; i <= height; i++) {
        if (matrix[x, y] <= matrix[i, y]) {
            return 0
        }
    }
    return 1
}

function isVisibleLeft(x, y) {
    if (y == 1) {
        return 1
    }
    for (i = y-1; i >= 1; i--) {
        if (matrix[x, y] <= matrix[x, i]) {
            return 0
        }
    }
    return 1
}

function isVisibleRight(x, y) {
    if (y == width) {
        return 1
    }
    for (i = y+1; i <= width; i++) {
        if (matrix[x, y] <= matrix[x, i]) {
            return 0
        }
    }
    return 1
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
            visible = isVisible(x, y)
            printf("%d", visible)
            totalVisible += visible
        }
        printf("\n")
    }
    printf("Total Visible: %d\n", totalVisible)
}