from collections import deque

lines = [list(line.strip().split()) for line in open(0)]

indicators = []
schematics = []
# jolts = []

for line in lines:
    indicator = list(line[0][1:-1])
    buttons = [list(map(int, button[1:-1].split(','))) for button in line[1:-1]]
    indicators.append(indicator)
    schematics.append(buttons)

def press(lights, button):
        newlights = lights[:]
        for i in button:
            if newlights[i] == '.':
                newlights[i] = '#'
            else:
                newlights[i] = '.'
        return newlights

total = 0
for m in range(len(indicators)):
    want = indicators[m]
    buttons = schematics[m]
    q = deque()
    q.append((list('.' * len(want)), 0))
    while q:
        lights, presses = q.popleft()
        if lights == want:
            # print(m, presses)
            total += presses
            break
        for button in buttons:
            newlights = press(lights, button)
            q.append((list(newlights), presses+1))

print(total)
