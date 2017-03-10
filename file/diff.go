package file

import (
	"strings"
	"math"
	"github.com/gin-gonic/gin"
	"contract/response"
	"contract/config"
	"errors"
)

const SepString = "string"       //以字符为单位
const SepLine = "line"           //以一句为单位
const SepParagraph = "paragraph" //以字符为单位
const Add string = "✐"
const Delete string = "✎"
const Replace string = "✏"

var Seps = map[string]string{SepString: "█", SepLine: " ", SepParagraph: " "}
var nowSep string
var cells [][]int16
var preCells [][]int8
var finalCells [][]int16
var lenStr1 int16
var lenStr2 int16
var str1 []string // 横轴 x
var str2 []string // 纵轴 y
var newStr1 []string
var newStr2 []string
var left float64
var above float64
var leftAbove float64
var keywords map[string]int

// [动态编程和基因序列比对](https://www.ibm.com/developerworks/cn/java/j-seqalign/)
func init() {
	keywords, _ = config.GetKeywords()
}

func _init(string1 string, string2 string, sep string) (err error) {
	switch sep {
	case SepString:
		str1 = strings.Split(string1, "")
		str2 = strings.Split(string2, "")
	case SepLine:
		str1 = strings.Split(string1, "。")
		str2 = strings.Split(string2, "。")
		for k, v := range str1 {
			if v == "" {
				str1 = append(str1[:k], str1[k+1:]...)
			}
		}
		for k, v := range str2 {
			if v == "" {
				str2 = append(str2[:k], str2[k+1:]...)
			}
		}
	case SepParagraph:
		str1 = strings.Split(string1, "\n")
		str2 = strings.Split(string2, "\n")
		for k, v := range str1 {
			if strings.TrimSpace(v) == "" {
				str1 = append(str1[:k], str1[k+1:]...)
			}
		}
		for k, v := range str2 {
			if strings.TrimSpace(v) == "" {
				str2 = append(str2[:k], str2[k+1:]...)
			}
		}
	}
	lenStr1 = int16(len(str1))
	lenStr2 = int16(len(str2))

	if lenStr1*lenStr2 == 0 {
		err = errors.New("长度不能为空")
		return
	}
	// 初始化cells
	cells = [][]int16{}
	preCells = [][]int8{}
	var y, x int16
	for x = 0; x < lenStr1+1; x++ {
		cell := make([]int16, lenStr2+1)
		preCell := make([]int8, lenStr2+1)
		cells = append(cells, cell)
		preCells = append(preCells, preCell)
		cells[x][0] = int16(x * (-2))
	}
	for y = 0; y < lenStr2+1; y++ {
		cells[0][y] = int16(y * (-2))
	}
	return
}

func fillIn() {
	var y, x int16
	var leftAboveFormat float64
	for y = 1; y < lenStr2+1; y++ {

		for x = 1; x < lenStr1+1; x++ {
			left = float64(cells[x-1][y])
			above = float64(cells[x][y-1])
			leftAbove = float64(cells[x-1][y-1])

			if str1[x-1] == str2[y-1] {
				leftAboveFormat = leftAbove + 1
			} else {
				leftAboveFormat = leftAbove - 1
			}
			max := math.Max(leftAboveFormat, math.Max(left-2, above-2))
			cells[x][y] = int16(max)
			switch max {
			case left - 2:
				preCells[x][y] = 0
			case above - 2:
				preCells[x][y] = 1
			case leftAboveFormat:
				preCells[x][y] = 2
			}
		}
	}
}

