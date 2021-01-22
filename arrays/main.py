def maxProductSubarray(arr,n):
    ans=helper(arr,0,1,0)
    return ans


def helper(arr,currentIndex,currentVal,ans):
    if ans<currentVal:
        print("The value of currentVal is ",currentVal)
        currentVal=ans
    if currentIndex==len(arr):
        print("The value of ans is ",ans)
        return ans
    return max(helper(arr,currentIndex+1,currentVal*arr[currentIndex],ans),helper(arr,currentIndex+1,arr[currentIndex],ans))



arr=[6,-3,-10,0,2]
n=5
ans=maxProductSubarray(arr,n)
print(ans)