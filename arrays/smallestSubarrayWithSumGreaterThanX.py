#given an array we need to find the smallest whose sum is greater than x


def smallestSubArray(a,x):
    currentSum=0
    start=0
    currentLength=float('inf')
    for i in range(len(a)):
        currentSum+=a[i]
        if currentSum>x:
            while currentSum-a[start]>x:
                currentSum-=a[start]
                start+=1
            temp=i-start+1
            currentLength=min(currentLength,temp)
    if currentLength==float('inf'):
        currentLength=len(a)
    print(currentLength)



a=[1, 10, 5, 2, 7]
t=9
smallestSubArray(a,t)

