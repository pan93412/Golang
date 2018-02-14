# Golang 補充
## 前言
- 請先翻閱 [這個網站](http://tour.go-zh.org) 以了解基礎知識。
- 你可以在[這個網站](http://play.golang.org)測試 95% 以上的程式
## 補充區
### Python 的 input() 如何於 Golang 實現？
```
// 第一種方式
reader := bufio.NewReader(os.Stdin)
fmt.Print("Enter text: ")
text, _ := reader.ReadString('\n')
fmt.Println(text)
```

```
// 第二種方式
v1 := ""
fmt.Scanln(&v1)
fmt.Println(v1)
```
Scan : 掃描 (input)

Print: 印出 (output)
### Golang 中 panic() 的使用
Python 的 exit()，在 Go 就是 panic()。

假如你有聽說過 Kernel panic — 就是 Linux 的 BSOD，核心當崩潰時會留下錯誤資訊。

而 panic() 就是這個功能，使程式崩潰並留下錯誤資訊。

panic() 可以單獨使用。 (`panic ("錯誤訊息")`)

panic() 也可以跟 Error() 搭配，使用方法請參閱以下程式碼：
```
package main
import "fmt"

// 建立一個 error
type ErrorMsg struct {
    Title, Msg string
}

func (em *ErrorMsg) Error() string{
    return fmt.Sprint("Error!\nTitle:%v\nMessage:%v", em.Title, em.Msg")
}

func Errorun error{
    return &ErrorMsg{"錯誤標題", "錯誤訊息"}
}

func main{
    panic(Errorun) //強制崩潰，顯示錯誤。
}
```
### 如何讓 Golang 的程式暫停 (pause)？
Python 的 time.sleep() 在 Go 是差不多的用法

`time.Sleep(Second * time.second)`

是的，你必須乘以時間單位。

你可以把 second 換成 minute, hour 看看效果...
### Python 的 random.randint(min, max) 在 Go 的用法
```
package main
import ("fmt"; "math/rand"; "time")
func main() {
    // 建立種子以產生數值
    // 這裡使用現在的 Unix 時間當作種子。
    rand.Seed(time.Now().Unix())
    // 然後產生隨機數目
    // 預設為 0~你輸入的值
    // 會產生 0~128 之間的數目：
    fmt.Println(rand.Intn(128))
    // 但假如我想產生 11~25 的呢
    // 記住這個小公式
    // (max = 欲產生最大值 / min = 欲產生最小值)：
    // rand.Intn(max - min) + min
    // 就可以達到一樣的效果囉：
    fmt.Println(rand.Intn(14) + 11)
}
```
### defer 的用途
```
package main
import ("fmt")
func defer1 {
    // defer 的用意是要在函數結束前 "反轉程序"
    defer fmt.Print("世界")
    fmt.Print("你好")
   // result: 你好世界
}
func main {
    // 因為它 "反轉" 了程序， for 的寫法會很好玩
    for i := 0, i < 10, i++{
        defer fmt.Println(i)
    }
  // 結果是： 10 9 8 7 6 5 4 ...
}
// 更多 defer 的用法將會在下一次的 Gotour 補充。
```
### fmt 的使用
- fmt 輸出文字：
  - fmt.Print()
    - 輸出字串不換行 (\n)
  - fmt.Println()
    - 輸出字串會自動換行 (\n)
  - fmt.Printf()
    - 可格式化字串
    - 可參閱 Golang 對於 fmt.Printf 的說明
- fmt 輸入文字：
  - fmt.Scan()
    - 換行視為空白
  - fmt.Scanln()
    - 換行不視為空白，只輸入第一行內容。
  - fmt.Scanf()
    - 可參閱 Golang 對於 fmt.Scanf 的說明
- fmt 格式化文字：
  - fmt.Sprint()
    - 把字串轉換為 string
  - fmt.Sprintln()
    - 其 string 尾端會換行
  - fmt.Sprintln()
    - 可格式化字串
- [fmt 說明文檔](https://golang.org/pkg/fmt)
## log 的使用
- log 的使用方法跟 fmt.Print 相同。
- log 的函式：
  - log.Print()
  - log.Println()
  - log.Printf()
- 跟 fmt.Print 差別：字串前端有時間
## flag 的使用
- flag 常用函式：
  - flag.[Type]
    - [Type] 包含 int, string, float32, float64, bool 等等
  - flag.[Type]var
  - flag.parse
- 使用教學
```
package main
import "flag"
/**
 * (定義名稱) := flag.int("(名稱)", (預設值 #1), (說明文字))
 * flag.intvar((定義名稱的 Pointer (&)), "(名稱)", (預設值 #1), (說明文字))
 * #1: 假如你寫 flag.int()，那預設值就要用 int 表示。
 *     假如你寫 flag.bool()，那預設值就要用 bool 表示。
 **/
flag1 := flag.bool("autorun", false, "自動執行程式")
var flag2 int
flag.intvar(&flag2, "count", 1, "執行次數")
flag.parse()/** 用戶可以輸入 ./example -a
utorun=true
 * 或者是 ./example --autorun=true
 **/
if *flag1 == True {
    fmt.Println("自動執行程式：開啟")
}
fmt.Printf("執行次數： %v 次\n", *count) // 注意！ flag 得到的值要用 Pointer 讀取。
```
### strings 的使用
#### Replace
- 使用方法： strings.Replace(原文字常數/變數, 欲替換文字, 替換文字, 替換次數)
- 範例：
```
package main
import ("fmt"; "strings")
func main() {
  const originalText string = "byStarTW is byStarTW's friend"
  // 解釋： 替換 originalText 中的 1 個 byStarTW 為 yami_odymel
  newText := strings.Replace(originalText, byStarTW, yami_odymel, 1)
  fmt.Println(newText) // 結果： "yami_odymel is byStarTW's friend"
}
```
### Split
- 使用方法： strings.Split(原文字常數/變數, 切割處)
- 範例：
```
package main
import "strings"
func main() {
  const orText string = "num1 num2 num3"
  // 解釋： 把 orText 中的每一個空白作為一個分割
  result := strings.Split(orText, " ")
  // 結果： ["num1", "num2", "num3"]
}
```
## Authors
作者： byStarTW

歡迎發 PR 補齊更多內容！