package common

import (
	"fmt"
	"strconv"
)

type Node[T any] interface {
	GetAdj() map[string]Node[T]
	GetValue() T
	Connect(Node[T]) Node[T]
}

type node[T any] struct {
	value T
	adj   map[string]Node[T]
}

type indexedAdjNode[T any] struct {
	value T
	adj   []Node[T]
}

func CreateNode[T any](value T) Node[T] {
	return &node[T]{value, make(map[string]Node[T], 0)}
}

func (n *node[T]) GetAdj() map[string]Node[T] {
	return n.adj
}

func (n *node[T]) GetValue() T {
	return n.value
}

func (n *node[T]) Connect(c Node[T]) Node[T] {
    str := fmt.Sprintf("%v", c.GetValue())
	n.adj[str] = c
	return n
}

func CreateIndexedAdjNode[T any](value T) Node[T] {
	return &indexedAdjNode[T]{value, make([]Node[T], 0)}
}

func (n *indexedAdjNode[T]) GetAdj() map[string]Node[T] {
    m := make(map[string]Node[T])
    for i, n := range(n.adj) {
        m[strconv.Itoa(i)] = n
    }
	return m
}

func (n *indexedAdjNode[T]) GetValue() T {
	return n.value
}

func (n *indexedAdjNode[T]) Connect(c Node[T]) Node[T] {
    n.adj = append(n.adj, c)
	return n
}

type Graph[T any] interface {
	AddNode(Node[T]) Graph[T]
	GetNodes() []Node[T]
	GetNode(string) Node[T]
	CheckEdge(string, string) bool
}

type graph[T any] struct {
	nodes map[string]Node[T]
}

func CreateGraph[T any]() Graph[T] {
	return &graph[T]{make(map[string]Node[T])}
}

func (g *graph[T]) AddNode(n Node[T]) Graph[T] {
    str := fmt.Sprintf("%v", n.GetValue())
	g.nodes[str] = n
	return g
}

func (g *graph[T]) GetNodes() []Node[T] {
	ns := make([]Node[T], 0)
	for _, n := range g.nodes {
		ns = append(ns, n)
	}
	return ns
}

func (g *graph[T]) GetNode(v string) Node[T] {
	return g.nodes[v]
}

func (g *graph[T]) CheckEdge(v1, v2 string) bool {
	n, ok := g.nodes[v1]
	if !ok {
		return false
	}
	_, ok = n.GetAdj()[v2]
	return ok
}
