package models

import (
    "fmt"
    "testing"
)

func TestHashing(t *testing.T) {

    b := []byte("hello world")

    fmt.Printf(
        "%x \n",
        hash128(b),
    )

    fmt.Printf(
        "%x \n",
        hash256(b),
    )

}
