package util

import (
	rand2 "math/rand"
	"strconv"
	"strings"
	"time"
)

/**
兑换码随机生成不重复处理
 */
func UniqueExchangeCode(keys map[string]string, rand int) map[string]string {
	max := RandNumMaxCount(rand)
	rand2.Seed(time.Now().UnixNano())
	key := rand2.Intn(max)
	if _, ok := keys[strconv.FormatInt(int64(key), 10)]; ok {
		UniqueExchangeCode(keys, rand)
	} else {
		str := aha_exchange_int2code(key, 8)
		keys[strconv.FormatInt(int64(key), 10)] = strconv.FormatInt(int64(key), 10)
		m := make(map[string]string)
		m["key"] = strconv.FormatInt(int64(key), 10)
		m["value"] = strings.ToUpper(str)
		return m
	}
	return nil
}

func aha_exchange_int2code(number, rand int) string {
	// 定义算法进制
	jz := float64(18)

	// 根据随机位数 计算出在18进制中最大的数是多少
	max := pow(jz, rand)
	// 根据随机位数 计算出在18进制中最小的数是多少
	min, _ := strconv.ParseInt(strconv.FormatFloat(pow(10, rand-1), 'f', -1, 64), int(jz), 32)
	// 根据传入的序号加上初始值
	num := number + int(min)

	// 如果大于最大值，就等于最大值
	if num >= int(max) {
		num = int(max)
	}
	code := strconv.FormatInt(int64(num), 18) // 10 to 18
	// 返回去除重复的字母
	return str_rpl(code)
}

func str_rpl(str string) string {
	rule := make(map[string]string)
	rule["0"] = "k"
	rule["1"] = "l"
	rule["2"] = "m"
	rule["5"] = "n"
	rule["8"] = "p"
	rule["9"] = "r"
	rule["b"] = "t"
	rule["g"] = "w"
	rule["i"] = "x"
	for k, v := range rule {
		str = strings.Replace(str, k, v, -1)
	}
	return str
}

func RandNumMaxCount(rand int) int {
	// 定义算法进制
	jz := float64(18)
	// 根据随机位数 计算出在18进制中最大的数是多少
	max := pow(jz, rand)

	// 根据随机位数 计算出在18进制中最小的数是多少
	min, _ := strconv.ParseInt(strconv.FormatFloat(pow(10, rand-1), 'f', -1, 64), int(jz), 32)
	// 返回数量
	return int(max) - int(min)
}

/*
	实现一个数的整数次方
	pow(x, n)
*/
func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	result := calPow(x, n)
	if n < 0 {
		result = 1 / result
	}
	return result
}

func calPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}

	// 向右移动一位
	result := calPow(x, n>>1)
	result *= result

	// 如果n是奇数
	if n&1 == 1 {
		result *= x
	}

	return result
}
