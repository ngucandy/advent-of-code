#!/usr/bin/awk -f

# Advent of Code 2022 Day 6

# Takes as input a single string on one line and finds the position in the
# string where the previous n characters contain no duplicates.  n can be
# specified using `-v wsize=n` and defaults to 4.  Outputs the characters
# in the current window for each position.
#
# e.g., windows size 14
# ./marker.awk -v wsize=14 input.txt
BEGIN {
    # separate each character into its own field
    FS = ""
    # window size
    wsize = wsize > 0 ? wsize : 4
}

{
    # packet: array of length wsize containing the characters in the current packet
    # counts: map containing the count of each character in the current packet

    # iteratre over fields/characters
    for (i = 1; i <= NF; i++) {
        # index into the packet array
        j = (i - 1) % wsize

        # replace the character at packet[j]
        old = packet[j]
        counts[old] = counts[old] > 0 ? counts[old] - 1 : 0  # ensure count doesn't go below 0
        counts[$i]++
        packet[j] = $i

        # build the current marker and sum the character counts
        for (x = 0; x < wsize; x++) {
            marker = marker packet[x]
            sum += counts[packet[x]]
        }
        printf("%d: %s\n", i, marker)

        # sum of character counts should equal the window size if there are no duplicates,
        # but only if we've read at least a full window size amount of characters
        if (i >= wsize && sum == wsize) {
            break
        }
        marker = ""
        sum = 0
    }
}
