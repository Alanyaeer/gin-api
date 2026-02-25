package agent

import (
	"chat-system/config"
	"chat-system/internal/model/customize/sse"
	"chat-system/pkg/protocol/header"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"

	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
)

// LLMStreamDemo 最简流式输出 demo
// GET /api/v1/agent/stream?prompt=你好
func LLMStreamDemo(ctx *gin.Context) {
	prompt := ctx.Query("prompt")
	if prompt == "" {
		prompt = "用一句话介绍你自己"
	}
	header.WrapperCtxHeaderForSse(ctx)

	
	agentCfg := config.Cfg.Agent
	chatModel, err := ark.NewChatModel(ctx.Request.Context(), &ark.ChatModelConfig{
		APIKey: agentCfg.APIKey,
		Model:  agentCfg.Model,
		Thinking: &model.Thinking{
			Type: model.ThinkingTypeDisabled,
		},
	})
	if err != nil {
		slog.Info(err.Error())
		ctx.String(http.StatusInternalServerError, "NewChatModel failed: %v", err)
		return
	}
	if err != nil {
		slog.Info(err.Error())
		ctx.String(http.StatusInternalServerError, "NewChatModel failed: %v", err)
		return
	}
	if err != nil {
		slog.Info(err.Error())
		ctx.String(http.StatusInternalServerError, "NewChatModel failed: %v", err)
		return
	}
	adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{

	})
	streamMsgs, err := chatModel.Stream(ctx.Request.Context(), []*schema.Message{
		{Role: schema.User, Content: prompt},
	})
	if err != nil {
		slog.Info(err.Error())
		ctx.String(http.StatusInternalServerError, "Stream failed: %v", err)
		return
	}
	defer streamMsgs.Close()

	for {
		msg, err := streamMsgs.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Recv failed: %v", err)
			return
		}
		if msg.Content == "" {
			continue
		}
		slog.Info(fmt.Sprintf("%v", msg.Content))
		eventResponse := sse.SseResponse{
			Content: msg.Content,
		}
		ctx.SSEvent("message", eventResponse)
	}
	ctx.Writer.Flush()
}
