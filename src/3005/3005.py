row, col = map(int, input().split())

EMPTY_STR = '#'

puzzle = []
for i in range(row):
        puzzle.append(list(input()))
        
word_list = []
for r_idx in range(row):
    word = ''
    for c_idx in range(col):
        character = puzzle[r_idx][c_idx]
        if character != EMPTY_STR:
            word += character
        else:
            if len(word) >= 2:
                word_list.append(word)
            word = ''
    if len(word) >= 2:
        word_list.append(word)
    word = ''

for c_idx in range(col):
    word = ''
    for r_idx in range(row):
        character = puzzle[r_idx][c_idx]
        if character != EMPTY_STR:
            word += character
        else:
            if len(word) >= 2:
                word_list.append(word)
            word = ''
    if len(word) >= 2:
        word_list.append(word)
    word = ''
    
print(sorted(word_list)[0])
        