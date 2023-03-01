# Marinesnow

## 概要
シンプルなGolang製 snowflake ID 採番ライブラリです。
`workerID` と `timestampOffset` を設定(どちらも任意)すれば、あとは `Generate()` 関数を叩くだけで snowflake ID を採番することが出来ます。

## 具体的な使用方法

コード例:
``` golang
package main

import (
	"fmt"
	"marinesnow"           // ID 採番機本体.
	"marinesnow/snowflake" // snowflake 型を提供するサブパッケージ.
)

func main() {
	// サーバーの ID など、マシン固有の ID を指定(任意).
	marinesnow.SetWorkerID(5)

	// 採番開始時刻をUnixtime(ミリ秒精度)で指定(任意).
	marinesnow.SetTimestampOffset(1677698659356)

	// snowflake を生成する.
	if sf, err := marinesnow.Generate(); err == nil {

		// snowflake を出力.
		fmt.Println("snowflake:", sf)

		// snowflake からタイムスタンプを取得.
		fmt.Println("timestamp:", sf.Timestamp())

		// snowflake から workerID を取得.
		fmt.Println("workerID:", sf.WorkerID())

		// snowflake からインクリメント値を取得.
		fmt.Println("increments:", sf.Increments())

		// 底を指定して文字列に変換することも可能.
		fmt.Println("base2: ", sf.ToString(2))
		fmt.Println("base8: ", sf.ToString(8))
		fmt.Println("base10:", sf.ToString(10))
		fmt.Println("base16:", sf.ToString(16))
		fmt.Println("base32:", sf.ToString(32))
	}

	// 既存の snowflake も使用できます.
	existing := snowflake.Snowflake(1490753660736864259)
	fmt.Println("existing:", existing)
}
```

出力例:
```
snowflake: 8298866692096
timestamp: 1978604
workerID: 5
increments: 0
base2:  1111000110000111011000000000101000000000000
base8:  170607300050000
base10: 8298866692096
base16: 78c3b005000
base32: 7hgtg0k00
existing: 1490753660736864259
```

## Author
[github](https://github.com/Daikonnbatake)
[twitter](https://twitter.com/_kagamaru_)

## License
[MIT](https://github.com/studigo/marinesnow/blob/main/LICENSE)