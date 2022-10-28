def productOfArrayExceptSelf(nums) -> list[int]:
    n = len(nums)
    left = [1] * n
    right = [1] * n # If n == 3, this creates an array [1, 1, 1]
    for i in range(1, n):
        left[i] = left[i - 1] * nums[i - 1]
    for i in range(n - 2, -1, -1): # start at n-2, end at 0 (-1), decrement -1
        right[i] = right[i + 1] * nums[i + 1]
    result = []
    for i in range(n):
        result.append(left[i] * right[i])

    return result


def main():
    result = productOfArrayExceptSelf([1, 2, 3, 4, 5])
    print(result)

if __name__ == "__main__":
    main()