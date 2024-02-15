package config

type Article struct {
	MaxVoteCount   int    `mapstructure:"max-vote-count" json:"max-vote-count" yaml:"max-vote-count"`    // 每人每类每天最大投票次数
	SubmissionTime string `mapstructure:"submission-time" json:"submission-time" yaml:"submission-time"` // 投稿时间段
	VoteTime       string `mapstructure:"vote-time" json:"vote-time" yaml:"vote-time"`                   // 投票时间段
	OpenVerify     bool   `mapstructure:"open-verify" json:"open-verify" yaml:"open-verify"`             // 开启验证
}
