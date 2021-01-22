#given an array we need to return true or false such that if there exist a tiple such that a+b+c is equal to the X
def findTriplet(arr,N,X):
    start=0
    arr.sort()
    for i in range(len(arr)-2):
        start=i+1
        end=len(arr)-1
        while start<end:
            current=arr[i]+arr[start]+arr[end]
            if current==X:
                return True
            elif current<X:
                start+=1
            else:
                end-=1
    return False


arr=[1,2,4,3,6]
ans=findTriplet(arr,5,10)
print(ans)