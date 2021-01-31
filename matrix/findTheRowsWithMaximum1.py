#row is sorted and consist of 0 and 1 and we need to calc the row with max number of 1
class Solution:
    def calcOne(self,matrix,low,high):
        if high>=low:
            mid=int((low+high)/2)
            if (mid==0 or matrix[mid-1]==0) and matrix[mid]==1:
                return mid
            elif matrix[mid]==1:
                return self.calcOne(matrix,low,mid-1)
            else:
                return self.calcOne(matrix,mid+1,high)
        return -1
    def rowWithMax1s(self,arr, n, m):
		# code here
        currentAns=0
        currentRow=-1
        length=len(arr[0])-1
        for i in range(len(arr)):
            if currentAns==m:
                return currentRow
            t=length-currentAns
            if arr[i][t]==1:
                t2=self.calcOne(arr[i],0,length)
                if length+1-t2>currentAns:
                    currentAns=length+1-t2
                    currentRow=i
        return currentRow
mat=[
    [0,0,1,1,1]
]
s=Solution()
ans=s.rowWithMax1s(mat,1,5)
print(ans)