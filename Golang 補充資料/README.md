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
#### Split
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
### runtime 的使用
- `runtime.GOOS`: 判斷系統。
- `runtime.GOARCH`: 判斷系統架構。
### for 的兩個小函式
- `continue`：跳過這一次迴圈
  - 範例
  ```
  for i := 0; i < 4; i++ {
    fmt.Println(i)
    if i == 2 { continue }
  }
  ```
  - 結果： 1, 3
- `break`：停止迴圈
  - 範例
  ```
  for i := 0; i < 4; i++ {
    fmt.Println(i)
    if i == 2 { break }
  }
  ```
  - 結果： 1
### os 的使用
#### 說明
```
/**
 * 這個使用教學會分三篇，分別是 「資料夾操作」、
 * 「檔案操作」 跟 「檔案資訊」。
 * 在看這篇文章之前，你可以看看前幾篇的 os.Stdout、
 * os.Stdin、os.Stderr 分別代表什麼意思。
 * 而在 「檔案操作」 這篇，我也會講關於
 * os.O_CREATE, os.O_RDONLY 等等常用的檔案操作函數。
 **/
```
#### os 資料夾操作
```
/**
 * 教學： os 使用 - 資料夾操作

 * 編撰者： byStarTW
 * 如果代碼有任何錯誤，歡迎發 issues 或 PR！
 **/
package main
import ("fmt"; "os")
/**
 * 資料夾操作的常用指令：
 *  - os.Mkdir(資料夾名稱, UNIX 權限) [error]
 *    稍微講一下 Unix 權限的寫法：
 *    4- 讀取 / 2- 寫入 / 1- 執行
 *    把你想要的權限數字加起來
 *    假如我要讀取跟執行權限，那就是 5。寫入跟執行權限就是 3、讀取跟寫入權限就是 6、
 *    所有權限就是 7。
 *    然後 OOO 依序是擁有者；群組；所有使用者 （有點不太確定?）
 *    如果是 Windows 系統你可以直接寫 0777
 *  - os.RemoveAll (資料夾名稱)
 *    os.Remove 用來移除檔案、os.RemoveAll 用來移除資料夾 （與其底下的所有檔案、資料夾）
 *  - os.Stat(資料夾名稱) [FileInfo, error]
 *    - Name() 資料夾名稱
 *    - IsDir() 是否為資料夾 (回傳 true / false)
 **/
func main() {
  os.Mkdir("My1stFolderWithGo") // 建立 My1stFolderWithGo 資料夾
  Hi, _ := os.Stat("My1stFolderWithGo")
  if Hi.Name() != "My1stFolderWithGo" {
    // 假如這個資料夾的名稱不是 My1stFolderWithGo (?)
    fmt.Println("哪裡出錯了... 他的名稱不是", Hi.Name())
  }
  if Hi.IsDir() != true {
    // 假如這個資料夾不是個資料夾
    fmt.Println("哪裡出錯了... 他不是個資料夾")
  }
  fmt.Println("資料夾名稱：", Hi.Name())
}
```
#### os 檔案操作
```
/**
 * 教學： os 使用 - 檔案操作
 * 編撰者： byStarTW
 * 如果代碼有任何錯誤，歡迎發 issues 或 PR！
 **/
package main
import ("fmt"; "os")
/**
 * os 關於檔案操作的函式，我個人覺得已經足夠詳細。
 * 那麼，開始囉！ :D
 * - os.Open("檔案名稱") [檔案操作函式, err] 開啟檔案
 * - os.Create("檔案名稱") [檔案操作函式, err] 新建檔案
 * - os.OpenFile("檔案名稱", flag, UNIX 權限) [檔案操作函式, err] 以上兩個函式的 Easy 版
 *   - 常用 flag
 *     - os.O_CREATE 若檔案不存在，建立一個新資料夾
 *     - os.O_TRUNC 如果檔案存在，清空檔案
 *     - os.O_RDWR 賜予讀取、寫入權限
 *     - os.O_RDONLY 僅賜予讀取權限
 *     - os.O_WRONLY 僅賜予寫入權限
 *   - 檔案操作函式
 *     - Close() 關閉檔案
 *     - Sync() 變更寫入檔案
 *     - Stat() 用法同 os.Stat()
 *     - 另外這個函式也可以直接用來讀取跟寫入。
 *       (io.Reader / io.Writer)
 * - os.Remove("檔案名稱") 移除檔案
 * - os.Stat("檔案名稱") [StatInfo, err]
 *   - Name() 檔案名稱
 *   - Size() 檔案大小 (bytes)
 **/
func main() {
  // 首先，我們先建立一個檔案
  // [!] 假如你想要執行這個程式碼，請確保你的電腦
  // 是否有 "My1stFileWithGo.txt"
  TheFile, _ := os.OpenFile("My1stFileWithGo.txt", os.O_CREATE|os.O_RDWR|os.TRUNC, 777)
  TheFile.Close() // 關閉檔案
  Hi, _ := os.Stat("My1stFileWithGo.txt")
  fmt.Println("檔案名稱：", Hi.Name(), "\n檔案大小：", Hi.Size(), " bytes")
  os.Remove("My1stFileWithGo.txt")
}
```
#### os 補充
```
/**
 * 教學： os 使用 - 補充
 * 編撰者： byStarTW
 * 如果代碼有任何錯誤，歡迎發 issues 或 PR！
 **/
package main
import ("fmt"; "os"; "os/exec")
/**
 * os 的補充
 * 這一篇會講關於 Exec 執行檔案、
 * os.Exit() 關閉程序等函式。
 *
 * - Exec 執行檔案 （使用前 import "os/exec"）
 *   - exec.Command("程式名稱", "flag 1", "flag 2", ...) [Command 參數]
 *     - Run() 開啟並等待程式執行完成
 *     - Start() 開啟但不等待程式執行完成
 *     - Command 參數.Stdout = os.Stdout 把程式輸出轉移到標準輸出 (os.Stdout)
 * - os.Exit(程式離開代碼 int)
 *   - 通常 os.Exit(0) 是正常結束，os.Exit(-1) 是程式錯誤
 *   - 假如這個錯誤嚴重，你可以把 os.Exit() 換成 panic()
 **/
func main() {
  // 開啟 grep，不等待 grep 執行完成
  p := exec.Command("grep") // 設置 grep 指令
  p.Start() // 執行
  // 清除畫面 (for Linux)
  p = exec.Command("clear") // 設置 clear 指令
  p.Stdout = os.Stdout // 把 clear 的輸出輸出至 os.Stdout 標準輸出
  p.Run() // 執行
  // 關閉程式
  os.Exit(0) // 正常關閉
}
```
#### 輸入、輸出跟錯誤
- `os.Stdout`：輸出
- `os.Stdin`：輸入
- `os.Stderr`：錯誤
### bufio 的教學
Reader 部分：
```
/**
 * 教學： bufio [READER]
 * 適用對象： Golang 基礎者
 **/
package main
import ("bufio"; "os"; "fmt")

/**
 * 關於 Reader 類別：
 * 讀取一個來源資料、然後輸出一個資料。
 * Reader 常用的類別有：
 * bufio.NewReader(Source) [Reader] 讀取來源資料指令
 * bufio.NewReaderSize(bufio.NewReader(Source), Size) 設定來源資料的緩衝區大小
 * Reader.ReadString(charset) 以一個 charset (e.g '\n'、'\t')
 * 分割一個字串，假如寫 Reader.ReadString('\n')，遇到 \n 就會
 * 自己截斷。
 * 若截斷的文字後面還有文字，可以再使用一次這個指令，
 * 把後面的內容讀取出來。
 * Reader.Buffered() 目前緩衝區剩餘的內容數量 (int)
 **/
func main {
  /** 第一種使用情況 **/
  MyReader := bufio.NewReader(os.Stdin) // 讀取來源 os.Stdin 的內容
  fmt.Println("輸入些什麼：")
  // 注意！ ReadString 會輸出 [string, error]，若只需要
  // 文字請自行指定。
  mrStr, mrErr := MyReader.ReadString('\n') // 讀取 os.Stdin 直到那個 charset 的內容。
  fmt.Printf("您輸出了 %v，錯誤為 %v。\n", mrStr, mrErr)
  /** 第二種使用情況 **/
  myFile := os.OpenFile("theText.txt", os.O_CREATE|os.RDWR, 666) // 開啟 theText.txt
  MyReader = bufio.NewReader(myFile) // 讀取 theText.txt 的檔案
  // 反覆的一直讀取資料，直到所有檔案讀取完畢為止。
  for MyReader.Buffered() == 0 {
    fmt.Println(MyReader.ReadString('\n'))
  }  
}
```
Writer 部分：
```
/**
 * 教學： bufio [WRITER]
 * 適用對象： Golang 基礎者
 **/
package main
import ("bufio"; "os"; "fmt")

/**
 * 關於 Writer 類別：
 * 建立一個欲寫入位置、寫入資料、並且把資料寫入原位置。
 * Writer 常用的類別有：
 * bufio.NewWriter(Source) [Writer] 設定寫入位置
 * bufio.NewWriterSize(bufio.NewWriter(Source), Size) 設定緩衝區大小
 * Writer.WriteString("字串") 寫入資料至寫入位置。
 * Writer.Flush() 把緩衝區內容寫入 「寫入位置」
 **/
func main {
  myFile := os.OpenFile("theText.txt", os.O_CREATE|os.RDWR, 666) // 開啟 theText.txt
  MyWriter = bufio.NewWriter(myFile) // 指定 myFile 為寫入位置
  MyWriter.WriteString("Hello! I am byStarTW") // 寫入字串 "Hello! I am byStarTW"
  MyWriter.Flush() // 寫入
  myFile.Sync() // 把 myFile 目前存在記憶體的檔案寫入實體硬碟內。
  myFile.Close() // 關閉檔案
}
```
## Authors
作者： byStarTW

歡迎發 PR 補齊更多內容！