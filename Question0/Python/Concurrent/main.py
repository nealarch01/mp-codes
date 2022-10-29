import threading

def leftProducts(nums, left) -> list[int]:
    n = len(nums)
    for i in range(1, n):
        left[i] = left[i - 1] * nums[i - 1]
    return left

def rightProducts(nums, right) -> list[int]:
    n = len(nums)
    for i in range(n - 2, -1, -1): # start at n-2, end at 0 (-1), decrement -1
        right[i] = right[i + 1] * nums[i + 1]
    return right

def productOfArrayExceptSelf(nums) -> list[int]:
    n = len(nums)
    left = [1] * n
    right = [1] * n # If n == 3, this creates an array [1, 1, 1]

    thread1 = threading.Thread(target = leftProducts, args = (nums, left))
    thread2 = threading.Thread(target = rightProducts, args = (nums, right))

    thread1.start()
    thread2.start()

    thread1.join()
    thread2.join() # Wait for the thread to join back to the main thread

    result = []
    for i in range(n):
        result.append(left[i] * right[i])

    return result


def main():
    result = productOfArrayExceptSelf([1, 2, 3, 4, 5])
    print(result)

if __name__ == "__main__":
    main()