package helpers

import (
	"fmt"
	"regexp"
)

func Find() {
	a := "I am learning Go language"

	re, _ := regexp.Compile("[a-z]{2,4}")
	// 找到第一个匹配的
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))

	// 找到所有匹配并保存到切片，n小于0表示返回所有匹配，如果大于0，则表示切片长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll:", all)

	// 查找第一个匹配，开始和结束位置的索引。
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex:", index)

	// 查找所有匹配的索引，n执行与上面相同的工作。
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex:", allindex)

	re2, _ := regexp.Compile("am(.*)lang(.*)")

	// 找到第一个子匹配和返回数组，第一个元素包含所有元素，第二个元素包含first()的结果，第三个元素包含second()的结果
	// 执行后输出如下
	// the first element:am learning Go language
	// the second element: learning Go
	// the third element:uage
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch:", submatch)

	for _, v := range submatch {
		fmt.Println(string(v))
	}

	// 类似于FindIndex()
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println("FindSubmatchIndex:", submatchindex)

	// FindAllSubmatch 找到所有子匹配
	allsubmatch := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println("FindAllSubmatch:", allsubmatch)

	// FindAllSubmatchIndex 找到所有子匹配的所以
	allsubmatchindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println("FindAllSubmatchIndex:", allsubmatchindex)

}
