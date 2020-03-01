package ioencode

type TestWriter struct {
  buffer []byte
  err error
}

func (w *TestWriter) Write(p []byte) (n int, err error) {
  if w.err != nil {
    return 0, w.err
  }

  w.buffer = append(w.buffer, p...)

  return len(p), nil
}
