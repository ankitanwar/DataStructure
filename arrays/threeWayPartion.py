"""
Given an array we need to partion that such that
all elements smaller than low value comes first
all elements smaller than low value and larger then high value should come in the mid 
all elements larger then high value should come in the last
"""

def partionAroundRange(array,a,b):
    first=-1
    last=len(array)
    i=0
    while i<len(array):
        if array[i]<a:
            first+=1
            array[i],array[first]=array[first],array[i]
            i+=1
        elif array[i]>b:
            last-=1
            array[i],arr[last]=array[last],array[i]
        else:
            i+=1
    print(array)


arr=[1, 14, 5, 20, 4, 2, 54, 20, 87, 98, 3, 1, 32]
partionAroundRange(arr,14,20)

