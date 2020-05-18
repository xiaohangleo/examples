/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2019-12-16 16:50 
# @File : wallet.go
# @Description : 
# @Attention : 
*/
package models

import (
	"myLibrary/go-library/go/wallet"
	wallet2 "vlink.com/v2/vlink-common/constants/usercc/wallet"
)

// 用户钱包容器

// 用户主钱包
type VlinkWallet struct {
	*wallet.HDWallet

	WalletTree WalletTree

	ChildWallets []*WalletTreeNode
}

func (v *VlinkWallet) CopyFromTheOther(w *VlinkWallet) {
	if w!=nil{
		v.WalletTree = w.WalletTree
		v.ChildWallets = w.ChildWallets
	}
}

func (v *VlinkWallet) GetByType(Type int) (*WalletTreeNode, bool) {
	for _, c := range v.ChildWallets {
		if c.Type == Type {
			return c, true
		}
	}
	return nil, false
}
func (v *VlinkWallet) NewVlinkChildWallet(Type int) (*VlinkChildWallet, error) {
	c := new(VlinkChildWallet)
	cc, e := v.NewChildWallet(Type)
	if nil != e {
		return nil, e
	}
	c.ChildWallet = cc

	// if nil == v.WalletTree {
	// 	v.WalletTree = *NewWalletTree()
	// }
	v.ChildWallets = append(v.ChildWallets, &WalletTreeNode{
		Type: Type,
		Path: c.Path,
		// WalletDetail: c,
	})
	// v.WalletTree.Insert(Type, &WalletTreeNode{
	// 	Type: Type,
	// 	Path: c.Path,
	// })

	return c, nil
}

func NewVlinkWallet(inOrOut int, pwd string) *VlinkWallet {
	v := new(VlinkWallet)
	v.HDWallet = wallet.NewHDWallet(wallet2.WALLET_TYPE_MAIN_WALLET, inOrOut, pwd)
	return v
}

type VlinkChildWallet struct {
	*wallet.ChildWallet
}
