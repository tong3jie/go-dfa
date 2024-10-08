package dfa

const (
	defaultInvalidWorlds = " ,~,!,@,#,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,¥,·"
	defaultReplaceStr    = "****"
)

type Option func(opts *Options)

type Options struct {
	star         int
	question     int
	defaultStr   string
	invalidWords string
}

func WithStar(longth int) Option {
	return func(opts *Options) {
		opts.star = longth
	}
}

func WithQuestion(longth int) Option {
	return func(opts *Options) {
		opts.question = longth
	}
}

func WithDefaultStr(defaultStr string) Option {
	return func(opts *Options) {
		opts.defaultStr = defaultStr
	}
}

func WithInvalidWords(invalidWords string) Option {
	return func(opts *Options) {
		opts.invalidWords = invalidWords
	}
}
func loadOptions(options ...Option) *Options {
	opts := new(Options)
	for _, option := range options {
		option(opts)
	}
	return opts
}
