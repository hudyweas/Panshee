package util

import "context"

type Metadata struct {
	ClientIp string
	UserAgent string
}

func GetMetadatas(ctx context.Context) (md Metadata){
	md.ClientIp = ctx.Value("clientIP").(string)
	md.UserAgent = ctx.Value("userAgent").(string)

	return
}
