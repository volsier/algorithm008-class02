func reverseStr(s string, k int) string {
    l := len(s)
    if l <= 1 {
        return s
    }
    b := []byte(s)
    for i:=0;i*k< l;i=i+2 {
        start := i * k
        end := (i+1)*k-1
        if end > l -1 {
            end = l -1
        }
        for start <= end {
            b[start],b[end] = b[end],b[start]
            start++
            end--
        }
    }
    return string(b)
}