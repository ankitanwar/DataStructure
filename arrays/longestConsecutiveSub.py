def findLongestConsecutiveSubsequence(arr,N):
    element={}
    ans=0
    for val in  arr:
        element[val]=False
    for key in element:
        t=0
        current=key
        backward=key
        element[key]=True
        while current+1 in element and element[current+1]==False:
            element[current+1]=True
            t+=1
            current+=1
        while backward-1 in element and element[backward-1]==False:
            element[backward-1]=False
            t+=1
            backward-=1
        if t!=0:
            t+=1
        ans=max(ans,t)
    return ans
