package resolvers

//// ChatGpt ...
//func (r *mutationResolver) ChatGpt(ctx context.Context, input *generated.ChatGPTData) (*generated.ChatGPTResp, error) {
//	apiEndpoint := "https://api.openai.com/v1/completions"
//	accessToken := "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"
//
//	return &generated.ChatGPTResp{
//		Message: "hello world",
//	}, nil
//
//	params := map[string]interface{}{
//		"prompt": strings.TrimSpace(input.Message),
//		//"temperature":       0.7,
//		//"max_tokens":        500,
//		//"top_p":             1,
//		//"frequency_penalty": 0.5,
//		//"presence_penalty":  0.0,
//
//		"model":             "text-davinci-003",
//		"temperature":       0.9,
//		"max_tokens":        999,
//		"top_p":             1,
//		"frequency_penalty": 0.0,
//		"presence_penalty":  0.6,
//		"stop":              []string{".end."},
//	}
//
//	rcode, rdata, err := urllib.Post(apiEndpoint).SetTimeout(time.Second*30).SetHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).SetJsonObject(params).Byte()
//	if err != nil {
//		panic(err)
//	}
//
//	if rcode != 200 {
//		return nil, errors.New(string(rdata))
//	}
//
//	return &generated.ChatGPTResp{
//		Message: string(rdata),
//	}, nil
//}
