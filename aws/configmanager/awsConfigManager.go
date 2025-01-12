package configmanager

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AwsConfigManager struct {
	ctx context.Context
	cfg aws.Config
}

func NewAwsConfigManager() *AwsConfigManager {
	return &AwsConfigManager{}
}

func (i *AwsConfigManager) InitConfig() {

	var err error

	i.ctx = context.TODO()
	i.cfg, err = config.LoadDefaultConfig(i.ctx, config.WithDefaultRegion("us-east-2"))

	if err != nil {
		panic(fmt.Sprintf(errSessionIntMsg, err.Error()))
	}
}

func (i *AwsConfigManager) GetContext() *context.Context {
	return &i.ctx
}

func (i *AwsConfigManager) GetConfig() aws.Config {
	return i.cfg
}
