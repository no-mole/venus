package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/no-mole/venus/agent/venus/auth"
	"strings"
)

// var token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYyI6eyJleHAiOjI1NDEwNTQ3NzZ9LCJ0dCI6ImFkIn0.zJUjZSQUe1uQpjVJrD30t7bKKeAbxbFeOSV0xUws7v8`
// var token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYyI6eyJleHAiOjI1NDEwNjAyNTZ9LCJ0dCI6ImFkIn0.e3F_p4Ns36NldlN07RLh3jYIOKplyBC7N_tzuUQ-_oA`
var token = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyYyI6eyJleHAiOjI1NDExMTUyMTV9LCJ0dCI6ImFkIn0.ua1gthH6eh4XsqAHMERYoQvMXJf_AP7WCueArJ03AHU`
var peerToken = `uXxPXZ3cMGb35kDPvL5OVQ`

func main() {
	data := auth.Claims{}
	aaa := `{"rc":{"exp":2541054547},"tt":"ad"}`
	err := json.Unmarshal([]byte(aaa), &data)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", data)

	parts := strings.Split(token, ".")
	part0, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		panic(err.Error())
	}
	println("part0", string(part0))
	part1, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		panic(err.Error())
	}
	println("part1", string(part1))

	tp := auth.NewTokenProvider([]byte(peerToken))
	token, err := tp.Parse(context.Background(), token)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", token)
	fmt.Printf("%+v\n", token.Claims)
}
