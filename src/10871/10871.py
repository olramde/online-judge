maxNum = int(input().split()[1])

arr = map(int, input().split())

for i in arr:
    if i < maxNum:
        print(i, end=" ")

