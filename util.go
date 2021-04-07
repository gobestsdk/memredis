package memredis

//作者：abc-mr
//链接：https://leetcode-cn.com/problems/wildcard-matching/solution/golangdong-tai-gui-hua-jie-jue-tong-pei-bu3cz/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func isMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	res := make([][]bool, sLen+1)
	for i := 0; i <= sLen; i++ {
		res[i] = make([]bool, pLen+1)
	}
	res[0][0] = true //初始条件
	for i := 0; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ { //res[1:sLen+1][0]都为false，不需要考虑j=0情况
			//match表示当前字符s[i-1]和p[j-1]是否相等,i=0表示s为空串，所以i>0
			match := (i > 0 && s[i-1] == p[j-1]) || p[j-1] == '?' || p[j-1] == '*'
			//只要当前两个字符相等，且s[:i-1]与p[:j-1]相匹配res[i][j]就为true,i=0表示s为空串，所以i>0
			res[i][j] = i > 0 && res[i-1][j-1] && match
			if p[j-1] == '*' {
				//两种情况\*做空串或者？
				res[i][j] = res[i][j-1] || (i > 0 && res[i-1][j])
			}
		}
	}
	return res[sLen][pLen]
}
