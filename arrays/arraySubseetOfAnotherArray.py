import collections

def main():
    n=int(input())
    for _ in range(n):
        solve()
def solve():
    notFound=True
    _,_ = map(int,input().strip().split())
    arr1 = list(map(int,input().strip().split()))
    arr2 = list(map(int,input().strip().split()))
    x=collections.Counter(arr1)
    for val in arr2:
        if x.get(val)==None:
            notFound=False
            print("No")
            break
    if notFound==True:
        print("Yes")
if __name__=="__main__":
    main()
