// Snowflake ID を生成するシンプルなライブラリ.
package marinesnow

import (
	"fmt"
	t "time"
)

var timestampOffset int64 // 採番開始日時.を保持する.
var workerID int64        // サーバー毎に割り振るIDを保持する.
var lastTimestamp int64   // 最後に採番した時間を保持する.
var increments int64      // 同一タイムスタンプのIDを区別するためのインクリメント値.

// workerID を設定する.
func SetWorkerID(id int64) error {

	// 負の値を弾く.
	if id < 0 {
		return fmt.Errorf("invalid value: %d", id)
	}

	// 10bitを超える数値を弾く.
	if WORKER_ID_MASK < id {
		return fmt.Errorf("invalid value: %d", id)
	}

	// workerID を設定する.
	workerID = id
	return nil
}

// 採番開始日時を設定する.
func SetTimestampOffset(offset int64) error {

	// 負の値を弾く.
	if offset < 0 {
		return fmt.Errorf("invalid value: %d", offset)
	}

	// 未来の時間を弾く.
	if t.Now().UnixMilli() < offset {
		return fmt.Errorf("invalid value: %d", offset)
	}

	// 採番開始日時を設定する.
	timestampOffset = offset
	return nil
}

// snowflakeを生成する.
func Generate() (Snowflake, error) {
	var time int64 = t.Now().UnixMilli()
	var newSF Snowflake

	// 時間が巻き戻っているなら弾く.
	if time < lastTimestamp {
		return newSF, fmt.Errorf("timestamp was rolled back")
	}

	// インクリメント値を設定する
	if time == lastTimestamp {
		increments++
	} else {
		increments = 0
	}

	// インクリメント値がカンストしたら弾く.
	if INCREMENTS_MASK < increments {
		return newSF, fmt.Errorf("limit reached")
	}

	// Snowflake を生成する.
	shiftedTimestamp := (time - timestampOffset) << TIMESTAMP_SHIFT
	shiftedWorkerID := workerID << WORKER_ID_SHIFT
	shiftedIncrements := increments << INCREMENTS_SHIFT

	newSF = Snowflake(shiftedTimestamp | shiftedWorkerID | shiftedIncrements)
	lastTimestamp = time

	return newSF, nil
}
