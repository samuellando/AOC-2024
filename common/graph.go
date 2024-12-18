package common

import "strconv"

type Node interface {
	GetAdj() map[string]Node
	GetValue() string
	Connect(Node) Node
}

type node struct {
	value string
	adj   map[string]Node
}

type indexedAdjNode struct {
	value string
	adj   []Node
}

func CreateNode(value string) Node {
	return &node{value, make(map[string]Node, 0)}
}

func (n *node) GetAdj() map[string]Node {
	return n.adj
}

func (n *node) GetValue() string {
	return n.value
}

func (n *node) Connect(c Node) Node {
	n.adj[c.GetValue()] = c
	return n
}

func CreateIndexedAdjNode(value string) Node {
	return &indexedAdjNode{value, make([]Node, 0)}
}

func (n *indexedAdjNode) GetAdj() map[string]Node {
    m := make(map[string]Node)
    for i, n := range(n.adj) {
        m[strconv.Itoa(i)] = n
    }
	return m
}

func (n *indexedAdjNode) GetValue() string {
	return n.value
}

func (n *indexedAdjNode) Connect(c Node) Node {
    n.adj = append(n.adj, c)
	return n
}

type Graph interface {
	AddNode(Node) Graph
	GetNodes() []Node
	GetNode(string) Node
	CheckEdge(string, string) bool
}

type graph struct {
	nodes map[string]Node
}

func CreateGraph() Graph {
	return &graph{make(map[string]Node)}
}

func (g *graph) AddNode(n Node) Graph {
	g.nodes[n.GetValue()] = n
	return g
}

func (g *graph) GetNodes() []Node {
	ns := make([]Node, 0)
	for _, n := range g.nodes {
		ns = append(ns, n)
	}
	return ns
}

func (g *graph) GetNode(v string) Node {
	return g.nodes[v]
}

func (g *graph) CheckEdge(v1, v2 string) bool {
	n, ok := g.nodes[v1]
	if !ok {
		return false
	}
	_, ok = n.GetAdj()[v2]
	return ok
}
