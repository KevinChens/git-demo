package lc344

/*
分析：递归
时间复杂度：O(n)
空间复杂度：O(1)
 */

func ReverseString2(s []byte) {
	res := make([]byte, 0)
	reverse(s, 0, &res)
	// res已经保存了反转的string
	//fmt.Println(res)
	for i := 0; i < len(s); i++ {
		s[i] = res[i]
	}
	//fmt.Printf("%c", s)
}

func reverse(s []byte, i int, res *[]byte) {
	if i == len(s) {
		return
	}
	reverse(s, i+1, res)
	*res = append(*res, s[i])
}

