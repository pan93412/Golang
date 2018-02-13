package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

/**
 * 名稱：StarCleaner
 * 作者：byStarTW
 * 發布於 Golang 範例程式。
 **/
var f1, f1mode string

/** 錯誤處理 **/
type theErrors struct {
	filename string
}

func (e theErrors) Error() string {
	return fmt.Sprint(e.filename, " 不存在。") // 輸出 [檔案名稱] 不存在。
}

func filenotexist() error {
	return &theErrors{f1}
}

/** CustomClean 資料夾建立 **/
func check() {
	info, err := os.Stat("CustomClean")
	switch {
	case err != nil:
		os.Mkdir("CustomClean", 0777)
	case info.IsDir() == false:
		os.Remove("CustomClean")
		os.Mkdir("CustomClean", 0777)
	}
}

/** 主程序 **/
func main() {
	/** 指令列選項設定 **/
	flag.StringVar(&f1, "custom", "", "清理除了 CustomClean 的其他檔案、資料夾。")
	flag.Parse()
	/** CustomClean 建立 **/
	check()
	/** 說明文字 **/
	fmt.Println("StarCleaner 自訂垃圾清理器")
	fmt.Println("請把垃圾投進去 CustomClean 以移除這些廢棄檔案。")
	fmt.Println("更多這個程式的用法，可以參閱 'StarCleaner --help'。")
	fmt.Println("\n將會清理的清單：")
	fmt.Println("* CustomClean 資料夾")
	if f1 != "" {
		info, err := os.Stat(f1)
		if err != nil {
			panic(filenotexist()) // 檔案不存在時，程式終止。
		} else {
			if info.IsDir() == true {
				f1mode = "資料夾"
			} else {
				f1mode = "檔案"
			}
		}
		fmt.Println("* 您自訂的", f1mode, "：", f1) // 自動判斷檔案、資料夾。
	}
	fmt.Print("輸入 C 清除，E 離開。")
	/** 用戶選擇 **/
	var UserChoice string
	fmt.Scanln(&UserChoice)
	/** 清理程序 **/
	if UserChoice == "C" {
		/** 清理 CustomClean 資料夾 **/
		log.Println("\n清除 CustomClean 資料夾 ...")
		os.RemoveAll("CustomClean")
		log.Println("清除完成。")
		/** 清理用戶自訂的資料夾。 **/
		if f1 != "" {
			// 只會在用戶使用 --custom="" 才會使用到這個部分。
			log.Println("清除", f1, f1mode, "...")
			if f1mode == "資料夾" {
				os.RemoveAll(f1)
			} else {
				os.Remove(f1)
			}
			log.Println("清除完成。")
		}
		/** 程序結束 **/
		log.Println("程序結束。")
	}
}
