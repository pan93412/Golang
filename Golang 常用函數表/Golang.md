# Golang 常用函數表
## 函數類
- fmt
  - 輸出文字類
    - Print
    - Println
    - Printf
  - 輸入文字類
    - Scan
    - Scanln
    - Scanf
  - 轉換文字類
    - Sprint
    - Sprintln
    - Scanf
- os
  - 資料夾
    - Chdir
    - Mkdir
  - 刪除、更名操作
    - Remove
      僅移除檔案
    - RemoveAll
      移除目錄下所有檔案
    - Rename
  - 檔案資訊
    - Stat
      - 用法： Stat [檔案資訊, err]
        - 檔案資訊.IsDir()
        - 該物件是不是資料夾？
          不是就是檔案，
          可使用 true / false 判斷。
  - 檔案執行： exec
    - Command
      - 用法： [程式檔名, 參數]
      - 例如： `Command("dir", "/s")`
    - Run
      - 程式會等待至執行完成。
    - Start
      - 程式**不會**等待至執行完成。
  - math
    - 亂數產生： Rand
      - Seed(種子)
        - `Seed(time.Now().Unix())`
      - Intn
        - max = 欲最大值；min = 欲最小值
        - `Intn(max - min) + min`
    - 根號： Sqrt
  - time
    - Now()
      - Unix() (時間戳)
    - Sleep()
      - `Sleep(second * time.Second)`
    - time.Second
    - time.Minute
    - time.Hour
  - panic() (程式結束)
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
    - 
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
          return &theError{'沒有原因'}
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