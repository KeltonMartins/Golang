package main
import "fmt"

func formatar (slice []int, k int, sep string) string{
    stringvolta := "[ "
    for _, v:= range slice{
        if v == k{
            stringvolta += fmt.Sprint(v) + "> "
        }else{
            stringvolta += fmt.Sprint(v) + sep
        }
    }
    stringvolta += "]"
    return stringvolta
}

func main() {
    var n, matador int
    fmt.Scanf("%d %d", &n, &matador)

    var matar []int
    for i:=1; i<=n; i++{
        matar = append(matar, i)
    }
    k := matador - 1
    fmt.Println(formatar(matar, matar[k], " "))

    for len(matar) > 1{
        indexDaMorte := (k+1) % len(matar)
        matar = append(matar[:indexDaMorte], matar[indexDaMorte+1:]...)
        k = indexDaMorte % len(matar)
        fmt.Println(formatar(matar, matar[k], " "))
    }
}
