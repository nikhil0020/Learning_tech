


def getSum(arr , n , i):
    if i == n :
        return 0
    return arr[i] + getSum(arr,n, i+1)


arr = [1,2,3,4,5]

n = len(arr)

print(getSum(arr,n,0))