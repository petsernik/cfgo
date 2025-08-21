package main

import "fmt"

func main() {
	d := &dsu{}
	d.Init(5)
	d.UnionSets(3, 4)
	for i := range 6 {
		fmt.Printf("%d in one group with %d\n", i, d.FindSet(i))
	}
	fmt.Println()

	dm := &dsuMap[string]{}

	dm.Init()
	creatures := []string{"cat", "dog", "man", "woman"}
	for _, creature := range creatures {
		dm.MakeSet(creature)
	}
	dm.UnionSets("cat", "dog")
	dm.UnionSets("man", "woman")
	for _, creature := range creatures {
		fmt.Printf("%s in one group with %s\n", creature, dm.FindSet(creature))
	}
}
