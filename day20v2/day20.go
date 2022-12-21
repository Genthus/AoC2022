package main

type node struct {
	prev *node
	next *node
	val  int
}

func (n *node) shiftLeftBy(times int) {
	n1 := n.next
	p1 := n.prev
	n1.prev = p1
	p1.next = n1
	for i := -1; i > times; i-- {
		p1 = p1.prev
	}
	p2 := p1.prev
	p1.prev = n
	n.next = p1
	n.prev = p2
	p2.next = n
}

func (n *node) ShiftRightBy(times int) {
	p1 := n.prev
	n1 := n.next
	p1.next = n1
	n1.prev = p1
	for i := 1; i < times; i++ {
		n1 = n1.next
	}
	n2 := n1.next
	n1.next = n
	n.prev = n1
	n.next = n2
	n2.prev = n
}

func (n *node) Shift(times int, length int) {
	times %= length - 1
	if times > 0 {
		n.ShiftRightBy(times)
	} else if times < 0 {
		n.shiftLeftBy(times)
	}
}

func newNode(val int, prev *node, next *node) *node {
	n := node{prev, next, val}
	return &n
}

func createNodeCircle(input []int) map[int]*node {
	nodeMap := make(map[int]*node)
	var prev *node
	for i, v := range input {
		if i == 0 {
			// first node
			prev = newNode(v, nil, nil)
			nodeMap[i] = prev
		} else if i == len(input)-1 {
			// last node
			nodeMap[i] = newNode(v, prev, nodeMap[0])
			prev.next = nodeMap[i]
			nodeMap[0].prev = nodeMap[i]
		} else {
			nodeMap[i] = newNode(v, prev, nil)
			prev.next = nodeMap[i]
			prev = nodeMap[i]
		}
	}
	return nodeMap
}
func createNodeCircleWithKey(input []int, key int) map[int]*node {
	nodeMap := make(map[int]*node)
	var prev *node
	for i, v := range input {
		if i == 0 {
			// first node
			prev = newNode(v*key, nil, nil)
			nodeMap[i] = prev
		} else if i == len(input)-1 {
			// last node
			nodeMap[i] = newNode(v*key, prev, nodeMap[0])
			prev.next = nodeMap[i]
			nodeMap[0].prev = nodeMap[i]
		} else {
			nodeMap[i] = newNode(v*key, prev, nil)
			prev.next = nodeMap[i]
			prev = nodeMap[i]
		}
	}
	return nodeMap
}

func mix(input []int) (map[int]*node, *node) {
	var zeroVal *node
	nodeMap := createNodeCircle(input)
	for i := 0; i < len(input); i++ {
		v := nodeMap[i]
		if v.val == 0 {
			zeroVal = nodeMap[i]
		} else {
			v.Shift(v.val, len(input))
		}
	}
	return nodeMap, zeroVal
}

func reMix(input map[int]*node) *node {
	var zeroVal *node
	for i := 0; i < len(input); i++ {
		v := input[i]
		if v.val == 0 {
			zeroVal = input[i]
		} else {
			v.Shift(v.val, len(input))
		}
	}
	return zeroVal
}

func findValAfter(zeroVal *node, index int, size int) int {
	for i := 0; i < (index % size); i++ {
		zeroVal = zeroVal.next
	}
	return zeroVal.val
}

func printCircle(zeroVal *node) {
	for zeroVal.val != 0 {
		zeroVal = zeroVal.next
	}
	a := false
	for {
		if zeroVal.val == 0 {
			if a {
				break
			}
			a = true
		}
		zeroVal = zeroVal.next
	}
}

func Day22(input []int) (int, int, int) {
	nodeMap, p := mix(input)
	a := findValAfter(p, 1000, len(nodeMap))
	b := findValAfter(p, 2000, len(nodeMap))
	c := findValAfter(p, 3000, len(nodeMap))
	return a, b, c
}

func Day22Part2(input []int, mixTimes int, key int) (int, int, int, *node) {
	nodeMap := createNodeCircleWithKey(input, key)
	var zeroVal *node
	for i := 0; i < mixTimes; i++ {
		zeroVal = reMix(nodeMap)
	}
	a := findValAfter(zeroVal, 1000, len(nodeMap))
	b := findValAfter(zeroVal, 2000, len(nodeMap))
	c := findValAfter(zeroVal, 3000, len(nodeMap))
	return a, b, c, zeroVal
}
