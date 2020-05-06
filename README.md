# Golang_testing
## 準備測試
首先先建立一個你想測試的檔案名稱"名稱_test.go"的檔案.
測試func的參數都要是*testing.T.
方法名稱也要是Test開頭.

## Table-Driven Test
使用表驅動測試, 就是給一組列表內有輸入和預期的結果;
都去執行相同的待測單元, 然後比對預期的結果.

*參考FizzBuzz_test.go內的TestFizzBuzzSuccess方法

go test  可以執行所有測試範例
go test -run "想要測試的方法" 可以執行指定的測試範例
    ex.go test -run TestFizzBuzzSuccess
go test -run -timeout 時間  可以加上超時時間
    ex.go test -run TestFizzBuzzSuccess -timeout 1ns 

## Benchmark Test

這裡測試方法名稱要以Benchmark開頭, 傳入參數為(b *testing.B)
b.N表示的是循環的次數, 因為是壓測, 所有要反覆的測試待測目標
測試時間預設是1秒鐘, 會顯示測試次數跟每次所花費的時間成本.

*參考FizzBuzz_test.go內的Benchmark_FizzBuzz方法

運行方法 
`
go test -bench=.
`
運行結果
`
goos: windows
goarch: amd64
Benchmark_FizzBuzz-4    33841203                32.1 ns/op
PASS
ok      _/C_/Users/user/Desktop/Golang_testing  3.378s
`
總共測試33841203次, 每次大約花費32.1 ns.


還能自己定義測試時間
加上-benchtime
`
go test -bench=. -benchtime=2s
`

記憶體佔用測試
加上-benchmem
`
go test -bench=.  -benchmem
`
結果
`
goos: windows
goarch: amd64
Benchmark_FizzBuzz-4    33331388                36.4 ns/op             4 B/op
PASS
ok      _/C_/Users/user/Desktop/Golang_testing  1.382s
`

4B/op 表示每次呼叫呼叫配置4Bytes
0 allocs/op 表示每次呼叫需要進行幾次分配對象記憶體.
這裡是0 是因為我們的待測目標沒有去new或是宣告任何參考型別.