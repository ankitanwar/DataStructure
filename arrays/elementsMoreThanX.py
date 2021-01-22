#we have to return all the elements which have occured more than n\k times
import collections

def moreThanN(arr,n,k):
    x=collections.Counter(arr)
    count=0
    t=int(n/k)
    for keys in x:
        if x[keys]>t:
            count+=1
    return count



arr=[2,3,3,2]
n=4
k=3
ans=moreThanN(arr,n,k)
print(ans)