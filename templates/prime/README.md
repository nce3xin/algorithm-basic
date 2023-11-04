# 质数

在大于1的整数中，如果只包含1和本身这两个约数，就被称为质数，或者叫素数。

## 质数的判定

试除法。

最简单写法，从定义出发，时间复杂度O(n)，效率比较低:

```
func isPrime(n int) bool{
    if n < 2 {
        return false
    }
    for i:=2;i<n;i++{
        if n % i == 0{
            return false
        }
    }
    return true
}
```

可以对其做优化。优化有一个很重要的性质，n的所有约数都是成对出现的，一个数 x 分解成两个数的乘积，则这两个数中，一定有一个数大于根号 x，一个数小于根号x。在枚举的时候，可以只枚举每对当中较小的那个。

## 模板

```
bool is_prime(int x)
{
    if (x < 2) return false;
    // 推荐写法 i <= x / i. 不推荐写法：i * i <= x. 因为如果i很大，i^2很容易溢出，超过int范围
    for (int i = 2; i <= x / i; i ++ )
        if (x % i == 0)
            return false;
    return true;
}
```

## 题目

- 866 试除法判定质数
