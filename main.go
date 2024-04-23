package main

import (
	"fmt"
	"github.com/yuno-payments/yuno-go-utils-lib/constants"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithService("cfg.AppName"),
		tracer.WithRuntimeMetrics(),
	)
	fmt.Println("Hello, World!")
	fmt.Println(constants.GetAlpha3CountryCode("United States"))
}
