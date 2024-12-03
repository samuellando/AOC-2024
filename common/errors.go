package common

func net[T any](v T, err error) T {
    if err != nil {
        panic(err)
    }
    return v
}
