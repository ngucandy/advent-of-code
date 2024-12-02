#!/usr/bin/env awk -f

# Advent of Code 2022 Day 15 part 1

BEGIN {
    MINX = 2^32
    MAXX = -MINX
}

{
    split($3, a, "[=,]")
    sensorX = a[2]
    split($4, a, "[=:]")
    sensorY = a[2]
    split($9, a, "[=,]")
    beaconX = a[2]
    split($10, a, "=")
    beaconY = a[2]
    DISTANCES[sensorX, sensorY] = distance(sensorX, sensorY, beaconX, beaconY)
    BEACONS[beaconX, beaconY] = 1
    printf("sensor[%d, %d], beacon[%d, %d], distance: %d\n", sensorX, sensorY, beaconX, beaconY, DISTANCES[sensorX, sensorY])
}

END {
#    coverage(9)
#    coverage(10)
#    coverage(11)
    coverage(2000000)
}

function abs(n) {
    return n < 0 ? -n : n
}

function min(x, y) {
    return x < y ? x : y
}

function max(x, y) {
    return x > y ? x : y
}

function distance(x1, y1, x2, y2) {
    return abs(x1 - x2) + abs(y1 - y2)
}
function printRow(first, last, a) {
    printf("%d ", first)
    for (x = first; x <= last; ++x) {
        printf("%s", x in a ? a[x] : ".")
    }
    printf(" %d\n", last)
}

function coverage(y) {
    for (sensor in DISTANCES) {
        split(sensor, s, SUBSEP)
        sX = s[1]
        sY = s[2]
        printf("sensor[%d, %d]\n", sX, sY)
        distanceY = abs(sY - y)
        if (distanceY > DISTANCES[sX, sY]) {
            printf("row %d is out of range: %d > %d\n", y, distanceY, DISTANCES[sX, sY])
            continue
        }
        dExtra = abs(DISTANCES[sX, sY] - distanceY)
        for (x = sX - dExtra; x <= sX + dExtra; ++x) {
            if (BEACONS[x, y] == 1) {
                col[x] = "B"
            }
            else if ((x, y) in DISTANCES) {
                col[x] = "S"
            }
            else {
                col[x] = "#"
            }
        }
        MINX = min(MINX, sX - dExtra)
        MAXX = max(MAXX, sX + dExtra)
    }
    count = 0
    for (x = MINX; x <= MAXX; ++x) {
        count += col[x] == "#" ? 1 : 0
    }
    print(count)
}