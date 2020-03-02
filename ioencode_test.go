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

type TestReader struct {
  buffer []byte
  err error
}

func (r *TestReader) Read(p []byte) (n int, err error) {
  if r.err != nil {
    return 0, r.err
  }

  toCopy := len(p)
  if toCopy > len(r.buffer) {
    toCopy = len(r.buffer)
  }

  copy(p, r.buffer)
  r.buffer = r.buffer[toCopy:]

  return toCopy, nil
}
