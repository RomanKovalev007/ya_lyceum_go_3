package step1log

import "log/slog"

func LogUserAction(logger *slog.Logger, user string, action string){
	logger.Info("user action",slog.String("User", user), slog.String("Action", action))
}