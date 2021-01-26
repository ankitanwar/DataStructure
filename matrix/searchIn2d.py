from typing import List
def searchMatrix(matrix: List[List[int]], target: int):
    row=-1
    for i in range(len(matrix)):
        if matrix[i][0]>target:
            row=i-1
            break
    if row==-1:
        if matrix[len(matrix)-1][len(matrix[0])-1]<target:
            return False
        else:
            for j in range(len(matrix[0])):
                if matrix[len(matrix)-1][j]==target:
                    return True
            return False

    for j in range(len(matrix[0])):
        if matrix[row][i]==target:
            return True
        
    return False

