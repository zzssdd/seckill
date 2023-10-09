package task

import "context"

func StartAllTask(ctx context.Context) {
	go StartDelayTask(ctx)
	go StartBackgroundTask(ctx)
}
