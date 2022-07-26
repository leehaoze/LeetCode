package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const MaxLevel = 7

type Node struct {
	Val  int
	Prev *Node
	Next *Node
	Down *Node
}

type Skiplist struct {
	Level       int
	HeadNodeArr []*Node
}

func Constructor() Skiplist {
	return Skiplist{
		Level:       -1,
		HeadNodeArr: make([]*Node, MaxLevel),
	}
}

func (this *Skiplist) getNode(target int) *Node {
	// 跳表为空数据
	if this.Level == -1 {
		return nil
	}

	// 从最顶级的索引开始查找
	level := this.Level
	// 获取到第一个节点
	node := this.HeadNodeArr[level].Next
	// 从左向右找到第一个比要查找值大的node，从此节点的前一个节点向下走
	for node != nil {
		// 索引里正好就是这个值
		if node.Val == target {
			return node
		}

		// 节点值比target大了，那说明在当前节点和他的前一个节点之间
		if node.Val > target {
			// node.Prev 可能是头结点
			if node.Prev.Down == nil {
				// 下移一层
				if level-1 >= 0 {
					node = this.HeadNodeArr[level-1].Next
				} else {
					// 最底层了 节点不存在
					node = nil
				}
			} else {
				// 正常下移
				node = node.Prev.Down
			}
			level = level - 1
		} else if node.Val < target {
			// 正常右移
			node = node.Next
			// 右移值如果为nil，则本层查找完了, 去下一层
			if node == nil {
				level -= 1
				if level >= 0 {
					node = this.HeadNodeArr[level].Next
				}
			}
		}
	}

	return nil
}

func (this *Skiplist) Search(target int) bool {
	ret := this.getNode(target)
	if ret == nil {
		return false
	}

	return true
}

func (this *Skiplist) Add(num int) {
	insertPostionArr := make([]*Node, MaxLevel)
	// 寻找每一层的插入位置
	if this.Level >= 0 {
		level := this.Level
		node := this.HeadNodeArr[level].Next
		for node != nil && level >= 0 {
			if node.Val == num {
				insertPostionArr[level] = node
				if node.Down == nil {
					// 降一层
					if level-1 >= 0 {
						node = this.HeadNodeArr[level-1].Next
					} else {
						node = nil
					}
				} else {
					node = node.Down
				}
				level -= 1
			} else if node.Val > num {
				// 应该插入到此节点和上一个节点之间
				insertPostionArr[level] = node.Prev
				if node.Prev.Down == nil {
					// 降一层
					if level-1 >= 0 {
						node = this.HeadNodeArr[level-1].Next
					} else {
						node = nil
					}
				} else {
					node = node.Prev.Down
				}
				level -= 1
			} else if node.Val < num {
				// 本层查完 直接插到最后一个节点的位置
				if node.Next == nil {
					insertPostionArr[level] = node
					level -= 1
					if level >= 0 {
						node = this.HeadNodeArr[level].Next
					}
				} else {
					node = node.Next
				}
			}
		}
	}
	this.insert(num, insertPostionArr)
}

func (this *Skiplist) insert(val int, insertPostionArr []*Node) {
	// 插入到最底层
	node := &Node{
		Val:  val,
		Prev: nil,
		Next: nil,
		Down: nil,
	}

	// 空跳表
	if this.Level < 0 {
		this.Level = 0
		this.HeadNodeArr[0] = &Node{}
		this.HeadNodeArr[0].Next = node
		node.Prev = this.HeadNodeArr[0]
		return
	}

	// 每一层尝试插入 从底层开始
	rootNode := insertPostionArr[0]
	nextNode := rootNode.Next
	rootNode.Next = node
	node.Prev = rootNode
	node.Next = nextNode
	if nextNode != nil {
		nextNode.Prev = node
	}

	currentLevel := 1
	// 通过摇点决定要不要再更高层插入
	for randLevel() && currentLevel <= this.Level+1 && currentLevel < MaxLevel {
		if insertPostionArr[currentLevel] == nil {
			rootNode = &Node{}
			this.HeadNodeArr[currentLevel] = rootNode
		} else {
			rootNode = insertPostionArr[currentLevel]
		}

		nextNode = rootNode.Next
		upNode := &Node{
			Val:  val,
			Prev: rootNode,
			Next: nextNode,
			Down: node,
		}
		rootNode.Next = upNode
		if nextNode != nil {
			nextNode.Prev = upNode
		}

		node = upNode
		currentLevel++
	}

	for i := 0; i < MaxLevel; i++ {
		if this.HeadNodeArr[i] == nil {
			this.Level = i - 1
			break
		}
	}

}

