func hammingWeight(num uint32) int {
    r := 0
    for num != 0 { 
        if num & 1 == 1 {
            r++
        }
        num >>= 1
    }
    return r
}