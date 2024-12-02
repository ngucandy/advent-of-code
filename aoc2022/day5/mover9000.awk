#!/usr/bin/awk -f

# Advent of Code 2022 Day 5 Part 1

# use lines that don't start with "move" to initialize the stacks
!/^move/ {
    gsub(/.{3} /, "&-", $0)
    n = split($0, a, /-/)
    for (i = 1; i <= n; i++) {
        crate = substr(a[i], 2, 1)
        gsub(/ /, "", crate)
        stacks[i] = sprintf("%s%s", crate, stacks[i])
    }
}

# empty line separates stack state and moves section
length($0) == 0 {
    # remove the stack number from the stack
    print("Initial stacks")
    for (i = 1; i <= length(stacks); i++) {
        stacks[i] = substr(stacks[i], 2, length(stacks[i]) - 1)
        printf("[%s]\n", stacks[i])
    }
    print("Moving")
    next
}

# lines that start with "move"
/^move/ {
    qty = $2
    src = $4
    dest = $6
    printf("%d: [%s] -> [%s]", qty, stacks[src], stacks[dest])
    for (i = 1; i <= qty; i++) {
        crate = substr(stacks[src], length(stacks[src]), 1)
        stacks[src] = substr(stacks[src], 1, length(stacks[src]) - 1)
        stacks[dest] = sprintf("%s%s", stacks[dest], crate)
    }
    printf(" ==> [%s] -> [%s]\n", stacks[src], stacks[dest])
}

END {
    print("Final stacks")
    for (i = 1; i <= length(stacks); i++) {
        printf("[%s]\n", stacks[i])
        message = sprintf("%s%s", message, substr(stacks[i], length(stacks[i]), 1))
    }
    print(message)
}
