package main

import "fmt"

// 堆排序的优缺点：就地排序不用辅助数组，最坏情况下o(nlgn)的时间复杂度。
func main()  {
	data:=[]int{10,1,23,24,7,8,9,6,4,288,99,3,34}
	for i:=0;i<len(data)-2;i++{
		//从最后一个父节点开始调整
		for j:=len(data[i:])/2-1;j>=0;j--{ //把这个二叉树变成可处理的二叉树
			heapSort(data[i:],j)
		}
	}
	fmt.Println(data)
}
//堆，即完全二叉树。所有的父节点都比其子节点大
func heapSort(data []int,i int){
	child:=2*i+1
	if 2*i+2<len(data){ //如果存在右孩子并且
		if data[2*i+1]<data[2*i+2]{
			child=2*i+2
		}
	}
	if data[i]>data[child]{
		return
	}else {
		data[i],data[child]=data[child],data[i] //父节点小于子节点换位置
	}
	if child<=(len(data)/2-1){//只要child 的序号还在 //只要当前孩子的索引在所有父节点索引内继续交换
		heapSort(data,child)
	}
}
