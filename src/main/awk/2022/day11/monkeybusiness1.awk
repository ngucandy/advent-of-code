#!/usr/bin/env awk -f

# Advent of Code 2022 Day 11 part 1

# start of a new monkey
/^Monkey/ {
    sub(/:/, "", $2)
    monkeyNum = $2
}

# starting item values for a monkey
/^  Starting/ {
    for (i = 3; i <= NF; i++) {
        sub(/,/, "", $i)
        # item values keyed by monkey number and item number
        items[monkeyNum, i-2] = $i
    }
    # track the total number of items per monkey
    numItems[monkeyNum] = NF - 3 + 1
}

/^  Operation/ {
    # operation (addition or multiplication) per monkey
    operator[monkeyNum] = $5
    # 2nd operand value for the operation (per monkey)
    operand[monkeyNum] = $6
}

/^  Test/ {
    # divisibility test value (per monkey)
    test[monkeyNum] = $4
}

/^    If true/ {
    # monkey number to receive item if test is true (per monkey)
    targetTrue[monkeyNum] = $6
}

/^    If false/ {
    # monkey number to receive item if test is false (per monkey)
    targetFalse[monkeyNum] = $6
}

END {
    for (round = 1; round <= 20; round++) {
        print("Round " round)
        for (monkey = 0; monkey <= monkeyNum; monkey++) {
            for (item = 1; item <= numItems[monkey]; item++) {
                currentValue = items[monkey, item]
                operand2 = operand[monkey] ~ /old/ ? currentValue : operand[monkey]
                newValue = operator[monkey] ~ /\+/ ? currentValue + operand2 : currentValue * operand2
                newValue = int(newValue / 3)
                targetMonkey = newValue % test[monkey] == 0 ? targetTrue[monkey] : targetFalse[monkey]
                # add new item value to target monkey
                numItems[targetMonkey]++
                items[targetMonkey, numItems[targetMonkey]] = newValue
                inspected[monkey]++
            }
            numItems[monkey] = 0
            printf("Monkey %d inspected items %d times\n", monkey, inspected[monkey])
        }
    }
}