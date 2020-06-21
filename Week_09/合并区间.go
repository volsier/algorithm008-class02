func merge(intervals [][]int) [][]int {
    l := len(intervals)
    if l ==0 {
        return nil
    }
    sort.Slice(intervals,func(i,j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    m := [][]int{intervals[0]}
    for i:=1;i<l ;i++ {
        t := m[len(m) -1]
        c := intervals[i]
        if c[0] > t[1] {
            m = append(m,c)
            continue
        }
        if c[1] >  t[1] {
            t[1] = c[1]
        }
    }
    return m
}