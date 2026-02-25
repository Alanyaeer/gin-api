package agent

import (
	"chat-system/internal/api/agent"
	"github.com/gin-gonic/gin"
)

type AgentRouteRegistrar struct{}

func NewRouter() *AgentRouteRegistrar {
	return &AgentRouteRegistrar{}
}

func (a *AgentRouteRegistrar) RegisterRoutes(r *gin.RouterGroup) {
	agentGroup := r.Group("/agent")
	{
		agentGroup.GET("/stream", agent.LLMStreamDemo)
	}
}