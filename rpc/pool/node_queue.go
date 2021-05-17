package pool

import (
	"sync"
)

// NodeQueue 管理NetNode的队列
type NodeQueue interface {
	Get()*NetNode
	Take()*NetNode
	Put(node *NetNode)
	Len()int
	Grow(size int)
	Clean()
}

type blockNodeQueue struct {
	container     chan *NetNode
	tempContainer chan *NetNode
	mutex         sync.Mutex
	once          sync.Once
}

func (queue *blockNodeQueue) Len() int {
	return len(queue.container)
}

func NewBlockNodeQueue(size int) *blockNodeQueue {
	return &blockNodeQueue{container: make(chan *NetNode,size)}
}

// Get 获取连接(非阻塞)
func (queue *blockNodeQueue) Get()*NetNode {
	if queue.tempContainer == nil {
		select {
		case node := <-queue.container:
			return node
		default:
			return nil
		}
	}else {
		select {
		case node := <-queue.container:
			return node
		case node := <-queue.tempContainer:
			return node
		default:
			return nil
		}
	}
}

// Take 获取连接(阻塞)
func (queue *blockNodeQueue) Take() *NetNode {
	select {
	case node := <-queue.container:
		return node
	case node := <-queue.tempContainer:
		return node
	}
}

// Put 添加连接(阻塞)
func (queue *blockNodeQueue) Put(node *NetNode) {
	if node.IsTemp() {
		queue.tempContainer <- node
	}else {
		queue.container <- node
	}
}

// Grow 队列扩充，新建chan存放临时连接
func (queue *blockNodeQueue) Grow(size int) {
	queue.once.Do(func() {
		queue.tempContainer = make(chan *NetNode,size)
	})
}

// Clean 临时连接清理
func (queue *blockNodeQueue) Clean() {
	if len(queue.tempContainer) == 0 {
		return
	}
	var count = 0
	select {
	case node := <- queue.tempContainer:
		if node.CanRecycle() {
			node.Close()
		}else {
			count ++
			queue.tempContainer <- node
			if count >= cap(queue.tempContainer)/2{
				return
			}
		}
	default:
		if count >= cap(queue.tempContainer)/2 {
			return
		}
	}
}