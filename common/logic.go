package common

import (
    "iter"
)

func ToSeq[T any](slice []T) iter.Seq[T] {
    return func(yield func(T) bool) {
        for _, v := range slice {
            if !yield(v) {
                return
            }
        }
    }
}

func Map[T, G any](in iter.Seq[T], f func(T) G) iter.Seq[G] {
    return func(yeild func(G) bool) {
        for v := range(in) {
            if !yeild(f(v)) {
                return
            }
        }
    }
}

func Any(in iter.Seq[bool]) bool {
    for v := range(in) {
        if v {
            return true
        }
    }
    return false
}

func All(in iter.Seq[bool])  bool {
    for v := range(in){
        if !v {
            return false
        }
    }
    return true
}
