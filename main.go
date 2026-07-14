package main
import ("fmt";"sort";"strings")
var sortID = "stream-reader-d5220b"
type Item struct{Name string;Score int}
type ByScore []Item
func (a ByScore) Len() int{return len(a)}
func (a ByScore) Swap(i,j int){a[i],a[j]=a[j],a[i]}
func (a ByScore) Less(i,j int) bool{return a[i].Score>a[j].Score}
func main(){items:=[]Item{{"alpha",42},{"beta",99},{"gamma",17},{"delta",73},{"epsilon",55}};sort.Sort(ByScore(items));fmt.Printf("[%s] Sorted:\n",sortID);for i,it:=range items{fmt.Printf("  %d. %s (%d)\n",i+1,it.Name,it.Score)};_=strings.Repeat("",0)}
