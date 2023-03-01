// snowflake を定義するパッケージ
package snowflake

import "strconv"

// スノーフレークの各セクションのシフト数
const (
	TIMESTAMP_SHIFT  = 22
	WORKER_ID_SHIFT  = 12
	INCREMENTS_SHIFT = 0
)

// スノーフレークの各セクションのサイズに対応するbitマスク
const (
	TIMESTAMP_MASK  = 0x1FFFFFFFFFF
	WORKER_ID_MASK  = 0x3FF
	INCREMENTS_MASK = 0xFFF
)

// snowflake ID を表す型
type Snowflake int64

// Snowflake からタイムスタンプを取得する.
func (v Snowflake) Timestamp() int64 {
	return int64(v) >> TIMESTAMP_SHIFT & TIMESTAMP_MASK
}

// Snowflake からワーカーID を取得する.
func (v Snowflake) WorkerID() int64 {
	return int64(v) >> WORKER_ID_SHIFT & WORKER_ID_MASK
}

// Snowflake からインクリメント情報を取得する.
func (v Snowflake) Increments() int64 {
	return int64(v) >> INCREMENTS_SHIFT & INCREMENTS_MASK
}

// Snowflake を底を指定して文字列に変換する.
func (v Snowflake) ToString(base int) string {
	return strconv.FormatInt(int64(v), base)
}
