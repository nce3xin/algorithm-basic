# 区间合并

## 应用场景

给很多区间，如果区间有交集，就合并。

如下图，5个蓝色区间，合并结果是2个绿色区间。注意：如果两个区间仅端点相交，也是可以合并的。

![](imgs/1.png)

## 思路

1. 按区间左端点排序。
2. 区间之间的关系只有3

![](imgs/2.png)

不会出现下图中红色的情况，因为区间是按左端点排序的。

![](imgs/3.png)

## 模板

```c++
// 将所有存在交集的区间合并
void merge(vector<PII> &segs)
{
    vector<PII> res;

    sort(segs.begin(), segs.end());

    int st = -2e9, ed = -2e9;
    for (auto seg : segs)
        if (ed < seg.first)
        {
            if (st != -2e9) res.push_back({st, ed});
            st = seg.first, ed = seg.second;
        }
        else ed = max(ed, seg.second);

    if (st != -2e9) res.push_back({st, ed});

    segs = res;
}
```
