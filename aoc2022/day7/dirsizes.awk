#!/usr/bin/awk -f

# Advent of Code 2022 Day 7

# Takes as input commands and outputs (one command or one output per line) representing filesystem
# navigation and examination.  Outputs the size of all directories.  The size of a directory is
# computed as the sum of all file sizes within that directory and subdirectories.

# e.g., compute sum of all directories whose size is at least 100,000 (uses day1 solution to compute sum)
# src/main/awk/2022/day7/dirsizes.awk input.txt | awk '$2 <= 100000 {print $2}' | src/main/awk/2022/day1/sumgroups.awk
#
# e.g., find the size of the smallest directory that would free up at least 30,000,000 units of space given total a
# total capacity of 70,000,000 units
# src/main/awk/2022/day7/dirsizes.awk input.txt | \
#     awk '/^Total used space/ {needed = 30000000 - (70000000 - $4); next} $2 >= needed {print $2}' | sort -n | head -1

# use `cd` commands to track the current working directory
# command `cd /`
/^\$ cd \// {
    cwd = "/"
    next
}
# command `cd ..`
/^\$ cd \.\./ {
    # chop off the last `/...` from the current directory
    sub(/\/[^\/]*$/, "", cwd)
    next
}
# command `cd <dir>`
/^\$ cd/ {
    cwd = cwd"/"$3
    next
}

# match a file with format:  `<size> <filename>`
/^[0-9]+/ {
    # for each file found, add its size to every directory in its path
    # keep track of directory sizes using a map keyed by the directory path
    dir = cwd
    do {
        sizes[dir] += $1
        # move up one directory (chop off the last `/...` from the directory)
        sub(/\/[^\/]*$/, "", dir)
    }
    while (length(dir) > 0)
    next
}

END {
    printf("Total used space: %d\n", sizes["/"])
    for (dir in sizes) {
        printf("%s: %d\n", dir, sizes[dir])
    }
}