package ioencode

import (
  "errors"
  "fmt"
  "testing"
)

func TestBinaryReader(t *testing.T) {
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
      "01000001",
      [][]byte{make([]byte, 1024)},
      []string{"A"},
      []int{1},
      []error{nil},
      []error{nil},
    },
    {
      "medium string",
      "010101000110100001100101001000000111000101110101011010010110001101101011001000000110001001101100011000010110001101101011001000000110011001101111011110000010000001101010011101010110110101110000011001010110010000100000011011110111011001100101011100100010000001110100011010000110010100100000011011000110000101111010011110010010000001100100011011110110011101110011",
      [][]byte{make([]byte, 1024)},
      []string{"The quick black fox jumped over the lazy dogs"},
      []int{45},
      []error{nil},
      []error{nil},
    },
    {
      "two medium strings",
      "010101000110100001100101001000000111000101110101011010010110001101101011001000000110001001101100011000010110001101101011001000000110011001101111011110000010000001101010011101010110110101110000011001010110010000100000011011110111011001100101011100100010000001110100011010000110010100100000011011000110000101111010011110010010000001100100011011110110011101110011",
      [][]byte{make([]byte, 20), make([]byte, 1024)},
      []string{"The quick black fox ", "jumped over the lazy dogs"},
      []int{20, 25},
      []error{nil, nil},
      []error{nil, nil},
    },
    {
      "error",
      "010101000110100001100101001000000111000101110101011010010110001101101011001000000110001001101100011000010110001101101011001000000110011001101111011110000010000001101010011101010110110101110000011001010110010000100000011011110111011001100101011100100010000001110100011010000110010100100000011011000110000101111010011110010010000001100100011011110110011101110011",
      [][]byte{make([]byte, 1024)},
      []string{""},
      []int{0},
      []error{anError},
      []error{anError},
    },
    {
      "read then error",
      "0101010001101000011001010010000001110001011101010110100101100011011010110010000001100010011011000110000101100011011010110010000001100110011011110111100000100000 some non binary",
      [][]byte{make([]byte, 20), make([]byte, 1024)},
      []string{"The quick black fox ", ""},
      []int{20, 0},
      []error{nil, anError},
      []error{nil, anError},
    },
    {
      "bad string",
      "010000010100001001000011 not binary anymore",
      [][]byte{make([]byte, 1024)},
      []string{"ABC"},
      []int{3},
      []error{nil},
      []error{ErrDecode(' ')},
    },
  }

  for _, tc := range tcs {
    t.Run(tc.description, func(t *testing.T) {
      tr := &TestReader{
       buffer: []byte(tc.input),
      }

      r := NewBinaryReader(tr)

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

