package main
 
import (
    "fmt"
 
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)
 
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}
 
func init() {
    orm.RegisterDataBase("default", "mysql",
        "root:root@tcp(52.77.219.84:3306)/laravel5?charset=utf8", 30)
    orm.RegisterModel(new(User))
    orm.RunSyncdb("default", true, true)
}
 
func main() {
    orm.Debug = true
    var o = orm.NewOrm()
    // o := orm.NewOrm()
    user := User{Name: "slene"}
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
 
    user1 := User{Name: "tom"}
    id, err3 := o.Insert(&user1)
    fmt.Printf("ID: %d, ERR: %v\n", id, err3)
 
    user2 := User{Name: "jerry"}
    id, err4 := o.Insert(&user2)
    fmt.Printf("ID: %d, ERR: %v\n", id, err4)
 
    user3 := User{Name: "mary"}
    id, err5 := o.Insert(&user3)
    fmt.Printf("ID: %d, ERR: %v\n", id, err5)
 
    user.Name = "astaxie"
    num, err6 := o.Update(&user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err6)
 
    u := User{Id: user.Id}
    err1 := o.Read(&u)
    fmt.Printf("ERR: %v\n", err1)
 
    num, err2 := o.Delete(&u)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err2)
}