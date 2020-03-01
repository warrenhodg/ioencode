package ioencode

import (
  "errors"
  "fmt"
  "testing"
)

func TestHexWriter(t *testing.T) {
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
      []string{"41"},
      []int{1},
      []error{nil},
    },
    {
      "medium string",
      []string{"The quick black fox jumped over the lazy dogs"},
      []string{"54686520717569636b20626c61636b20666f78206a756d706564206f76657220746865206c617a7920646f6773"},
      []int{45},
      []error{nil},
    },
    {
      "two medium strings",
      []string{"The quick black fox ", "jumped over the lazy dogs"},
      []string{"54686520717569636b20626c61636b20666f7820", "54686520717569636b20626c61636b20666f78206a756d706564206f76657220746865206c617a7920646f6773"},
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
      []string{"54686520717569636b20626c61636b20666f7820", "54686520717569636b20626c61636b20666f7820"},
      []int{20, 0},
      []error{nil, anError},
    },
  }

  for _, tc := range tcs {
    t.Run(tc.description, func(t *testing.T) {
      tw := &TestWriter{}

      w := NewHexWriter(tw)

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
