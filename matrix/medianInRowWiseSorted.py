import math
def median(matrix,r,c):
    ans=search(matrix,10)
    print(ans)

def search(matrix,target):
    #given a matrix which is sorted both row wise and column wise and we need to find the given key
    i=0
    j=len(matrix[0])-1

    while i>=0 and i<len(matrix) and j>=0 and j<len(matrix[0]):
        if matrix[i][j]==target:
            return True
        if matrix[i][j]>target:
            j-=1
        else:
            i+=1
    return False


mat=[
    [1,2,3,4],
    [5,6,7,8],
    [9,10,11,12],
    [13,14,15,16]
]

median(mat,1,1)


