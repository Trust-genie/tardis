package storage

import (
	"sync/atomic"
	"tardis/internals/logger"
	"testing"
)

func TestMain(m *testing.M) {
	logger.Log.Info("Starting Tests")
	testStore = *NewStore(10, 15)

	//new function
	chaos := func(n int32) int32 {
		v := atomic.AddInt32(&n, 1) - 1
		return v
	}

	testStore.Put("Strings", "I meant what I said and I said what i Meant")
	testStore.Put("List of Strings", []string{"Horton", "Little Red Riding Hood", "The three little Pigs"})
	testStore.Put("Digits", 23)
	testStore.Put("List of Digits", []int{0, 10000, -23, 46})
	testStore.Put("Structs", struct {
		Name   string
		School string
		Age    uint
	}{"Harry Potter", "Hogworths?", 14})
	testStore.Put("Functions", chaos)

	m.Run()

	logger.Log.Info("Test Ended.")
}

var testStore Storage

type testStruct struct {
	string
}

func TestStorage_Put(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Positive Tests: Strings", &testStore, args{"El classico Score", "5-1"}, false},
		{"Positive Tests: list of strings", &testStore, args{"list", []string{"Jonathan", "Livingston", "Seagull"}}, false},
		{"Posirive Tests: Digits", &testStore, args{"key", 8081}, false},
		{"Positive Tests: List of Digits", &testStore, args{"nums", []int{-21, 0, 4, 7, 14, -0}}, false},
		{"Positive Tests: Structs", &testStore, args{"Structs A", testStruct{"Never"}}, false},
		{"Positive Tests:list of Structs", &testStore, args{"Tester Struct", []testStruct{
			testStruct{"Never"},
			testStruct{"Gonna"},
			testStruct{"Give"},
			testStruct{"You"},
			testStruct{"Up"},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Put(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Storage.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkStorage_Put(b *testing.B) {
	new := NewStore(4, 10)
	for i:= 0; i < b.N; i++{
		err := new.Put("Key", "value")
		if err != nil{
			b.Skip()
		}
	}
}

func TestStorage_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		s    *Storage
		args args
		//want    interface{}
		wantErr bool
	}{
		// Positive Tests
		{"Positive Test: strings", &testStore, args{"Strings"}, false},
		{"Positive Test: List of strings", &testStore, args{"List of Strings"}, false},
		{"Positive Test: Integers", &testStore, args{"Digits"}, false},
		{"Positive Test: List of Integers", &testStore, args{"List of Digits"}, false},
		{"Positive Test: Structs", &testStore, args{"Structs"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
