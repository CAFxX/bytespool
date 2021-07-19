package bytespool

// Append appends the specified bytes to the buf byte slice.
// The buf byte slice passed to Append MUST NOT be used after
// Append is called: only the byte slice returned should be
// used.
func Append(buf []byte, v ...byte) []byte {
  if len(v) <= cap(buf)-len(buf) {
    return append(buf, v...)
  }
  return appendSlow(buf, v...)
}

func appendSlow(buf []byte, v ...byte) []byte {
  pbuf := GetBytesSlicePtr((len(v)+len(buf))*2)
  nbuf := append(append((*pbuf)[:0], buf...), v...)
  *pbuf = buf
  PutBytesSlicePtr(pbuf)
  return nbuf
}
