package recover

import "testing"

func TestRecoverCase1(t *testing.T) {
	// 无法recover
	defer t.Log(recover()) //     recover_test.go:7: <nil>
	panic(1)               // panic: 1
}

func TestRecoverCase2(t *testing.T) {
	// recover 必须在defer匿名函数里
	defer func() {
		res := recover() //  recover_test.go:15: 1
		t.Log(res)
	}()
	panic(1)
}

func TestRecoverCase3(t *testing.T) {
	// recover 必须在defer匿名函数里
	defer func() {
		t.Log(recover(), 1) //     recover_test.go:23: <nil> 1
	}()
	defer func() {
		t.Log(recover(), 2) // 先打印，先recover    recover_test.go:26: 1 2
	}()
	panic(1)
}

func TestRecoverCase4(t *testing.T) {
	// recover 必须在defer匿名函数里
	defer1 := func() {
		t.Log(recover(), 1) //    recover_test.go:34: 1 1
		panic(2)            // recover failed
	}
	defer defer1()
	panic(1)
}

func TestRecoverCase5(t *testing.T) {
	// recover 必须在defer匿名函数里
	defer1 := func() {
		t.Log(recover(), 1) //        recover_test.go:49: 1 2
		panic(2)
	}

	defer2 := func() {
		t.Log(recover(), 2) //     recover_test.go:49: 1 2
	}
	defer defer1()
	defer defer2()
	panic(1)
}

/*
=== RUN   TestRecoverCase5
    recover_test.go:49: 1 2
    recover_test.go:44: <nil> 1
--- FAIL: TestRecoverCase5 (0.00s)
panic: 2 [recovered]
	panic: 2
*/
