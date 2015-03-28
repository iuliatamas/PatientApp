package main

type ActionQueue []Action

func (pq ActionQueue) Len() int { return len(pq) }

func (pq ActionQueue) Less(i, j int) bool {
	return pq[i].Time().After(pq[j].Time())
}

func (pq ActionQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *ActionQueue) Push(x interface{}) {
	item := x.(Action)
	*pq = append(*pq, item)
}

func (pq *ActionQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
