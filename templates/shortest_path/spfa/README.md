# spfa算法

## 原理

其实是对bellman-ford算法做了优化。用宽搜做优化。

终于来到SPFA算法了！之前已经说明过了Bellman_ford算法 ,我们今天说明的SPFA算法仅仅只是对该算法的一个优化。

Bellman_ford算法会遍历所有的边，但是有很多的边遍历了其实没有什么意义，我们只用遍历那些到源点距离变小的点所连接的边即可，只有当一个点的前驱结点更新了，该节点才会得到更新；因此考虑到这一点，我们将创建一个队列每一次加入距离被更新的结点。

虽然说正权图用dijkstra, 但是其实也可以用spfa做，很多都可以用spfa过掉。

## 基本思想

更新过谁，我再拿谁来更新别人。

![](imgs/1.png)

dist[b]变小，一定是因为dist[a]变小了。也就是，只有a变小，b才会变小。只有a变小了，后继才会变小。所有用宽搜来做优化。

## 模板

时间复杂度，平均情况下 O(m)，最坏情况下 O(n*m), n表示点数，m 表示边数。

```
int n;      // 总点数
int h[N], w[N], e[N], ne[N], idx;       // 邻接表存储所有边
int dist[N];        // 存储每个点到1号点的最短距离
bool st[N];     // 存储每个点是否在队列中

// 求1号点到n号点的最短路距离，如果从1号点无法走到n号点则返回-1
int spfa()
{
    memset(dist, 0x3f, sizeof dist);
    dist[1] = 0;

    queue<int> q;
    q.push(1);
    st[1] = true;

    while (q.size())
    {
        auto t = q.front();
        q.pop();

        st[t] = false;

        for (int i = h[t]; i != -1; i = ne[i])
        {
            int j = e[i];
            if (dist[j] > dist[t] + w[i])
            {
                dist[j] = dist[t] + w[i];
                if (!st[j])     // 如果队列中已存在j，则不需要将j重复插入
                {
                    q.push(j);
                    st[j] = true;
                }
            }
        }
    }

    if (dist[n] == 0x3f3f3f3f) return -1;
    return dist[n];
}
```

## 题目

- 851
