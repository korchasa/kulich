package context

//type dryRunCtxKeyType string
//
//const dryRunCtxKey dryRunCtxKeyType = "dry_run"
//
//func WithDryRun(ctx context.Context, dryRun bool) context.Context {
//	return context.WithValue(ctx, dryRunCtxKey, dryRun)
//}
//
//func DryRun(ctx context.Context) bool {
//	dryRun, ok := ctx.Value(dryRunCtxKey).(bool)
//	if !ok {
//		log.Errorf("can't extract dryRun key from context")
//		return true
//	}
//	return dryRun
//}
