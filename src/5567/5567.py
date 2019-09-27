numOfCowoker = int(input())
listLength = int(input())
friendRelationList = []
friendOfFriendRelationList = []
friendSet = set()
friendOfFriendSet = set()

for i in range(listLength):    
    friendRelationList.append(list(map(int, input().split())))

for f1, f2 in friendRelationList:
    if f1 == 1:
        friendSet.add(f2)
    elif f2 == 1:
        friendSet.add(f1)
    else:
        friendOfFriendRelationList.append([f1, f2])

for f1, f2 in friendOfFriendRelationList:
    if f1 in friendSet:
        friendOfFriendSet.add(f2)
    elif f2 in friendSet:
        friendOfFriendSet.add(f1)

print(len(friendSet.union(friendOfFriendSet)))