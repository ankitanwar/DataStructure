import math
def ProductSubarray(arr,n):
    result=arr[0]
    currentMax=arr[0]
    currentMin=arr[0]

    for i in range(1,len(arr)):
        t1=currentMin*arr[i]
        t2=currentMax*arr[i]
        t3=arr[i]
        currentMax=max(max(t1,t2),t3)
        currentMin=min(min(t1,t2),t3)
        result=max(result,currentMax)
    print(result)


arr=[8 ,-2, -2, 0, 8, 0, -6, -8, -6, -1]
ProductSubarray(arr,10)
