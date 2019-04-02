package linkedNumber

import (
	"fmt"
)

//每个node所记录数据的进位值
const max = 1e9

//链表节点
type node struct {
	previousNode *node
	number       int
	nextNode     *node
}

//链表
type LinkedNumber struct {
	firstNode *node
	lastNode  *node
}

//初始化链表
func Init(i int) LinkedNumber {
	firstNode, lastNode := buildNode(i)
	return LinkedNumber{firstNode: firstNode, lastNode: lastNode}
}

//构建节点
/**
传入一个初始值,如果这个值大于节点最大值1e9则创建多个节点并把它们串联,返回fisrtNode和lastNode
*/
func buildNode(i int) (*node, *node) {
	var firstNode *node
	var lastNode *node
	var tmpNode *node
	last := i % max
	lastNode = &node{previousNode: nil, number: last, nextNode: nil}
	firstNode = lastNode
	tmpNode = lastNode
	for previous := i / max; previous > 0; previous = previous / max {
		firstNode = &node{previousNode: nil, number: previous % max, nextNode: tmpNode}
		tmpNode.previousNode = firstNode
		tmpNode = firstNode
	}
	return firstNode, lastNode
}

//链表的结果值
func (l *LinkedNumber) String() string {
	return l.firstNode.String()
}

//打印建立节点结果字符串
func (n *node) String() string {
	var number string

	if n.number == 0 {
		//如果节点值为零且不是firstNode则是中间的节点,返回9个0,否则则输出0
		if n.previousNode != nil {
			number = "000000000"
		} else {
			number = "0"
		}
	} else {
		//如果节点不是firstNode则在数字前补0到9位
		var format string
		if n.previousNode != nil {
			format = "%09d"
		} else {
			format = "%d"
		}
		number = fmt.Sprintf(format, n.number)
	}
	//如果是最后一个节点则停止递归
	if n.nextNode == nil {
		return number
	} else {
		return fmt.Sprintf("%v%v", number, n.nextNode)
	}
}

//乘法
func (l *LinkedNumber) Multiply(i int) {
	tmpNode := l.lastNode
	plus := 0
	//利用乘法分配律的原理对每一个节点做乘法运算并把进位传给下一个节点
	for ; ; tmpNode = tmpNode.previousNode {
		tmpNode.number, plus = multiply(i, tmpNode.number, plus)
		if tmpNode.previousNode == nil {
			break
		}
	}
	//如果便利所有节点之后有剩余的进位.则将剩余的进位建立新的节点.并连接回原来的
	if plus != 0 && tmpNode.previousNode == nil {
		firstNode, lastNode := buildNode(plus)
		lastNode.nextNode = tmpNode
		tmpNode.previousNode = lastNode
		l.firstNode = firstNode
	}

}

//乘法换成n次循环的加法
func multiply(i int, number int, more int) (result int, plus int) {
	result = 0
	plus = 0
	for j := 0; j < i; j++ {
		result += number
		if result >= max {
			plus += result / max
			result %= max
		}
	}
	result += more
	if result >= max {
		plus += result / max
		result %= max
	}
	return
}
