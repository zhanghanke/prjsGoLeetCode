/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.



先来一波brute force，结果被吊打~~~
982 / 983 test cases passed.
Status: Time Limit Exceeded
Submitted: 0 minutes ago
Last executed input:
*/

func lengthOfLongestSubstring(s string) int {
    end := len(s)
	maxlen := 1
    if (len(s) == 0){
        return 0
    }
	for end > 0 {
		s = s[:end]
		start := 0
		for start < end-1 {
			ss := s[start:end]
			for ii, cc := range ss {
				ii++
				for _, ccc := range ss[ii:] {
					if cc == ccc {
						goto BK_LABEL
					}
				}
			}
			if (end - start) > maxlen {
				maxlen = end - start
			}
		BK_LABEL:
			start++
		}
		end--
	}
	return maxlen
}
/*
做了两点改进：1，并不需要全部组合相邻字符串，只需要用两个指针：前、后指针，如果两个指针中间的内容没有重复，则移动前指针，如果有重复的，则移动后指针
    2.前、后指针的内容不用数组，而是用一个set(golang没找到set容器，用map代替一下).
    
    虽然我改进思路跟discuss与solution里面是一样的，但在实现的时候还是经常多N多错误与调试。再次感叹自己编码水平实在太弱鸡了。
    
*/
func lengthOfLongestSubstring(s string) int {
	var maxlen, start, current, length int = 0, 0, 0, len(s)
	if length == 0 {
		return 0
	}
	var set map[int]int = make(map[int]int)
	for start < length && current < length {
		v := int(s[current])
		_, ok := set[v]
		if !ok {
			set[v] = v
			current++
			if (current - start) > maxlen {
				maxlen = current - start
			}

		} else {
			delete(set, int(s[start]))
			start++

		}
	}
	return maxlen
}

/*
再次做了改进：当前指针发现前面重复的字符x，则后指针与x之间的指针是不需要一步步去比较的，可以一步就移动到x后面一个.  
同时发现前面len == 0是多余的判断
*/
func lengthOfLongestSubstring2(s string) int {
	var maxlen, start, current, length int = 0, 0, 0, len(s)
	var set map[int]int = make(map[int]int)
	for current < length {
		ch := int(s[current])
		v, ok := set[ch]
		if ok {
			if v > start {
				start = v
			}
		}
		if (current - start + 1) > maxlen {
			maxlen = current - start + 1
		}
		set[ch] = current + 1
		current++
	}
	return maxlen
}
