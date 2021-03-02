package messagebus

import (
	"fmt"
	"path"

	uuid "github.com/satori/go.uuid"
)

const (
	// agentTopicPrefix is the prefix for messages to specifc agents.
	agentTopicPrefix = "Agent"
	// c2vTopicPrefix is the prefix for all message topics from cloud domain to local NATS domain.
	c2vTopicPrefix = "c2v"
	// v2cTopicPrefix is the prefix for all message topics sent from local NATS to cloud domain.
	v2cTopicPrefix = "v2c"
)

// V2CTopic returns the topic used in the Vizier NATS domain to send messages from Vizier to Cloud.
func V2CTopic(topic string) string {
	// TODO(zasgar/michelle): Add validation.
	return fmt.Sprintf("%s.%s", v2cTopicPrefix, topic)
}

// C2VTopic returns the topic used in the Vizier NATS domain to get messages from the Cloud.
func C2VTopic(topic string) string {
	// TODO(zasgar/michelle): Add validation.
	return fmt.Sprintf("%s.%s", c2vTopicPrefix, topic)
}

// AgentUUIDTopic topic used to communicate to a specific agent.
func AgentUUIDTopic(agentID uuid.UUID) string {
	return AgentTopic(agentID.String())
}

// AgentTopic topic used to communicate to a specific agent.
func AgentTopic(agentID string) string {
	return path.Join(agentTopicPrefix, agentID)
}
