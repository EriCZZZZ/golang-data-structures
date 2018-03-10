# README

<!-- TOC -->

- [README](#readme)
- [AVL树](#avl树)
    - [简介](#简介)
    - [性质](#性质)
    - [效率](#效率)
    - [实现](#实现)
        - [节点构成](#节点构成)
        - [插入](#插入)
            - [两个定义](#两个定义)
            - [原理](#原理)
- [二叉堆 *以小顶堆为例*](#二叉堆-以小顶堆为例)
    - [特性](#特性)
        - [结构特性](#结构特性)
        - [堆序特性](#堆序特性)
    - [实现原理](#实现原理)
    - [操作](#操作)
        - [插入](#插入-1)
            - [原理](#原理-1)
            - [细节](#细节)
        - [去除最小值](#去除最小值)
            - [原理](#原理-2)
            - [细节](#细节-1)

<!-- /TOC -->

# AVL树

## 简介
自平衡二叉查找树

## 性质
1. 左右子树的高度差绝对值不超过1。
2. 左右子树均为AVL树。

## 效率
插入/删除/查找均为O(logN)

## 实现
每个节点包含一个平衡因子 *(balance factor)*

### 节点构成
```golang
type Node struct {
	Val int
	L *Node
	R *Node
	bf int		// balance factory 定义为左子树高度减去右子树高度
}
```

### 插入

#### 两个定义
1. 旋转 由于AVL树的左右子树高度差小于等于1，每次插入时都可能需要通过旋转来维持两子树的高度差。旋转操作共分四种。
    - 左旋
    - 右旋
    - 左旋后右旋
    - 右旋后左旋

	![AVL旋转图示](./github.com/ek1o/collections/AVL/AVLRotate.png)

2. 平衡因子 *balance factor* 平衡因子为两子树的高度差，**只能**是-2 -1 0 1 2.

    > 约定：实现中平衡因子=左子树高度-右子树高度

#### 原理
1. 若根节点为空，初始化根节点，返回；若值与根节点值相同，返回。
2. 若值大于根节点值，对右子树递归调用插入函数；若小于，对左子树调用。
3. 刷新平衡因子。
    - 左子树从无到有或插入到左子树且左子树的平衡因子从0变为1/-1，则根节点平衡因子加一。
    - 右子树从无到有或插入到右子树且右子树的平衡因子从0变为1/-1，则根节点平衡因子减一。
4. 检测平衡因子，如果需要，则通过四种旋转恢复平衡。

    | 条件（插入在） |         平衡因子        |     操作     |    旋转后的平衡因子    |
    |:--------------:|:-----------------------:|:------------:|:----------------------:|
    | 左子树的左子树 |       BF/L.BF==2/1      |   直接右旋   |      BF/L.BF==0/0      |
    | 左子树的右子树 |  BF/L.BF/L.R.BF==2/-1/1 | 先左旋后右旋 | BF/L.BF/L.R.BF==-1/0/0 |
    |                | BF/L.BF/L.R.BF==2/-1/-1 |              |  BF/L.BF/L.R.BF==0/1/0 |
    | 右子树的左子树 |  BF/R.BF/R.L.BF==-2/1/1 | 先右旋再左旋 | BF/R.BF/R.L.BF==0/-1/0 |
    |                | BF/R.BF/R.L.BF==-2/1/-1 |              |  BF/R.BF/R.L.BF==1/0/0 |
    | 右子树的右子树 |      BF/R.BF==-2/-1     |   直接左旋   |      BF/R.BF==0/0      |

# 二叉堆 *以小顶堆为例*

## 特性

### 结构特性

- 一个 **完全二叉树**

### 堆序特性

- 对于任意非叶节点，其值小于任意子节点 **此处的小于是因为小顶堆的原因，大顶堆则为大于**

## 实现原理

- 对于堆通常使用数组具体实现 因为 **对于节点i来说 _i是idx_ 左子节点的idx是2i 右子节点的idx是2i+1**
- 由于上一条，有两结论
    - 堆的大小需要预估
    - 数组idx==0需要空出，根节点的idx为1

## 操作

### 插入

#### 原理

 上滤 首先将新值直接放入最后 比较 新值与父节点的大小 如果新值更小 则交换两者 重复该过程

#### 细节

1. 上滤的结束条件为 **上滤到根节点或父节点小于新值**
1. 上滤过程中 为了减少交换次数 实现中使用假装将新值放入的方法

### 去除最小值

#### 原理

下滤 将 **最后一个值** 提取出来 然后递归的，从根节点开始，取较小的子节点填入当前节点，直至能将 **最后一个值** 填入为止

#### 细节

1. 下滤的结束条件为 当前节点为叶节点或者最小子节点大于 **最后一个值**
1. 下滤过程中需要处理单子节点的情况 通过如下trick避免讨论
```golang
if idx * 2 + 1 < t.pIdx && min > t.heap[idx * 2 + 1] {
	min = t.heap[idx * 2 + 1]
	minIdx = idx * 2 + 1
}
```