// 通过抛硬币决定是否加入下一层
func randLevel() bool {
	randNum := rand.Intn(2)
	if randNum == 0 {
		return true
	}
	return false
}

func (this *Skiplist) Erase(num int) bool {
	node := this.getNode(num)
	if node == nil {
		return false
	}

	for node != nil {
		prev := node.Prev
		next := node.Next

		prev.Next = next
		if next != nil {
			next.Prev = prev
		}
		node = node.Down
	}

	// 清理
	for i := 0; i < MaxLevel; i++ {
		// 此层空了
		if this.HeadNodeArr[i] == nil {
			this.Level = i - 1
			break
		}
		if this.HeadNodeArr[i].Next == nil {
			this.HeadNodeArr[i] = nil
			this.Level = i - 1
			break
		}
	}
	return true
}

func print(list *Skiplist) {
	fmt.Println("====================start===============" + strconv.Itoa(list.Level))
	for i := list.Level; i >= 0; i-- {
		node := list.HeadNodeArr[i].Next
		for node != nil {
			fmt.Print(strconv.Itoa(node.Val) + " -> ")
			node = node.Next
		}
		fmt.Println()
	}
	fmt.Println("====================end===============")
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

func main() {
	skipList := Constructor()

	oplist := []string{"Skiplist", "add", "add", "add", "add", "add", "add", "add", "add", "add", "erase", "search", "add", "erase", "erase", "erase", "add", "search", "search", "search", "erase", "search", "add", "add", "add", "erase", "search", "add", "search", "erase", "search", "search", "erase", "erase", "add", "erase", "search", "erase", "erase", "search", "add", "add", "erase", "erase", "erase", "add", "erase", "add", "erase", "erase", "add", "add", "add", "search", "search", "add", "erase", "search", "add", "add", "search", "add", "search", "erase", "erase", "search", "search", "erase", "search", "add", "erase", "search", "erase", "search", "erase", "erase", "search", "search", "add", "add", "add", "add", "search", "search", "search", "search", "search", "search", "search", "search", "search"}
	val := []int{0, 16, 5, 14, 13, 0, 3, 12, 9, 12, 3, 6, 7, 0, 1, 10, 5, 12, 7, 16, 7, 0, 9, 16, 3, 2, 17, 2, 17, 0, 9, 14, 1, 6, 1, 16, 9, 10, 9, 2, 3, 16, 15, 12, 7, 4, 3, 2, 1, 14, 13, 12, 3, 6, 17, 2, 3, 14, 11, 0, 13, 2, 1, 10, 17, 0, 5, 8, 9, 8, 11, 10, 11, 10, 9, 8, 15, 14, 1, 6, 17, 16, 13, 4, 5, 4, 17, 16, 7, 14, 1}
	for idx, op := range oplist {
		var ret bool

		// print(&skipList)

		if op == "add" {
			skipList.Add(val[idx])
		} else if op == "search" {
			ret = skipList.Search(val[idx])
		} else if op == "erase" {
			ret = skipList.Erase(val[idx])
		}
		fmt.Println(fmt.Sprintf("idx: %v, op: %v, val: %v, ret: %v", idx, op, val[idx], ret))
	}
}
