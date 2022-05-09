package lc933

/*
分析：rcueue
我们可以用一个队列维护发生请求的时间，当在时间t收到请求时，将时间t入队
由于每次收到的请求的时间都比之前的大，因此从队首到队尾的时间值是单调递增的
当在时间t收到请求时，为了求出[t−3000,t]内发生的请求数，
我们可以不断从队首弹出早于t−3000的时间。
循环结束后队列的长度就是[t−3000,t]内发生的请求数。
时间复杂度：O(1), 每个元素至多入队出队各一次
空间复杂度：O(L), 其中L为队列的最大元素个数
 */

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (rc *RecentCounter) Ping(t int) int {
	for len(rc.queue) > 0 && rc.queue[0] < t - 3000 {
		rc.queue = rc.queue[1:]
	}
	rc.queue = append(rc.queue, t)
	return len(rc.queue)
}
