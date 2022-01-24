package test

import (
	"context"
	"goframe-serverless-template/internal/cmd"
	"goframe-serverless-template/internal/misc"
	"os"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/stretchr/testify/assert"
)

var ctx context.Context

func initServer() {
	ctx = gctx.New()
	go func() {
		cmd.Main.Run(ctx)
	}()

	for {
		response := g.Client().GetVar(ctx, "http://localhost:8080/hello")
		if !response.IsNil() {
			g.Log().Debug(ctx, "init success, response =", response)
			break
		}
	}
	resetDB()
}

func resetDB() {
	err := misc.ExecSQL(ctx)
	if err != nil {
		panic(err)
	}
}

func TestMain(m *testing.M) {
	initServer()
	code := m.Run()
	os.Exit(code)
}

func assertResponseSuccess(t *testing.T, response *gvar.Var) {
	isSuccess := response.MapStrVar()["code"] != nil && response.MapStrVar()["code"].Int() == 0
	if !isSuccess {
		t.Logf("failed response = %v", response)
	}
	assert.True(t, isSuccess)
}

func TestTask(t *testing.T) {
	g.Client().PostVar(ctx, "http://localhost:8080/task", g.Map{"title": "task3"})

	response := g.Client().GetVar(ctx, "http://localhost:8080/task")
	assertResponseSuccess(t, response)
	assert.Equal(t, 1, len(response.MapStrVar()["data"].Slice()))
}
