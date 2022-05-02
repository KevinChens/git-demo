package lc591

import (
	"strings"
	"unicode"
)

/*
分析：stack+字符串遍历
难题，复杂, 还没懂，暂略
由于标签具有最先开始的标签最后结束的特性，因此我们可以考虑使用一个栈存储当前开放的标签
如果当前的字符为<，那么需要考虑下面的四种情况，这四种情况分不清
如果当前的字符为其它字符，那么根据规则1，栈中需要存在至少一个开放的标签。
在遍历完成后，我们还需要保证此时栈中没有任何（还没有结束的）标签
时间复杂度：O(n), n是字符串code的长度, 只需要对code进行一次遍历
空间复杂度：O(n), 栈存储标签名称需要使用的空间
 */

func isValid(code string) bool {
	tags := make([]string, 0)
	for code != "" {
		if code[0] != '<' {
			if len(tags) == 0 {
				return false
			}
			code = code[1:]
			continue
		}
		if len(code) == 1 {
			return false
		}
		if code[1] == '/' {
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			if len(tags) == 0 || tags[len(tags)-1] != code[2:j] {
				return false
			}
			tags = tags[:len(tags)-1]
			code = code[j+1:]
			if len(tags) == 0 && code != "" {
				return false
			}
		} else if code[1] == '!' {
			if len(tags) == 0 || len(code) < 9 || code[2:9] != "[CDATA[" {
				return false
			}
			j := strings.Index(code, "]]>")
			if j == -1 {
				return false
			}
			code = code[j+1:]
		} else {
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			tagName := code[1:j]
			if tagName == "" || len(tagName) > 9 {
				return false
			}
			for _, ch := range tagName {
				if !unicode.IsUpper(ch) {
					return false
				}
			}
			tags = append(tags, tagName)
			code = code[j+1:]
		}
	}
	return len(tags) == 0
}
