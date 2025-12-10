from collections import deque

lines = [list(line.strip().split()) for line in open(0)]

indicators = []
schematics = []

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
    q = deque([(list('.' * len(want)), [])])
    while q:
        lights, presses = q.popleft()
        if lights == want:
            total += len(presses)
            break
        for i, button in enumerate(buttons):
            if len(presses) > 0 and i == presses[-1]:
                continue
            newlights = press(lights, button)
            q.append((list(newlights), presses+[i]))

print(total)
