input = open("input.txt")
elves = []
elf = []
for line in input:
    if line == "\n":
        elves.append(elf)
        elf = []
    elf.append(line.strip())
elves.append(elf)

calories = []
for elf in elves:
    sum = 0
    for snack in elf:
        if snack != '':
            sum += int(snack)
    calories.append(sum)
calories.sort()
print(calories[-1])
print(calories[-1]+calories[-2]+calories[-3])
