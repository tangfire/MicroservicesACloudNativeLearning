package main

import (
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"math/rand"
	"net/http"
)

var tracer = otel.Tracer("roll")

// 创建一个 tracer
func roll(w http.ResponseWriter, r *http.Request) {
	// 创建一个子span
	_, span := tracer.Start(r.Context(), "roll")
	defer span.End()

	// 业务逻辑
	number := 1 + rand.Intn(6)

	// 往span里记录属性
	rollValueAttr := attribute.Int("roll.value", number)
	span.SetAttributes(rollValueAttr)

	_, _ = fmt.Fprintln(w, number)

}
