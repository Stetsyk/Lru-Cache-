package main

import (
	"fmt"

	"testing.com/list"
)

type LRUCache struct {
	mKeyList map[int]*list.List // key list map
	head     *list.List
	capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{mKeyList: make(map[int]*list.List), capacity: capacity}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.mKeyList[key]
	value := -1
	if exist {
		value = node.Value
		this.head = this.head.Erase(node)
		delete(this.mKeyList, key)
		this.head, this.mKeyList[key] = this.head.Append(key, value)
	}
	return value
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.mKeyList[key]
	if exist {
		this.head = this.head.Erase(node)
		delete(this.mKeyList, key)
	}
	if this.capacity == len(this.mKeyList) {
		tmpKey := this.head.Key
		this.head = this.head.Erase(this.head)
		delete(this.mKeyList, tmpKey)
	}
	this.head, this.mKeyList[key] = this.head.Append(key, value)
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 2)
	lru.Put(2, 3)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(1))
	lru.Put(4, 4)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(4))
}
