package leetcode

//其实就是一个双向链表+map
//为什么需要双向链表呢，因为需要从队首+队尾快速操作

// 定义一个双向链表的结构
// 这个双向链表中的值有key,value
type DLinkedNode struct {
	key, value int
	prev       *DLinkedNode
	next       *DLinkedNode
}

type LRUCache struct {
	head *DLinkedNode
	tail *DLinkedNode
	//根据值获取节点
	cacheMap map[int]*DLinkedNode
	capacity int //上限容量
	size     int //当前存储容量
}

func initNode(k, v int) *DLinkedNode {
	return &DLinkedNode{
		key:   k,
		value: v,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
		capacity: capacity,
		cacheMap: map[int]*DLinkedNode{},
	}
	//定义好双向链表
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cacheMap[key]
	if !ok {
		return -1
	}
	this.removeToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cacheMap[key]
	//注意下历史node的复用

	if !ok {
		newNode := initNode(key, value)
		//如果不存在，则直接插入首位
		this.addToHead(newNode)
		this.size++
		this.cacheMap[key] = newNode
		if this.size > this.capacity {
			//如果容量超了
			deleteNode := this.removeTail()
			delete(this.cacheMap, deleteNode.key)
			this.size--
		}
	} else {
		node.value = value
		this.removeToHead(node)
		this.cacheMap[key] = node
	}

}

//其实一些细节也挺有意思的

// 记录双向链表的一些操作，是否需要删除由业务层去决定
// 需要定义一些双向链表的操作，比如头部插入数据，尾部删除数据
func (this *LRUCache) addToHead(node *DLinkedNode) {
	//新插入的节点的指针的指向
	node.next = this.head.next
	node.prev = this.head

	//头部节点的指针走向
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeFromTail() {
	this.tail.prev.next = this.tail.next
	this.tail.next.prev = this.tail.prev
}

// 需要返回相关的节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// 删除节点
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 要删除里面的节点
func (this *LRUCache) removeToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
