func relativeSortArray(arr1 []int, arr2 []int) []int {
    r := make([]int,0)
    l := len(arr1)
    l2 := len(arr2)
    if l == 0 {
        return r
    }
    if l2 == 0 {
        sort.Ints(arr1)
        return arr1
    }
    m := make(map[int]bool)
    for i:=0;i<l2;i++{
        m[arr2[i]] = true
    }
    t1 := make([]int,0)
    t2 := make([]int,0)
    for i:=0;i<l;i++ {
        if _,ok := m[arr1[i]];ok{
            t1 = append(t1,arr1[i])
        }else{
            t2 = append(t2,arr1[i])
        }
    }
    tmp := make([]int,0)
    for i:=0;i<l2;i++ {
        for j:=0;j<len(t1);j++ {
            if t1[j] == arr2[i] {
                tmp = append(tmp,t1[j])
            }
        }
    }
    sort.Ints(t2)
    return append(tmp,t2...)
}