package ioencode

import (
  "errors"
  "fmt"
  "testing"
)

func TestBinaryWriter(t *testing.T) {
  anError := fmt.Errorf("an error")

  tcs := []struct{
    description string
    input []string
    output []string
    n []int
    err []error
  }{
    {
      "Empty",
      []string{""},
      []string{""},
      []int{0},
      []error{nil},
    },
    {
      "1 char",
      []string{"A"},
      []string{"01000001"},
      []int{1},
      []error{nil},
    },
    {
      "medium string",
      []string{"The quick black fox jumped over the lazy dogs"},
      []string{"010101000110100001100101001000000111000101110101011010010110001101101011001000000110001001101100011000010110001101101011001000000110011001101111011110000010000001101010011101010110110101110000011001010110010000100000011011110111011001100101011100100010000001110100011010000110010100100000011011000110000101111010011110010010000001100100011011110110011101110011"},
      []int{45},
      []error{nil},
    },
    {
      "two medium strings",
      []string{"The quick black fox ", "jumped over the lazy dogs"},
      []string{"0101010001101000011001010010000001110001011101010110100101100011011010110010000001100010011011000110000101100011011010110010000001100110011011110111100000100000", "010101000110100001100101001000000111000101110101011010010110001101101011001000000110001001101100011000010110001101101011001000000110011001101111011110000010000001101010011101010110110101110000011001010110010000100000011011110111011001100101011100100010000001110100011010000110010100100000011011000110000101111010011110010010000001100100011011110110011101110011"},
      []int{20, 25},
      []error{nil, nil},
    },
    {
      "simple error",
      []string{"The quick black fox jumped over the lazy dogs"},
      []string{""},
      []int{0},
      []error{anError},
    },
    {
      "write then error",
      []string{"The quick black fox ", "jumped over the lazy dogs"},
      []string{"0101010001101000011001010010000001110001011101010110100101100011011010110010000001100010011011000110000101100011011010110010000001100110011011110111100000100000", "0101010001101000011001010010000001110001011101010110100101100011011010110010000001100010011011000110000101100011011010110010000001100110011011110111100000100000"},
      []int{20, 0},
      []error{nil, anError},
    },
  }

  for _, tc := range tcs {
    t.Run(tc.description, func(t *testing.T) {
      tw := &TestWriter{}

      w := NewBinaryWriter(tw)

      for i := range tc.input {
        tw.err = tc.err[i]
        n, err := w.Write([]byte(tc.input[i]))

        if n != tc.n[i] {
          t.Errorf("expected n=%d but received %d", tc.n[i], n)
        }

        if !errors.Is(err, tc.err[i])  {
          t.Errorf("expected err=%v but received %v", tc.err[i], err)
        }
        
        if string(tw.buffer) != tc.output[i] {
          t.Errorf("expected output=%v but received %v", tc.output[i], string(tw.buffer)) 
        }
      }
    })
  }
}
