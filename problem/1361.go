package problem

func ValidateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// 统计入度为1的节点数是否为n
	count := 0 // 统计树中节点被访问的次数，有效的二叉树每个节点只能被访问一次，count=n
	// 找到根节点
	// 统计所有节点的入度， 默认入度为0, 节点值做索引，入度作为val
	// n = len(leftChild) = len(rightChild)
	inDegree := make([]int, n)
	for i := range leftChild {
		lNodeVal := leftChild[i]
		rNodeVal := rightChild[i]
		if lNodeVal != -1 {
			inDegree[lNodeVal] += 1
		}
		if rNodeVal != -1 {
			inDegree[rNodeVal] += 1
		}
	}
	// 统计入度为0的节点，有效二叉树的0入度的节点只有一个，即根节点，如果这种节点有多个，说明树不贯通，return flase
	countZero := func(indegree []int) int {
		cnt := 0
		for i := range indegree {
			if indegree[i] == 0 {
				cnt++
			}
		}
		return cnt
	}(inDegree)
	if countZero != 1 { // 根节点不唯一，直接返回false
		return false
	} else {
		// get root node, 一定是入度为0的节点，找到该索引
		root := func(indegree []int) int {
			for i := range indegree {
				if indegree[i] == 0 {
					return i
				}
			}
			return -1
		}(inDegree)
		// 根据层序遍历，统计节点数
		queue := []int{root}
		for len(queue) > 0 {
			node := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			count++
			if count > n {
				break
			}
			if leftChild[node] != -1 {
				queue = append(queue, leftChild[node])
			}
			if rightChild[node] != -1 {
				queue = append(queue, rightChild[node])
			}
		}
		return count == n
	}
}
