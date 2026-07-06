package middleware

import (
	"context"
	"testing"

	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hibiken/asynq"
)

type crontabLogStub struct {
	output string
}

func (s *crontabLogStub) Model(ctx context.Context) *gdb.Model {
	return nil
}

func (s *crontabLogStub) GetPageList(ctx context.Context, req *model.PageListReq, in *req.SettingCrontabLogSearch) (rs []*res.SettingCrontabLog, total int, err error) {
	return nil, 0, nil
}

func (s *crontabLogStub) Delete(ctx context.Context, ids []int64) error {
	return nil
}

func (s *crontabLogStub) AddLog(ctx context.Context, id int64, status int, exceptionInfo string, startTime *gtime.Time, endTime *gtime.Time, output string) error {
	s.output = output
	return nil
}

func TestLoggingMiddlewareSavesWorkerOutput(t *testing.T) {
	logStub := &crontabLogStub{}
	service.RegisterSettingCrontabLog(logStub)

	handler := LoggingMiddleware(asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		glob.SetTaskOutput(ctx, "worker output")
		return nil
	}))
	task := asynq.NewTask("test", []byte(`{"crontab_id":1,"data":{}}`))

	if err := handler.ProcessTask(context.Background(), task); err != nil {
		t.Fatalf("ProcessTask() error = %v", err)
	}
	if logStub.output != "worker output" {
		t.Fatalf("AddLog output = %q, want %q", logStub.output, "worker output")
	}
}
