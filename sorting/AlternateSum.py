def calc(arr,index):
    print("i am working")
    if index>=len(arr):
        return 0
    else:
        return max(arr[index]+calc(arr,index+2),calc(arr,index+1))

arr = [5, 3, 4, 11, 2]
index=0
ans=calc(arr,0)
print(ans)