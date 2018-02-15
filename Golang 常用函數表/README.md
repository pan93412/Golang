# Golang 常用函數表
## 函數類
- fmt
  - 輸出文字類
    - `Print()`
    - `Println()`
    - `Printf()`
  - 輸入文字類
    - `Scan()`
    - `Scanln()`
    - `Scanf()`
  - 轉換文字類
    - `Sprint()`
    - `Sprintln()`
    - `Scanf()`
- strings
  - `Replace()`
    - (原文字常數/變數), 欲替換文字, 替換文字, 替換次數
    - 範例程式請參考 Golang 補充資料。
  - `Split()`
    - (原文字常數/變數), 切割處
    - 範例程式請參考 Golang 補充資料。
- log
  - 與 fmt 比較： log 會在文字前端增加時間。
  - `Print()`
  - `Println()`
  - `Printf()`
- os
  - `Exit([返回代號])`
    - 範例： `os.Exit(0)`
  - 資料夾
    - `Chdir()`
    - `Mkdir()`
  - 刪除、更名操作
    - `Remove()`
      僅移除檔案
    - `RemoveAll()`
      移除目錄下所有檔案
    - `Rename()`
  - 檔案資訊
    - `Stat()`
      - 用法： `檔案資訊, err := os.Stat(檔案)`
        - `檔案資訊.IsDir()`
        - 該物件是不是資料夾？
          不是就是檔案，
          可使用 true / false 判斷。
  - 檔案執行： `exec`
    - Command
      - 用法： [程式檔名, 參數]
      - 例如： `os.Exec.Command("dir", "/s")`
    - `Run(程式檔名)`
      - 程式會等待至執行完成。
      - 會回傳 error。(`err := (變數).Run()`)
    - `Start(程式檔名)`
      - 程式**不會**等待至執行完成。
      - 會回傳 error。(`err := (變數).Run()`)
  - math
    - 亂數產生： `Rand`
      - Seed(種子)
        - `math.Rand.Seed(time.Now().Unix())`
      - Intn
        - max = 欲最大值；min = 欲最小值
        - `math.Rand.Intn(max - min) + min`
    - 根號： `Sqrt(數字)`
  - time
    - `Now()`
      - `Unix()` (時間戳)
    - `Sleep()`
      - `Sleep(second * time.Second)`
    - `time.Second`
    - `time.Minute`
    - `time.Hour`
  - flag
    - `flag.[Type]()`
      - [Type] 可以是 int, string, float32, float64, bool 等等
    - `flag.[Type]Var()`
    - `flag.Parse()`
    - 使用方法可參閱 Golang 補充資料。
  - `panic()` (程式結束)
  - func
    - `func (函數名稱) (輸入) (輸出) {}`
  - for
    - `for {}` (無限迴圈)
    - `for i<100 {}` (while)
    - `for i := 0; i<100; i++ {}` (for)
  - if
    - `if (什麼) {(做)} else {(做)}`
    - switch
      - case
      - default
      - 使用範例
      ```
      switch a {
          case 1:
            fmt.Println("a==1")
          case 2:
            fmt.Println("a==2")
          default:
            fmt.Println("a==a")
      }
      ```
      ```
      switch {
          case a < 2:
            fmt.Println("a < 2")
          case a > 2:
            fmt.Println("a > 2")
          default:
            fmt.Println("a == 2")
      }
      ```
- 非函數
  - 類型
    - string
    - int
    - float(32, 64)
    - bool
    - array
      - `(Array 名稱)[(項目數量)](類型){(項目)}`
    - slice
      - `make([](類型), 項目數量, 項目最大長度)`
    - map
      - `make(map[](類型))`
        - {'A': {'b', 'c'}, 'B'}
    - struct
      - `type (名稱) struct {}`
    - interface
    - error
      - 範例程序
      ```
      type theError struct {
          warn string
      }
      func (e theError) Error() string{
          return fmt.Springln("喔喔！程式壞了。原因：", e.string)
      }
      func ProgramErr error{
          return &theError{"沒有原因"}
      }
      func main() {
          panic(ProgramErr)
      }
      ```
  - 格式化字串
    - %v (通用)
    - %s (String)
    - %p (Pointer)

## 相關網頁
  - [Go 官方網頁中文版](http://go-zh.org)
  - [Tour of go 中文](http://tour.go-zh.org)

## Authors
- 製作： @byStarTW
  - 協助： @YamiOdymel, @koru1130
- 歡迎送 PR 來補更多資訊！