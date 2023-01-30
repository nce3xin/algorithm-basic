# 树和图的遍历

## 注意

BFS和DFS遍历的时候，一定要记住，每个点只会遍历一次。所以一般要开一个bool数组，存一下哪些点已经被遍历过了。

## DFS

时间复杂度 O(n+m)，n 表示点数，m 表示边数。

### DFS模板

```
int dfs(int u)
{
    st[u] = true; // st[u] 表示点u已经被遍历过

    for (int i = h[u]; i != -1; i = ne[i])
    {
        int j = e[i];
        if (!st[j]) dfs(j);
    }
}
```

### DFS题目

- 846 树的重心

## BFS

### BFS模板

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

### BFS题目

- 847 图中点的层次
