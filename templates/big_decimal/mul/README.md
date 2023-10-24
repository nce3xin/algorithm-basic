# 大整数乘法

高精度乘法，一般考的是，一个大整数和一个普通整数相乘，不会考两个大整数相乘。

注意，下面用竖式乘法模拟，是把 B 作为一个整体，和 A 的每一位相乘。和之前教的不太一样。

假设 t 是进位，那么 C[i] = （A[i]*B[i]+t）% 10。

![](imgs/mul.png)

## 模板

```
// C = A * b, A >= 0, b >= 0
vector<int> mul(vector<int> &A, int b)
{
    vector<int> C;

    int t = 0;
    for (int i = 0; i < A.size() || t; i ++ )
    {
        if (i < A.size()) t += A[i] * b;
        C.push_back(t % 10);
        t /= 10;
    }

    while (C.size() > 1 && C.back() == 0) C.pop_back();

    return C;
}
```

## 题目

- 793
