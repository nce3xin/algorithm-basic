package main

// 给定字符串（包含若干单词），单词之间只有一个空格，输出每个单词
// 这是最简单的双指针应用
// 模板就是为了节省大家思考的时间
func printWords(s string) []string {
	res := make([]string, 0)
	n := len(s)
	for i := 0; i < n; i++ {
		j := i
		for j < n && s[j] != ' ' {
			j++
		}
		// 此时j指向空格
		// 每道题目的具体逻辑
		res = append(res, s[i:j])
		i = j
	}
	return res
}
