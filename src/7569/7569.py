import copy

m, n, h = list(map(int, input().split())) # ROW COL HEI

box = []
for i in range(h):
    height = []
    for j in range(n):
        row = list(map(int, input().split()))
        height.append(row)
    box.append(height)

# [H, N, M]

def tommorow_box(box):
    next_box = copy.deepcopy(box)
    
def get_adjacent_idx_list(shape, idx):
    adjacent_idx_list = [
        [idx[0]+1, idx[1], idx[2]], [idx[0]-1, idx[1], idx[2]],
        [idx[0], idx[1]+1, idx[2]], [idx[0], idx[1]-1, idx[2]],
        [idx[0], idx[1], idx[2]+1], [idx[0], idx[1], idx[2]-1]
    ]
    return list(filter(is_valid_idx(shape), adjacent_idx_list))
    
def is_valid_idx(shape):
    def is_valid_idx(idx):
        for i in range(3):
            if not (idx[i] >= 0 and idx[i] < shape[i]):
                return False
        return True
    return is_valid_idx