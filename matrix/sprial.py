def sprialMatrix(matrix):
    LtoR=0
    topToBottom=len(matrix[0])
    bottomToLeft=len(matrix)
    bottomTOUp=len(matrix)
    ans=[]
    while LtoR<bottomTOUp:
        for i in range(topToBottom):
            ans.append(matrix[LtoR][i])
        LtoR+=1
        for i in range(LtoR,bottomToLeft):
            ans.append(matrix[i][bottomToLeft])
        topToBottom-=1
        for i in range(bottomToLeft-1,bottomTOUp,-1):
            ans.append(ans[bottomToLeft][i])
        for i in range(bottomToLeft,LtoR,-1):
            print("The value is ",bottomTOUp,i)
            matrix[i][bottomTOUp]
        bottomToLeft-=1
        bottomTOUp+=1
    
    print("The final ans is",ans)



mat=[
    [1,2,3,4],
    [5,6,7,8],
    [9,10,11,12],
    [13,14,15,16]
]

sprialMatrix(mat)



