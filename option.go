package wstatic

import "embed"

type Option func(*Options)

type Options struct {
	fs        embed.FS
	root      string
	indexes   bool
	urlPrefix string
	embed     bool
}

func WithEmbedFs(fs embed.FS) Option {
	return func(opts *Options) {
		opts.fs = fs
	}
}

func WithRoot(root string) Option {
	return func(opts *Options) {
		opts.root = root
	}
}

func WithIndexes(indexes bool) Option {
	return func(opts *Options) {
		opts.indexes = indexes
	}
}

func WithUrlPrefix(urlPrefix string) Option {
	return func(opts *Options) {
		opts.urlPrefix = urlPrefix
	}
}

func WithEmbed(embed bool) Option {
	return func(opts *Options) {
		opts.embed = embed
	}
}
