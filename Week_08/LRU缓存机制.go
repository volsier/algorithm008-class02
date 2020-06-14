type LRUCache struct {
    l int
    m map[int]int
    s []int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{l:capacity,m:make(map[int]int),s:make([]int,0)}
}


func (this *LRUCache) Get(key int) int {
    r,ok := this.m[key]
    if ok {
        index := 0
        if this.s[index] != key {

            for i:=1;i<len(this.s);i++ {
                if this.s[i] == key {
                    index = i
                    break
                }
            }
        }
        this.s = append(this.s[0:index],this.s[index+1:]...)
        this.s = append(this.s,key)
        return r
    }else{
        return -1
    }
}


func (this *LRUCache) Put(key int, value int)  {
    _,ok:=this.m[key]
    if ok {
        index := 0
        if this.s[index] != key {
            for i:=1;i<len(this.s);i++ {
                if this.s[i] == key {
                    index = i
                    break
                }
            }
        }
        this.s = append(this.s[0:index],this.s[index+1:]...)
    }
    if !ok && this.l == len(this.m)  {
        dK := this.s[0]
        delete(this.m,dK)
        this.s = this.s[1:]
    }
    this.m[key] = value
    this.s = append(this.s,key)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */