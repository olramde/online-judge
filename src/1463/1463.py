from math import inf

def isPossibleDivThree(x):
    if x % 3 == 0:
        return True
    return False

def isPossibleDivTwo(x):
    if x % 2 == 0:
        return True
    return False

num = int(input())

dp = [None]*(num+1)

if num == 0 or num == 1:
    print(0)
else:
    dp[0] = 0
    dp[1] = 0

    for i in range(2, num+1):
        min = inf
        if isPossibleDivThree(i):
            if min > dp[i//3]:
                min = dp[i//3]
        if isPossibleDivTwo(i):
            if min > dp[i//2]:
                min = dp[i//2]
        if min > dp[i-1]:
            min = dp[i-1]
        dp[i] = min+1

    print(dp[num])