// marinesnow 内で使用するエラーを定義するパッケージ.
package error

import "fmt"

// 不正値が入力された場合に使用するエラー.
func InvalidValue(value int64) error {
	return fmt.Errorf("invalid value '%d'", value)
}

// 時間が巻き戻った場合に使用するエラー.
func TimeIsRewound() error {
	return fmt.Errorf("timestamp was rolled back")
}

// IDの生成限界に達した時に使用するエラー.
func NoMoreID() error {
	return fmt.Errorf("limit reached")
}
