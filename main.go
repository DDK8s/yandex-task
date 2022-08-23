package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Graph : represents a Graph G
type Graph struct {
	head  int
	nodes []*GraphNode
	edges int
}

// GraphNode : represents a Graph node - w
type GraphNode struct {
	id          int    // name
	combination string // s
	link        int    // other nodes adjacent to the node
	weight      int    // weight of the edge which starts in this node
}

func main() {
	var letters []string

	graph := New()
	word := input()

	for _, v := range word {
		letters = splitTheWords(v)
		graph.AddNodes(letters)

		fmt.Println(len(graph.nodes))
		fmt.Println(graph.edges)

		graph.printer()
		graph.nodes = nil
	}

}

func (g *Graph) printer() {

	g.nodes[len(g.nodes)-1].link = g.head
	for i := 0; i < len(g.nodes); i++ {
		fmt.Println(g.nodes[i].combination, g.nodes[g.nodes[i].link].combination, g.nodes[i].weight)
	}
}

func (g *Graph) printNodes() {
	fmt.Println(len(g.nodes))
}

func (g *Graph) printEdges() {
	fmt.Println(g.edges)
}

func input() []string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	T := sc.Text()

	t, err := strconv.Atoi(T)
	if err != nil {
		panic(err)
	}
	words := inputWords(t)
	return words
}

func inputWords(t int) []string {
	var words []string
	for i := 0; i < t; i++ {
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		T := sc.Text()
		words = append(words, T)
	}

	return words
}

func splitTheWords(word string) []string {
	a := regexp.MustCompile(``)
	Characters := a.Split(word, -1)
	return Characters
}

/*
func splitTheWords(word string) []string {
	a := regexp.MustCompile(``)
	Characters := a.Split(word, -1)
	return Characters
}*/

// New : returns a new instance of a Graph
func New() *Graph {
	return &Graph{
		nodes: []*GraphNode{},
	}
}

// создание вершины
func (g *Graph) CreateNewNode(newCombination string) (id int) {

	id = len(g.nodes)
	link := id + 1
	if g.nodes != nil {
		for i := range g.nodes {
			if g.nodes[i].combination == newCombination {
				g.nodes[i].weight += 1
				return
			}
		}
	}

	g.edges = g.edges + 1
	g.nodes = append(g.nodes, &GraphNode{
		id:          id,
		combination: newCombination,
		link:        link,
		weight:      1,
	})

	if g.head == 0 {
		g.head = g.nodes[0].id
	}

	return
}

func (g *Graph) AddNodes(words []string) {

	n := len(words) // 30

	for i := 1; i != n-2; i++ {
		newCombination := strings.Join(words[:3], "")
		g.CreateNewNode(newCombination)
		//newCombination = nil
		words = words[1:]
	}

}
