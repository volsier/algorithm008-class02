func isIsomorphic(s string, t string) bool {
    l := len(s)
    if l == 0 {
        return true
    }
    c := make([]int,l)
    m := make(map[byte]int)
    f := make(map[byte]int)
    num := 0
    num1 := 0
    for i:=0;i<l;i++{
        n,ok := m[s[i]]
        if ok {
            c[i] = n
        }else {
            num++
            c[i],m[s[i]]=num,num
        }
    }
    for i:=0;i<l;i++{
        n,ok := f[t[i]]
        if ok{
            if n != c[i] {
                return false
            }
        }else {
            num1++
            f[t[i]] = num1
            if num1 != c[i]  {
                return false
            }
        }
    }
    return true
}