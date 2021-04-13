package memredis

type Node struct {
	Score  int
	Values []interface{}
	Left   *Node
	Right  *Node
	Bf     int
	Parent *Node
}

//平衡二叉树，也叫 AVL 树（平衡二叉树作者的名字首字母），是自平衡的二叉查找树，
//要求每个节点的左子树和右子树的高度差至多等于 1，
//这个高度（深度）差的值叫做平衡因子 BF，也就是说 BF 的值不能大于1，
//距离插入节点最近的，且平衡因子绝对值大于 1 的节点为根的子树，叫做最小不平衡子树，
//一旦出现最小不平衡子树时，就进行左旋、右旋或双旋处理，以保持自身始终平衡
//算法复杂度：O(logn)
type AvlTree struct {
	Tree *Node
}

const (
	LH = 1
	EH = 0
	RH = -1
)

//中序遍历
func (t *AvlTree) MidOrderTraverse() (result []interface{}) {
	result = append(result, t.Tree.MidOrderTraverse()...)
	return
}

//中序遍历
func (t *Node) MidOrderTraverse() (result []interface{}) {
	if t == nil {
		return
	}
	result = append(result, t.Left.MidOrderTraverse()...)
	result = append(result, t.Values...)
	result = append(result, t.Right.MidOrderTraverse()...)
	return
}

//先序遍历
func (t *AvlTree) PreOrderTraverse() (result []interface{}) {
	result = append(result, t.Tree.PreOrderTraverse()...)
	return
}

//先序遍历
func (t *Node) PreOrderTraverse() (result []interface{}) {
	if t == nil {
		return
	}

	result = append(result, t.Values...)
	result = append(result, t.Left.PreOrderTraverse()...)
	result = append(result, t.Right.PreOrderTraverse()...)
	return
}

//后序遍历
func (t *AvlTree) PostOrderTraverse() (result []interface{}) {
	result = append(result, t.Tree.PreOrderTraverse()...)
	return
}

//后序遍历
func (t *Node) PostOrderTraverse() (result []interface{}) {
	if t == nil {
		return
	}
	result = append(result, t.Left.PostOrderTraverse()...)
	result = append(result, t.Right.PostOrderTraverse()...)
	result = append(result, t.Values...)
	return
}

func (t *AvlTree) Insert(score int, v interface{}) {
	t.InsertNode(score, v)
	t.Reset() //将指针恢复到整棵树的根节点处
}

//查找并返回节点
func (t *AvlTree) Slice(startscore, leng int) (result []interface{}) {
	p := t.Tree
	for {
		if startscore <= p.Score {

		}
	}
}

func (t *AvlTree) SearchNode(score int) (me *Node) {
	me = t.Tree

	for {
		if me == nil {
			break
		}
		if me.Score == score {
			break
		}
		if me.Score < score {
			me = me.Right
		} else {
			me = me.Left
		}
	}

	return
}

//SearchNodeOrLeft score==me.score &&
func (t *AvlTree) SearchNodeOrLeft(score int) (me *Node) {
	me = t.Tree

	for {
		if me == nil {
			break
		}
		if me.Score == score {
			break
		}
		if me.Score < score {
			me = me.Right
		} else {
			me = me.Left
		}
	}

	return
}

func (t *AvlTree) Search(score int) []interface{} {
	n := t.SearchNode(score)
	if n != nil {
		return n.Values
	}
	return nil
}

func (t *AvlTree) Reset() {
	tree := t.Tree
	for tree.Parent != nil {
		tree = tree.Parent
	}
	t.Tree = tree
}