func getTraceBack(sep string) {
	keyX := lenStr1
	keyY := lenStr2
	finalCells = [][]int16{}
	for {
		cell := []int16{keyX, keyY}
		finalCells = append(finalCells, cell)
		if keyX == 0 || keyY == 0 {
			break
		}
		switch preCells[keyX][keyY] {
		case 0:
			keyX = keyX - 1
		case 1:
			keyY = keyY - 1
		case 2:
			keyX = keyX - 1
			keyY = keyY - 1
		}
	}
	newStr1 = []string{}
	newStr2 = []string{}
	var y, x, i, preX, preY int16
	for i = int16(len(finalCells) - 1); i >= 0; i-- {
		x = finalCells[i][0]
		y = finalCells[i][1]
		switch preCells[x][y] {
		case 0:
			preX = x - 1
			preY = y
		case 1:
			preX = x
			preY = y - 1
		case 2:
			preX = x - 1
			preY = y - 1

		}
		if x == 0 || y == 0 {
			continue
		}
		if x-preX == 1 {
			newStr1 = append(newStr1, str1[x-1])
		} else {
			newStr1 = append(newStr1, Seps[sep])
		}
		if y-preY == 1 {
			newStr2 = append(newStr2, str2[y-1])
		} else {
			newStr2 = append(newStr2, Seps[sep])
		}
	}
}

func Comparison(string1 string, string2 string, sep string) (str1 []string, str2 []string, err error) {
	s1 := strings.Split(string1, "\n")
	s2 := strings.Split(string2, "\n")
	for k := range s1 {
		if strings.TrimSpace(s1[k]) == " " {
			s1 = append(s1[:k], s1[k+1:]...)
		}
	}
	for k := range s2 {
		if strings.TrimSpace(s2[k]) == " " {
			s2 = append(s2[:k], s2[k+1:]...)
		}
	}
	string1 = strings.Join(s1, "\n")
	string2 = strings.Join(s2, "\n")
	if sep == SepString {
		switch {
		case string1 == " " && string2 == " ":
			return
		case string1 == " ":
			str2 = strings.Split(string2, "")
			for range str2 {
				str1 = append(str1, Seps[SepString])
			}
			return
		case string2 == " ":
			str1 = strings.Split(string1, "")
			for range str1 {
				str2 = append(str2, Seps[SepString])
			}
			return
		}
	}
	err = _init(string1, string2, sep)
	if err != nil {
		return
	}
	fillIn()
	getTraceBack(sep)
	str1 = newStr1
	str2 = newStr2
	return
}

func ComparisonAll(string1 string, string2 string) (str1 []string, str2 []string, err error) {
	// 比较段落
	nowSep = SepParagraph
	paragraph1, paragraph2, err := Comparison(string1, string2, SepParagraph)
	if err != nil {
		return
	}
	var p1, p2 []string
	if len(paragraph1) > 1 && len(paragraph2) > 1 {
		for k := range paragraph1 {
			if paragraph1[k] != paragraph2[k] {
				if k > 0 {
					if p1[len(p1)-1] == p2[len(p2)-1] {
						p1 = append(p1, paragraph1[k])
						p2 = append(p2, paragraph2[k])
					} else {
						p1[len(p1)-1] = p1[len(p1)-1] + "\n" + paragraph1[k]
						p2[len(p2)-1] = p2[len(p2)-1] + "\n" + paragraph2[k]
					}
				} else {
					p1 = append(p1, paragraph1[k])
					p2 = append(p2, paragraph2[k])
				}
			} else {
				p1 = append(p1, paragraph1[k])
				p2 = append(p2, paragraph2[k])
			}
		}
	} else {
		p1 = append(p1, string1)
		p2 = append(p2, string2)
	}
	nowSep = SepString
	// 	比较字
	for k := range p1 {
		if p1[k] != p2[k] {
			s1, s2, err1 := Comparison(p1[k], p2[k], SepString)
			if err1 != nil {
				err = err1
				return
			}
			p1[k] = strings.Join(s1, "")
			p2[k] = strings.Join(s2, "")
		}
	}

	str1 = p1
	str2 = p2
	return
}

