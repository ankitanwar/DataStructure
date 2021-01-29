def commonElement(matrix):
    values={}
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if matrix[i][j] not in values:
                t=[i,1]
                values[matrix[i][j]]=t
            else:
                t=values[matrix[i][j]]
                if t[0]!=i:
                    t[0]=i
                    t[1]+=1
                    values[matrix[i][j]]=t
    for key in values:
        t=values[key]
        if t[1]==len(matrix):
            print(key)


mat = [[1, 2, 1, 4, 8], 
       [3, 7, 8, 5, 1], 
       [8, 7, 7, 3, 1], 
       [8, 1, 2, 7, 9]] 
commonElement(mat)

