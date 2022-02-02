package logger

import (
	"context"
	"fmt"
)

type contextKey string

const (
	errValueNotFound = "not found"
)

func GetLogPrefix(ctx context.Context, isDebug bool) string {
	var prefix string

	if isDebug {
		prefix += "[DEBUG] "
	}

	if host, err := GetHostName(ctx); err == nil {
		prefix += fmt.Sprintf("[host: %s] ", host)
	}
	if vmName, err := GetVirtualMachineName(ctx); err == nil {
		prefix += fmt.Sprintf("[virtual machine: %s] ", vmName)
	}

	return prefix
}

var (
	keyHost           contextKey = "host"
	keyVirtualMachine contextKey = "virtual-machine"
)

func SetHostName(parents context.Context, hostName string) context.Context {
	return context.WithValue(parents, keyHost, hostName)
}
func GetHostName(ctx context.Context) (string, error) {
	v := ctx.Value(keyHost)

	hostname, ok := v.(string)
	if !ok {
		return "", fmt.Errorf(errValueNotFound)
	}
	return hostname, nil
}

func SetVirtualMachineName(parents context.Context, vmName string) context.Context {
	return context.WithValue(parents, keyVirtualMachine, vmName)
}
func GetVirtualMachineName(ctx context.Context) (string, error) {
	v := ctx.Value(keyVirtualMachine)

	vmName, ok := v.(string)
	if !ok {
		return "", fmt.Errorf(errValueNotFound)
	}
	return vmName, nil
}
