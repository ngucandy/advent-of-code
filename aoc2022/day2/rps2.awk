#!/usr/bin/awk -f

# Advent of Code 2022 Day 2 part 2

# $ awk/2022/day2/rps2.awk resources/2022/day2.txt | awk/2022/day1/sumgroups.awk
# 13490

# A = rock, B = paper, C = scissors
# X = lose, Y = draw, Z = win
# lose = 0, draw = 3, win = 6
# rock = 1, paper = 2, scissors = 3

# rock vs. scissors (lose)
/A X/ {
    print(0 + 3)
}

# rock vs. rock (draw)
/A Y/ {
    print(3 + 1)
}

# rock vs. paper (win)
/A Z/ {
    print(6 + 2)
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

# scissors vs. paper (lose)
/C X/ {
    print(0 + 2)
}

# scissors vs. scissors (draw)
/C Y/ {
    print(3 + 3)
}

# scissors vs. rock (win)
/C Z/ {
    print(6 + 1)
}
