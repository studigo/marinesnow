package marinesnow

import "testing"

// タイムスタンプを正常に取得できるかをテストする.
func TestSnowflakeTimestamp(t *testing.T) {
	sf := Snowflake(1630499798209155075)
	result := sf.Timestamp()
	expect := int64(388741445114)

	if result != expect {
		t.Error("\n出力: ", result, "\n期待: ", expect)
	}
}

// workerID を正常に取得できるかをテストする.
func TestSnowflakeWorkerID(t *testing.T) {
	sf := Snowflake(1630499798209155075)
	result := sf.WorkerID()
	expect := int64(421)

	if result != expect {
		t.Error("\n出力: ", result, "\n期待: ", expect)
	}
}

// インクリメント値を正常に取得できるかをテストする.
func TestSnowflakeIncrements(t *testing.T) {
	sf := Snowflake(1630499798209155075)
	result := sf.Increments()
	expect := int64(3)

	if result != expect {
		t.Error("\n出力: ", result, "\n期待: ", expect)
	}
}
