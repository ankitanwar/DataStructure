def RotateMatrix(matrix):
    low=0
    high=len(matrix)-1
    while high>low:
        matrix[low],matrix[high]=matrix[high],matrix[low]
        low+=1
        high-=1
    i=0
    while i+1<len(matrix):
        start=i+1
        right=i+1

        while start<len(matrix) and right<len(matrix[0]):
            matrix[start][i],matrix[i][right]=matrix[i][right],matrix[start][i]
            start+=1
            right+=1
        i+=1
    print(matrix)



matrix = [[1,2,3],[4,5,6],[7,8,9]]
RotateMatrix(matrix)