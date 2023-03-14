# Floyd

这个算法超简单。求多源汇最短路。用邻接矩阵来存储图。就是三重循环。原理是基于动态规划。

循环的时候注意一定要先循环k，i和j的顺序随便颠倒。

重边取最短的那条，自环直接把环删了。

如果出现segment fault，采用删代码法，看是哪块的问题。

## 时间复杂度

O(n^3), n表示点数。

## 模板

```
初始化：
    for (int i = 1; i <= n; i ++ )
        for (int j = 1; j <= n; j ++ )
            if (i == j) d[i][j] = 0;
            else d[i][j] = INF;

// 算法结束后，d[a][b]表示a到b的最短距离
void floyd()
{
    for (int k = 1; k <= n; k ++ )
        for (int i = 1; i <= n; i ++ )
            for (int j = 1; j <= n; j ++ )
                d[i][j] = min(d[i][j], d[i][k] + d[k][j]);
}
```
