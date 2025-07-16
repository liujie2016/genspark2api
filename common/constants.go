package common

import "time"

var StartTime = time.Now().Unix() // unit: second
var Version = "v1.12.6"           // this hard coding will be replaced automatically when building, no need to manually change

var DefaultOpenaiModelList = []string{
	"gpt-4.1",
	"o3-pro",
	"o4-mini-high",
	"gemini-2.5-pro",
	"gemini-2.0-flash",
	"deep-seek-r1",
	"claude-sonnet-4",
	"claude-opus-4",
	"grok-4-0709",

	"flux",
	"flux-speed",
	"flux-pro/ultra",
	"ideogram",
	"ideogram/V_2A",
	"recraft-v3",
	"dall-e-3",
	"imagen4",
	"gpt-image-1",
	"flux-pro/kontext/pro",
	"flux-pro/kontext/max",
}

var TextModelList = []string{
	"gpt-4.1",
	"o3-pro",
	"o4-mini-high",
	"gemini-2.5-pro",
	"gemini-2.0-flash",
	"deep-seek-r1",
	"claude-sonnet-4",
	"claude-opus-4",
	"grok-4-0709",
}

var MixtureModelList = []string{
	"gpt-4o",
	"claude-3-7-sonnet",
	"gemini-2.0-flash",
}

var ImageModelList = []string{
	"flux",
	"flux-speed",
	"flux-pro/ultra",
	"ideogram",
	"ideogram/V_2A",
	"recraft-v3",
	"dall-e-3",
	"imagen4",
	"gpt-image-1",
	"flux-pro/kontext/pro",
	"flux-pro/kontext/max",
}

var VideoModelList = []string{
	"gemini/veo3",
	"gemini/veo3/fast",
	"kling/v2.1/master",
	"fal-ai/bytedance/seedance/v1/pro",
	"minimax/hailuo-02/standard",
	"pixverse/v4.5/turbo",
	"fal-ai/bytedance/seedance/v1/lite",
}

//
