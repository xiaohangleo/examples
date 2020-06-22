/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2020-06-22 10:22 
# @File : base.go
# @Description : 
# @Attention : 
*/
package tree

import (
	"fmt"
)

type TreeNode struct {
	Data interface{}
	LeftNode *TreeNode
	RightNode *TreeNode
}

// 前序遍历 根左右
func preorderTree(node *TreeNode){
	fmt.Println(node.Data)
	preorderTree(node.LeftNode)
	preorderTree(node.RightNode)
}

// 非递归先序遍历
func preorderTreeWithStack(root *TreeNode)[]interface{}{
	if nil==root {
		return nil
	}
	data:=make([]interface{},0)
	nodes:=make([]*TreeNode,0)
	for nil!=root && len(nodes)!=0{
		// 先序遍历,需要先到最左孩子
		for root!=nil{
			// 保存根节点,既数据信息
			data=append(data,root.Data)
			nodes=append(nodes,root)
			root=root.LeftNode
		}
		// 此时的root为nil,代表的是最左孩子的左侧,意味着到了底部,此时已经保存了左孩子A和A的根节点PA的值
		// 这个节点为根节点,并且已经没有左孩子了,nodes[len(nodes)-2] 的值为该节点的父节点
		node:=nodes[len(nodes)-1]
		//  因为该节点已经数据保存完毕了,该弹出了
		nodes=nodes[:len(nodes)-1]
		// 此时已经保存了左孩子和根节点,所以需要跳到右孩子
		root=node.LeftNode
	}
	return data
}