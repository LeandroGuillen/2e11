package aiplayer

import "testing"

func TestGetNextMove(t *testing.T) {
  
  v := getNextMove()
  
  if v < 0 || v > 3 {
    t.Error("Returned value is out of range")
  }
}
