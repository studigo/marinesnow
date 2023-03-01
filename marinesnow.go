// Snowflake ID を生成するシンプルなライブラリ.
package marinesnow

import (
	t "time"

	er "github.com/studigo/marinesnow/error"
	sf "github.com/studigo/marinesnow/snowflake"
)

var timestampOffset int64 // 採番開始日時.を保持する.
var workerID int64        // サーバー毎に割り振るIDを保持する.
var lastTimestamp int64   // 最後に採番した時間を保持する.
var increments int64      // 同一タイムスタンプのIDを区別するためのインクリメント値.

// workerID を設定する.
func SetWorkerID(id int64) error {

	// 負の値を弾く.
	if id < 0 {
		return er.InvalidValue(id)
	}

	// 10bitを超える数値を弾く.
	if sf.WORKER_ID_MASK < id {
		return er.InvalidValue(id)
	}

	// workerID を設定する.
	workerID = id
	return nil
}

// 採番開始日時を設定する.
func SetTimestampOffset(offset int64) error {

	// 負の値を弾く.
	if offset < 0 {
		return er.InvalidValue(offset)
	}

	// 未来の時間を弾く.
	if t.Now().UnixMilli() < offset {
		return er.InvalidValue(offset)
	}

	// 採番開始日時を設定する.
	timestampOffset = offset
	return nil
}

// snowflakeを生成する.
func Generate() (sf.Snowflake, error) {
	var time int64 = t.Now().UnixMilli()
	var newSF sf.Snowflake

	// 時間が巻き戻っているなら弾く.
	if time < lastTimestamp {
		return newSF, er.TimeIsRewound()
	}

	// インクリメント値を設定する
	if time == lastTimestamp {
		increments++
	} else {
		increments = 0
	}

	// インクリメント値がカンストしたら弾く.
	if sf.INCREMENTS_MASK < increments {
		return newSF, er.NoMoreID()
	}

	// Snowflake ID を生成する.
	shiftedTimestamp := (time - timestampOffset) << sf.TIMESTAMP_SHIFT
	shiftedWorkerID := workerID << sf.WORKER_ID_SHIFT
	shiftedIncrements := increments << sf.INCREMENTS_SHIFT

	newSF = sf.Snowflake(shiftedTimestamp | shiftedWorkerID | shiftedIncrements)
	lastTimestamp = time

	return newSF, nil
}
