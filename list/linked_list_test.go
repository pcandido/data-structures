package list_test

import (
	"testing"
	"data_structures/list"
)

func TestAdd(t *testing.T){
	var list list.LinkedList[int]

	list.AddAfter(1)
	list.AddAfter(2)
	list.AddAfter(3)

	var listStr = list.ToString()
	if(listStr != "[ 1, 2, 3 ]"){
		t.Errorf("Expected: [ 1, 2, 3 ], received: %s", listStr)
	}	
}
