package main

import "container/list"

type Node struct {
	keys  map[string]struct{}
	count int
}

type AllOne struct {
	*list.List
	nodes map[string]*list.Element
}

func Constructor() AllOne {
	return AllOne{
		List:  list.New(),
		nodes: map[string]*list.Element{},
	}
}

func (this *AllOne) Inc(key string) {
	if cur := this.nodes[key]; cur != nil { // 在链表中
		curNode := cur.Value.(Node)
		// 后边为空，直接追加到后边
		if nxt := cur.Next(); nxt == nil || nxt.Value.(Node).count > curNode.count+1 {
			this.nodes[key] = this.InsertAfter(Node{
				keys: map[string]struct{}{
					key: {},
				},
				count: curNode.count + 1,
			}, cur)
		} else { // 后边不为空 并且 后一个的count是 等于 当前count+1的 合并到后边的count里
			nxt.Value.(Node).keys[key] = struct{}{}
			this.nodes[key] = nxt
		}
		delete(curNode.keys, key)
		if len(curNode.keys) == 0 {
			this.Remove(cur)
		}
	} else { // 不在链表
		if this.Front() == nil || this.Front().Value.(Node).count > 1 { // 插入到前边
			this.nodes[key] = this.PushFront(Node{
				keys: map[string]struct{}{
					key: {},
				},
				count: 1,
			})
		} else {
			// 合并到前边
			this.Front().Value.(Node).keys[key] = struct{}{}
			this.nodes[key] = this.Front()
		}
	}
}

func (this *AllOne) Dec(key string) {
	cur := this.nodes[key]
	curNode := cur.Value.(Node)
	if curNode.count > 1 {
		if pre := cur.Prev(); pre == nil || pre.Value.(Node).count < curNode.count-1 { // 插入到前一个里
			this.nodes[key] = this.InsertBefore(Node{
				keys: map[string]struct{}{
					key: {},
				},
				count: curNode.count - 1,
			}, cur)
		} else {
			pre.Value.(Node).keys[key] = struct{}{}
			this.nodes[key] = pre
		}

	} else {
		delete(this.nodes, key)
	}
	delete(curNode.keys, key)
	if len(curNode.keys) == 0 {
		this.Remove(cur)
	}
}

func (this *AllOne) GetMaxKey() string {
	if b := this.Back(); b != nil {
		for key := range b.Value.(Node).keys {
			return key
		}
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if b := this.Front(); b != nil {
		for key := range b.Value.(Node).keys {
			return key
		}
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
