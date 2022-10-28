// Corban's code:

#include <iostream>
#include <vector>

using namespace std;

vector<int> productOfArrayExceptSelf(vector<int> &nums) {
    vector<int> lar = {};
    vector<int> rar = {};
    for(int i = 0; i < nums.size(); i++)
    {
        lar.push_back(1);
        rar.push_back(1);
    }

    for(int i = 1; i < nums.size(); i++)
    {
        lar[i] = nums[i - 1] * lar[i - 1];
    }
    for(int i = nums.size() - 2; i  >= 0; i--)
    {
        rar[i] = rar[i + 1] * nums[i + 1];
    }
    for(int i = 0; i < nums.size(); i++)
    {
        nums[i] = lar[i] * rar[i];
    }
    return nums;
}

int main() {
    vector<int> v;
    for (int i = 1; i <= 5; i++) {
        v.push_back(i);
    }

    vector<int> result = productOfArrayExceptSelf(v);
    for (int i = 0; i < result.size(); i++) {
        cout << v[i] << " ";
    }
}