# 堆优化的Dijkstra

## 优化思路

看一下朴素算法的时间复杂度：

```
for(i:1 ~ n)//n次
{
    t <- 没有确定最短路径的节点中距离源点最近的点;//每次遍一遍历dist数组，n次的复杂度是O(n^2)
    state[t] = 1;
    更新 dist;//每次遍历一个节点的出边，n次遍历了所有节点的边，复杂度为O(e)
}
```

算法的主要耗时的步骤是从dist 数组中选出：没有确定最短路径的节点中距离源点最近的点 t。只是找个最小值而已，没有必要每次遍历一遍dist数组。

在一组数中每次能很快的找到最小值，很容易想到使用小根堆。可以使用库中的小根堆（推荐）或者自己编写。

## 模板

时间复杂度 O(mlogn)，n 表示点数，m 表示边数。

```
typedef pair<int, int> PII;

int n;      // 点的数量
int h[N], w[N], e[N], ne[N], idx;       // 邻接表存储所有边
int dist[N];        // 存储所有点到1号点的距离
bool st[N];     // 存储每个点的最短距离是否已确定

// 求1号点到n号点的最短距离，如果不存在，则返回-1
int dijkstra()
{
    memset(dist, 0x3f, sizeof dist);
    dist[1] = 0;
    priority_queue<PII, vector<PII>, greater<PII>> heap;
    heap.push({0, 1});      // first存储距离，second存储节点编号

    while (heap.size())
    {
        auto t = heap.top();
        heap.pop();

        int ver = t.second, distance = t.first;

        if (st[ver]) continue;
        st[ver] = true;

        for (int i = h[ver]; i != -1; i = ne[i])
        {
            int j = e[i];
            if (dist[j] > distance + w[i])
            {
                dist[j] = distance + w[i];
                heap.push({dist[j], j});
            }
        }
    }

    if (dist[n] == 0x3f3f3f3f) return -1;
    return dist[n];
}
```

## 题目

- 850
