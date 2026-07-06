package glob

import "context"

type taskOutputContextKey struct{}

type taskOutput struct {
	value string
}

func WithTaskOutput(ctx context.Context) context.Context {
	return context.WithValue(ctx, taskOutputContextKey{}, &taskOutput{})
}

func SetTaskOutput(ctx context.Context, output string) {
	if v, ok := ctx.Value(taskOutputContextKey{}).(*taskOutput); ok {
		v.value = output
	}
}

func GetTaskOutput(ctx context.Context) string {
	if v, ok := ctx.Value(taskOutputContextKey{}).(*taskOutput); ok {
		return v.value
	}
	return ""
}
