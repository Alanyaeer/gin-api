package agent

import (
	"chat-system/internal/api/agent"
	"github.com/gin-gonic/gin"
)

func RegisterAgentRoutes(r *gin.RouterGroup) {
	agentGroup := r.Group("/agent")
	{
		agentGroup.GET("/stream", agent.LLMStreamDemo)
	}
}
