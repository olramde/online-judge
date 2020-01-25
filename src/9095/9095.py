num_case = int(input())
test_case = []
for _ in range(num_case):
    test_case.append(int(input()))
max_test_case = max(test_case)

dp = []
dp.append(0) # 0
dp.append(1) # 1
dp.append(2) # 2
dp.append(4)


for i in range(4, max_test_case+1):
    dp.append(dp[i-1] + dp[i-2] + dp[i-3])
    
for case in test_case:
    print(dp[case])