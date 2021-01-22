#Given an array of chocolate we need to give 1 packet (packet may contain different chocolates inside them)to each student such that difference between the minimum chocolate given to the student and maximum chocolate given to the student should be minimum
import math
def main():
    n=int(input())
    for _ in range(n):
        solve()
def solve():
    _ = int(input())
    lst = list(map(int,input().strip().split()))
    student=int(input())
    lst.sort()
    ans=math.inf
    forward=student-1
    start=0
    while forward<len(lst):
        if lst[forward]-lst[start]<ans:
            ans=lst[forward]-lst[start]
        forward+=1
        start+=1
    print(ans)


if __name__=="__main__":
    main()






