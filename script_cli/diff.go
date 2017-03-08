package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"contract/file"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	s1, _ := ioutil.ReadFile("a.txt")
	s2, _ := ioutil.ReadFile("b.txt")
	string1 := string(s1)
	string2 := string(s2)
	//string1 = "GCCCTAGCG"
	//string2 = "GCGCAATG"
	//fmt.Println(string1, string2)
	time1 := time.Now().UnixNano()
	m1, _ := mem.VirtualMemory()
	str1, str2, err := file.Comparison(string1, string2, file.SepString)
	m2, _ := mem.VirtualMemory()
	fmt.Println(str1)
	fmt.Println(str2)
	if err != nil {
		panic(err.Error())
	}
	for k := range str1 {
		fmt.Println(k, str1[k])
	}
	fmt.Println("------------------------------------")
	for k := range str2 {
		fmt.Println(k, str2[k])
	}
	time2 := time.Now().UnixNano()
	fmt.Println("====================================")
	fmt.Println(time2-time1, "纳秒")
	fmt.Println((time2-time1)/1000, "微妙")
	fmt.Println((time2-time1)/1000000, "毫秒")
	fmt.Println((time2-time1)/1000000000, "秒")
	fmt.Println("内存占用", m2.Used-m1.Used, "byte")
	fmt.Println("内存占用", (m2.Used-m1.Used)/1000, "k")
	fmt.Println("内存占用", (m2.Used-m1.Used)/1000000, "M")
	fmt.Println("内存占用", (m2.Used-m1.Used)/1000000000, "G")
}
