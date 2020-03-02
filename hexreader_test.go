package ioencode

import (
  "errors"
  "fmt"
  "testing"
  "encoding/hex"
)

func TestHexReader(t *testing.T) {
  anError := fmt.Errorf("an error")

  tcs := []struct{
    description string
    input string
    buffer [][]byte
    output []string
    n []int
    inputErr []error
    outputErr []error
  }{
    {
      "Empty",
      "",
      [][]byte{make([]byte, 0, 1024)},
      []string{""},
      []int{0},
      []error{nil},
      []error{nil},
    },
    {
      "1 byte",
      "41",
      [][]byte{make([]byte, 1024)},
      []string{"A"},
      []int{1},
      []error{nil},
      []error{nil},
    },
    {
      "medium string",
      "54686520717569636b20626c61636b20666f78206a756d706564206f76657220746865206c617a7920646f6773",
      [][]byte{make([]byte, 1024)},
      []string{"The quick black fox jumped over the lazy dogs"},
      []int{45},
      []error{nil},
      []error{nil},
    },
    {
      "two medium strings",
      "54686520717569636b20626c61636b20666f78206a756d706564206f76657220746865206c617a7920646f6773",
      [][]byte{make([]byte, 20), make([]byte, 1024)},
      []string{"The quick black fox ", "jumped over the lazy dogs"},
      []int{20, 25},
      []error{nil, nil},
      []error{nil, nil},
    },
    {
      "error",
      "54686520717569636b20626c61636b20666f78206a756d706564206f76657220746865206c617a7920646f6773",
      [][]byte{make([]byte, 1024)},
      []string{""},
      []int{0},
      []error{anError},
      []error{anError},
    },
    {
      "read then error",
      "54686520717569636b20626c61636b20666f78206a756d706564206f76657220746865206c617a7920646f6773",
      [][]byte{make([]byte, 20), make([]byte, 1024)},
      []string{"The quick black fox ", ""},
      []int{20, 0},
      []error{nil, anError},
      []error{nil, anError},
    },
    {
      "bad string",
      "414243 not hex anymore",
      [][]byte{make([]byte, 1024)},
      []string{"ABC"},
      []int{3},
      []error{nil},
      []error{hex.InvalidByteError(' ')},
    },
  }

  for _, tc := range tcs {
    t.Run(tc.description, func(t *testing.T) {
      tr := &TestReader{
       buffer: []byte(tc.input),
      }

      r := NewHexReader(tr)

      for i := range tc.buffer {
        tr.err = tc.inputErr[i]
        n, err := r.Read(tc.buffer[i])
        tc.buffer[i] = tc.buffer[i][:n]

        if n != tc.n[i] {
          t.Errorf("expected n=%d but received %d", tc.n[i], n)
        }

        if !errors.Is(err, tc.outputErr[i])  {
          t.Errorf("expected err=%v but received %v", tc.outputErr[i], err)
        }
        
        if string(tc.buffer[i]) != tc.output[i] {
          t.Errorf("expected output %d(:%v) but received %d(%v)", len(tc.output[i]), tc.output[i], len(tc.buffer[i]), string(tc.buffer[i])) 
        }
      }
    })
  }
}

