# BFS

BFS像一个眼观六路耳听八方的人、稳重的人。它是一层一层搜的。每次只会扩展一层，不会离家太远。当这一层全部扩展完以后，才会去扩展下一层。

- 求最小
- 基于迭代，不会爆栈。

# 模板

```
queue<int> q;
st[1] = true; // 表示1号点已经被遍历过
q.push(1);

while (q.size())
{
    int t = q.front();
    q.pop();

    for (int i = h[t]; i != -1; i = ne[i])
    {
        int j = e[i];
        if (!st[j])
        {
            st[j] = true; // 表示点j已经被遍历过
            q.push(j);
        }
    }
}
```

## 题目

- 845 八数码
- 844 走迷宫
- Flood fill类型问题
