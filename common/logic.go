package common

func Map[T, G any](itter func() (T, bool), f func(T) G) func() (G, bool) {
    return func() (G, bool) {
        var zero G
        in, ok := itter()
        if !ok {
            return zero, false
        } else {
            return f(in), true
        }
    }
}

func Any(itter func() (bool, bool)) bool {
    for {
        val, ok := itter()
        if !ok {
            break
        } else if val {
            return true
        }
    }
    return false
}

func All(itter func() (bool, bool))  bool {
    for {
        val, ok := itter()
        if !ok {
            break
        } else if !val {
            return false
        }
    }
    return true
}