func (t *AvlTree) InsertNode(score int, v interface{}) bool {
	if t.Tree == nil {
		t.Tree = &Node{Score: score, Bf: EH, Values: []interface{}{v}}
		return true
	}
	tree := t.Tree
	if score < tree.Score {
		t.Tree = tree.Left
		if !t.InsertNode(score, v) {
			return false
		} else {
			if t.Tree.Parent == nil {
				t.Tree.Parent = tree
			}
			if tree.Left == nil {
				tree.Left = t.Tree
			}

			switch tree.Bf {
			case LH:
				t.LeftBalance(tree)
				return false
			case EH:
				tree.Bf = LH
				t.Tree = tree
				return true
			case RH:
				tree.Bf = EH
				return false
			}
		}
	} else if score > tree.Score {
		t.Tree = tree.Right
		if !t.InsertNode(score, v) {
			return false
		} else {
			if t.Tree.Parent == nil {
				t.Tree.Parent = tree
			}
			if tree.Right == nil {
				tree.Right = t.Tree
			}

			switch tree.Bf {
			case RH:
				t.RightBalance(tree)
				return false
			case EH:
				tree.Bf = RH
				t.Tree = tree
				return true
			case LH:
				tree.Bf = EH
				return false
			}
		}
	} else {
		tree.Values = append(tree.Values, v)
	}
	return true
}

func (t *AvlTree) LeftBalance(tree *Node) {
	subTree := tree.Left
	if subTree != nil {
		switch subTree.Bf {
		case LH:
			// 新插入节点在左子节点的左子树上要做右单旋处理
			tree.Bf = EH
			subTree.Bf = EH
			t.RightRotate(tree)
		case RH:
			// 新插入节点在左子节点的右子树上要做双旋处理
			subTree_r := subTree.Right
			if subTree_r != nil {
				switch subTree_r.Bf {
				case LH:
					tree.Bf = RH
					subTree.Bf = EH
				case RH:
					tree.Bf = EH
					subTree.Bf = LH
				}
				subTree_r.Bf = EH
				t.LeftRotate(subTree)
				t.RightRotate(tree)
			}

		}
	}
}

func (t *AvlTree) RightBalance(tree *Node) {
	subTree := tree.Right
	if subTree != nil {
		switch subTree.Bf {
		case RH:
			//新插入节点在右子节点的右子树上要做左单旋处理
			tree.Bf = EH
			subTree.Bf = EH
			t.LeftRotate(tree)
		case LH:
			//新插入节点在右子节点的左子树上要做双旋处理
			subTree_l := subTree.Left
			if subTree_l != nil {
				switch subTree_l.Bf {
				case LH:
					tree.Bf = EH
					subTree.Bf = RH
				case RH:
					tree.Bf = LH
					subTree.Bf = EH
				}
				subTree_l.Bf = EH
				t.RightRotate(subTree)
				t.LeftRotate(tree)
			}

		}
	}
}

//右单旋
func (t *AvlTree) RightRotate(tree *Node) {
	subTree := tree.Left
	isLeft := false
	if tree.Parent != nil {
		subTree.Parent = tree.Parent //更新新子树的父节点
		if tree.Parent.Left == tree {
			isLeft = true
		}
	} else {
		subTree.Parent = nil
	}
	tree.Left = subTree.Right //原来左节点的右子树挂到老的根节点的左子树
	if subTree.Right != nil {
		subTree.Right.Parent = tree
	}
	tree.Parent = subTree //原来的左节点变成老的根节点的父节点
	subTree.Right = tree  //原来的根节点变成原来左节点的右子树
	tree = subTree
	if tree.Parent == nil { //旋转的是整棵树的根节点
		t.Tree = tree
	} else {
		if isLeft { //更新老的子树根节点父节点指针指向新的根节点
			tree.Parent.Left = tree
		} else {
			tree.Parent.Right = tree
		}
	}
}

//左单旋
func (t *AvlTree) LeftRotate(tree *Node) {
	subTree := tree.Right
	isLeft := false
	if tree.Parent != nil {
		subTree.Parent = tree.Parent
		if tree.Parent.Left == tree {
			isLeft = true
		}
	} else {
		subTree.Parent = nil
	}
	tree.Right = subTree.Left
	if subTree.Left != nil {
		subTree.Left.Parent = tree
	}
	tree.Parent = subTree
	subTree.Left = tree
	tree = subTree
	if tree.Parent == nil {
		t.Tree = tree
	} else {
		if isLeft {
			tree.Parent.Left = tree
		} else {
			tree.Parent.Right = tree
		}
	}
}
