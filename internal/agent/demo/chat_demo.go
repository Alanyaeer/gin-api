package main

import (
	"chat-system/config"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/schema"
)
var agentCfg = config.Cfg.Agent
func DemoLLMQuickStart() {
	ctx := context.Background()

	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		// 如果您想使用 Azure OpenAI 服务，请设置这两个字段。
		// BaseURL: "https://{RESOURCE_NAME}.openai.azure.com",
		// ByAzure: true,
		// APIVersion: "2024-06-01",
		APIKey:  agentCfg.APIKey,
		Model:   agentCfg.Model,
		BaseURL: agentCfg.BaseURL,
		ByAzure: func() bool {
			return false
		}(),
		ReasoningEffort: openai.ReasoningEffortLevelLow,
	})
	if err != nil {
		log.Fatalf("NewChatModel failed, err=%v", err)
	}

	resp, err := chatModel.Generate(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "as a machine, how do you answer user's question?",
		},
	})
	if err != nil {
		log.Fatalf("Generate failed, err=%v", err)
	}
	fmt.Printf("output: \n%v", resp)

}

func DemoLLMQuickStartStream() {
	ctx := context.Background()
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		// 如果您想使用 Azure OpenAI 服务，请设置这两个字段。
		// BaseURL: "https://{RESOURCE_NAME}.openai.azure.com",
		// ByAzure: true,
		// APIVersion: "2024-06-01",
		APIKey:  agentCfg.APIKey,
		Model:   agentCfg.Model,
		BaseURL: agentCfg.BaseURL,
		ByAzure: func() bool {
			return false
		}(),
		
		ReasoningEffort: openai.ReasoningEffortLevelLow,
	})
	if err != nil {
		log.Fatalf("NewChatModel failed, err=%v", err)
	}

	streamMsgs, err := chatModel.Stream(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "为什么速度足够快，可以摆脱地球的引力",
		},
	})
	if err != nil {
		log.Fatalf("Stream of openai failed, err=%v", err)
	}

	defer streamMsgs.Close()

	fmt.Printf("typewriter output:")
	for {
		msg, err := streamMsgs.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Recv of streamMsgs failed, err=%v", err)
		}
		fmt.Print(msg.Content)
	}

	fmt.Print("\n")
}
