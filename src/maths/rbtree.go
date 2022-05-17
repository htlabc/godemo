package maths

import "errors"

const (
	RED   bool = true
	BLACK bool = false
)

const (
	LEFTROTATE  bool = true
	RIGHTROTATE bool = false
)

type RBNODE struct {
	value               uint64
	color               bool
	left, right, parent *RBNODE
}

//初始化红黑树
func NewRBNode(value uint64) *RBNODE {
	return &RBNODE{color: RED, value: value}
}

func (rbnode *RBNODE) getGrandParent() *RBNODE {
	if rbnode.parent == nil {
		return nil
	}

	return rbnode.parent.parent
}

//获取兄弟节点

func (rbnode *RBNODE) getBrother() *RBNODE {
	if rbnode.parent == nil {
		return nil
	}

	if rbnode == rbnode.parent.right {
		return rbnode.parent.left
	} else {
		return rbnode.parent.right
	}

}

func (rbnode *RBNODE) rotate(isRotateleft bool) (*RBNODE, error) {

	var root *RBNODE
	if rbnode == nil {
		return root, nil
	}

	if !isRotateleft && rbnode.left == nil {
		return root, errors.New("右旋节点不能为空")
	} else if isRotateleft && rbnode.right == nil {
		return root, errors.New("左旋节点不能为空")
	}

	parent := rbnode.parent

	var isleft bool
	if parent != nil {
		isleft = parent.left == rbnode
	}
	//左旋
	if isRotateleft {
		grandson := rbnode.right.left
		rbnode.right.left = rbnode
		rbnode.parent = rbnode.right
		rbnode.right = grandson
	} else {
		//右旋
		grandson := rbnode.left.right
		rbnode.left.right = rbnode
		rbnode.parent = rbnode.left
		rbnode.left = grandson
	}

	if parent == nil {
		rbnode.parent.parent = nil
		root = rbnode.parent
	} else {
		if isleft {
			parent.left = rbnode.parent
		} else {
			parent.right = rbnode.parent
		}
		rbnode.parent.parent = parent
	}
	return root, nil
}

func (rbnode *RBNODE) getMostLeftChild() *RBNODE {
	if rbnode.left == nil {
		return rbnode
	}
	return rbnode.left.getMostLeftChild()
}
