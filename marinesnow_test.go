package marinesnow

import (
	"testing"
	"time"
)

// 要件を満たさない値をちゃんと弾くか,
// 正しい値をちゃんと設定できるかテストする.
func TestSetTimestampOffset(t *testing.T) {

	t.Run("正常値", func(t *testing.T) {
		now := time.Now().UnixMilli()
		err := SetTimestampOffset(now)

		if err != nil {
			t.Fail()
		}
	})

	t.Run("負の時間", func(t *testing.T) {
		minus := int64(-1)
		err := SetTimestampOffset(minus)

		if err == nil {
			t.Fail()
		}
	})

	t.Run("未来の時間", func(t *testing.T) {
		future := time.Now().UnixMilli() + 1
		err := SetTimestampOffset(future)

		if err == nil {
			t.Fail()
		}
	})
}

// 要件を満たさない値をちゃんと弾くか,
// 正しい値をちゃんと設定できるかテストする.
func TestSetWorkerID(t *testing.T) {

	t.Run("正常値", func(t *testing.T) {
		val := int64(5)
		err := SetWorkerID(val)

		if err != nil {
			t.Fail()
		}
	})

	t.Run("負の値", func(t *testing.T) {
		val := int64(-1)
		err := SetWorkerID(val)

		if err == nil {
			t.Fail()
		}
	})

	t.Run("10bitを超える値", func(t *testing.T) {
		val := int64(0x400)
		err := SetWorkerID(val)
		if err == nil {
			t.Fail()
		}
	})
}

// 要件を満たさない値をちゃんと弾くか,
// 正しいIDをちゃんと吐くかテストする.
func TestGenerate(t *testing.T) {

	t.Run("正常値", func(t *testing.T) {
		SetTimestampOffset(0)
		SetWorkerID(0)

		if _, err := Generate(); err != nil {
			t.Fail()
		}
	})

	t.Run("時間が巻き戻った場合", func(t *testing.T) {
		lastTimestamp = time.Now().UnixMilli() + 1

		if _, err := Generate(); err == nil {
			t.Fail()
		}

		SetTimestampOffset(0)
	})

	t.Run("インクリメント値がカンストした場合", func(t *testing.T) {
		SetTimestampOffset(0)
		increments = 0xFFF

		if _, err := Generate(); err == nil {
			t.Fail()
		}

		increments = 0
	})
}
