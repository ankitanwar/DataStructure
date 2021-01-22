#at any point we need to find the greatest to the left and greater to the right
def trappingWater(arr,n):
    greaterToLeft=[]
    greaterToRight=[]
    #greater to the left
    greaterToLeft.append(0)
    stack=[]
    stack.append(arr[0])
    for i in range(1,len(arr)):
        while len(stack)!=0 and stack[len(stack)-1]<=arr[i]:
            stack=stack[:len(stack)-1]
        if len(stack)==0:
            stack.append(arr[i])
            greaterToLeft.append(0)
        else:
            greaterToLeft.append(stack[len(stack)-1])
    greaterToRight.append(0)
    stack=stack[:0]
    stack.append(arr[len(arr)-1])
    for i in range(len(arr)-2,-1,-1):
        while len(stack)!=0 and stack[len(stack)-1]<=arr[i]:
            stack=stack[:len(stack)-1]
        if len(stack)==0:
            stack.append(arr[i])
            greaterToRight.append(0)
        else:
            greaterToRight.append(stack[len(stack)-1])
    greaterToRight=greaterToRight[::-1]
    ans=0
    for i in range(len(greaterToRight)):
        current=min(greaterToRight[i],greaterToLeft[i])
        if current!=0:
            temp=current-arr[i]
            ans+=temp
    return ans


arr=[6,9,9]
trappingWater(arr,3)