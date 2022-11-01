# Question 0 
## Question: Product of Array Except Self


Implemented in:
- C++ (Sequential and Concurrent)
- Python (Sequential and Concurrent)
- Go (Sequential and Concurrent)

### C++
Sequential: ```g++ -std=c++11 main.cpp```
Concurrent: ```g++ -std=c++11 -lpthread main.cpp```

### Go:
Sequential and concurrent: ```go run main.go```

### Python:
Sequential and concurrent: ```python3 main.py```


#### Run down of the solution:<br>
Build an array that contains the sequence of products. One for the left, one for the right.

For example: given ```arr = [1, 2, 3, 4, 5]``` <br>
At ```i = 2 or arr[2]```, we can utilize the sequence of left array and right array. (Note: when i = 2, we are attempting to find the product of 1 * 2 * 4 * 5)<br>
```left_array = [1, 1, 2, 6, 24]```<br>
```right_array = [120, 60, 20, 5, 1]```<br>
The sequence of the left array when i = 2 is 2<br>
The sequence of the right array when i = 2 is 20<br>
Notice how 2 <=> (1 * 2) and 20 <=> (4 * 5), this is exactly the same as (1 * 2 * 4 * 5) which was mentioned above.<br>
Two last things to mention: 
1. The very first element of the sequence array should always begin with one. That is because the first and last index is the sequence of the right/left multiplied by one.

2. If you are appending to the right array, you will need to reverse your iteration or the right array to obtain the product.