package array

import (
	"fmt"
	"math/rand"
)

type RandomizedCollection struct {
	ValueIndexesMap map[int]map[int]struct{}
	Values []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		ValueIndexesMap: make(map[int]map[int]struct{}),
	}
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	_, has := this.ValueIndexesMap[val]
	if !has {
		this.ValueIndexesMap[val] = make(map[int]struct{})
	}
	this.ValueIndexesMap[val][len(this.Values)] = struct{}{}

	this.Values = append(this.Values, val)

	return !has
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	_, has := this.ValueIndexesMap[val]
	if !has {
		return has
	}
	index := 0
	for i := range this.ValueIndexesMap[val] {
		index = i
		break
	}

	delete(this.ValueIndexesMap[val], index)
	n := len(this.ValueIndexesMap[val])
	if n == 0 {
		delete(this.ValueIndexesMap, val)
	}

	n = len(this.Values) - 1
	if index < n {
		this.Values[index] = this.Values[n]
		delete(this.ValueIndexesMap[this.Values[n]], n)
		this.ValueIndexesMap[this.Values[n]][index] = struct{}{}
	}

	this.Values = this.Values[:n]
	return has
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	rIndex := rand.Intn(len(this.Values))
	return this.Values[rIndex]
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func Example() {
	obj := Constructor()
	_ = obj.Insert(0)
	_ = obj.Insert(1)
	_ = obj.Remove(0)
	_ = obj.GetRandom()
	_ = obj.Insert(2)
	_ = obj.Remove(1)
	param_7 := obj.GetRandom()
	fmt.Printf("%d\n", param_7)
}