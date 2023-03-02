package middlewares

import (
	"context"
	"fmt"

	"github.com/dollarkillerx/common/pkg/logger"
	"github.com/dollarkillerx/eim/internal/utils"
	"runtime/debug"
)

// RecoverFunc ...
func RecoverFunc(ctx context.Context, err interface{}) (userMessage error) {
	logger.Errorf("[ReqID]:%s\n[Message]:%+v\n[Panic]:%s", utils.GetRequestIdByContext(ctx), err, string(debug.Stack()))
	return fmt.Errorf("Panic: %+v", err)
}
