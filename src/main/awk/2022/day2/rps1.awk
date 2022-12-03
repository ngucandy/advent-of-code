#!/usr/bin/awk -f

# Advent of Code 2022 Day 2 part 1

# $ awk/2022/day2/rps1.awk resources/2022/day2.txt | awk/2022/day1/sumgroups.awk
# 17189

# A = rock, B = paper, C = scissors
# X = rock, Y = paper, Z = scissors
# lose = 0, draw = 3, win = 6
# rock = 1, paper = 2, scissors = 3

# rock vs. rock (draw)
/A X/ {
    print(3 + 1)
}

# rock vs. paper (win)
/A Y/ {
    print(6 + 2)
}

# rock vs. scissors (lose)
/A Z/ {
    print(0 + 3)
}

# paper vs. rock (lose)
/B X/ {
    print(0 + 1)
}

# paper vs. paper (draw)
/B Y/ {
    print(3 + 2)
}

# paper vs. scissors (win)
/B Z/ {
    print(6 + 3)
}

# scissors vs. rock (win)
/C X/ {
    print(6 + 1)
}

# scissors vs. paper (lose)
/C Y/ {
    print(0 + 2)
}

# scissors vs. scissors (draw)
/C Z/ {
    print(3 + 3)
}