// 增加 Add Delete Replace
func DealComparisonResult(s1 []string, s2 []string) (resStr1 []string, resStr2 []string, err error) {
	if len(s1) != len(s2) || len(s1)*len(s2) == 0 {
		err = errors.New("s1和s2长度要相同且都不能为空")
		return
	}
	for k := range s1 {
		key1 := s1[k]
		key2 := s2[k]
		switch {
		case s1[k] == Seps[SepString]:
			if k > 0 && strings.HasSuffix(resStr1[len(resStr1)-1], Add) {
				resStr1[len(resStr1)-1] = strings.TrimRight(resStr1[len(resStr1)-1], Add)
				resStr1 = append(resStr1, key2+Add)
			} else {
				resStr1 = append(resStr1, Add+key2+Add)
			}
			resStr2 = append(resStr2, s2[k])
		case s2[k] == Seps[SepString]:
			if k > 0 && strings.HasSuffix(resStr1[len(resStr1)-1], Delete) {
				resStr1[len(resStr1)-1] = strings.TrimRight(resStr1[len(resStr1)-1], Delete)
				resStr1 = append(resStr1, key1+Delete)
			} else {
				resStr1 = append(resStr1, Delete+key1+Delete)
			}
		case s1[k] != s2[k]:
			if keywords_v1, ok := keywords[s1[k]]; ok {
				if keywords_v2, ok := keywords[s2[k]]; ok {
					if keywords_v1 == keywords_v2 {
						resStr1 = append(resStr1, s1[k])
						resStr2 = append(resStr2, s2[k])
						continue
					}
				}
			}
			if k > 2 && strings.HasSuffix(resStr1[len(resStr1)-1], Replace) {
				keys := strings.Split(resStr1[len(resStr1)-1], Replace)
				resStr1[len(resStr1)-1] = Replace + keys[1] + key1 + Replace + keys[2] + key2 + Replace
			} else {
				resStr1 = append(resStr1, Replace+key1+Replace+key2+Replace)
			}
			resStr2 = append(resStr2, s2[k])
		default:
			resStr1 = append(resStr1, s1[k])
			resStr2 = append(resStr2, s2[k])
		}
	}
	return
}

func PostDiff(c *gin.Context) {
	var requestData = struct {
		Str1 string    `binding:"required"`
		Str2 string    `binding:"required"`
	}{}
	type responseData struct {
		Str1 []string
		Str2 []string
	}
	if err := c.BindJSON(&requestData); err != nil {
		response.Error("参数格式不正确,原因"+err.Error(), c)
		return
	}
	s1, s2, err := ComparisonAll(requestData.Str1, requestData.Str2)
	if err != nil {
		response.Error(err.Error(), c)
		return
	}
	for k := range s1 {
		resS1, resS2, err := DealComparisonResult(strings.Split(s1[k], ""), strings.Split(s2[k], ""))
		if err != nil {
			response.Error(err.Error(), c)
			return
		}
		s1[k] = strings.Join(resS1, "")
		s2[k] = strings.Join(resS2, "")
	}
	response.Ok(strings.Split(strings.Join(s1, "\n"), ""), c)
	//response.Ok(responseData{s1, s2}, c)
}

func PostDiffOne(c *gin.Context) {
	var requestData = struct {
		Str1 string    `binding:"required"`
		Str2 string    `binding:"required"`
	}{}
	type responseData struct {
		Str1 string
		Str2 string
	}
	if err := c.BindJSON(&requestData); err != nil {
		response.Error("参数格式不正确,原因"+err.Error(), c)
		return
	}
	s1, s2, err := Comparison(requestData.Str1, requestData.Str2, SepString)
	if err != nil {
		response.Error(err.Error(), c)
		return
	}
	resS1, resS2, err := DealComparisonResult(s1, s2)
	if err != nil {
		response.Error(err.Error(), c)
		return
	}
	response.Ok(responseData{strings.Join(resS1, ""), strings.Join(resS2, "")}, c)
}
