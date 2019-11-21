package main

import "bytes"

//填充最后一个分组的函数
// src—》原始数据
// blockLength -> 每个分组的长度
func AddSecret(src []byte,blockLength int)[]byte{
	//1.求出最后一个数组需要多少个字节
	padding := blockLength - len(src)%blockLength
	//2.创建新的切片，并初始化，切片的字节数为padding，创建一个新的字符串，并和src拼接到一起
	padText := bytes.Repeat([]byte{byte(padding)},padding)
	//3.连接最后分组和添加的字符串切片
	newSrc :=append(src,padText...)
	//4.返回新的字符串
	return newSrc
}

//解密时删除添加的字符串

func RemoveSecret(src []byte)[]byte{
	//1.求出要处理的切片长度
	len := len(src)
	//2.取出最后一个字符，得到它的整型值
	number := int(src[len-1])
	//3.将末尾的number个字节删除
	newSrc := src[:len-number]
	return newSrc
}


func main() {
	
}
