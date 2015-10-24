package party_test

import (
    "github.com/ymonk/party"
    "testing"
    "fmt"
)

func TestDynamicTemplate(t *testing.T) {
    th := party.New("main1.html", nil, true)
    fmt.Println(th.String())
}

func TestStaticTemplate(t *testing.T) {
    th := party.New("main2.html", nil, false)
    fmt.Println(th.String())
}


