package redis

import (
	"context"
)

// Pipeliner is an mechanism to realise Redis Pipeline technique.
//
// Pipelining is a technique to extremely speed up processing by packing
// operations to batches, send them at once to Redis and read a replies in a
// singe step.
// See https://redis.io/topics/pipelining
//
// Pay attention, that Pipeline is not a transaction, so you can get unexpected
// results in case of big pipelines and small read/write timeouts.
// Redis client has retransmission logic in case of timeouts, pipeline
// can be retransmitted and commands can be executed more then once.
// To avoid this: it is good idea to use reasonable bigger read/write timeouts
// depends of your batch size and/or use TxPipeline.
type CustomCmdable interface {
	StatefulCmdable
	Do(ctx context.Context, args ...interface{}) *Cmd
	Process(ctx context.Context, cmd Cmder) error
}

type CustomOperator interface {
	Process(ctx context.Context, cmd Cmder) error
	Pipeline() Pipeliner
	Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error)
}

var _ CustomCmdable = (*Custom)(nil)

// Pipeline implements pipelining as described in
// http://redis.io/topics/pipelining. It's safe for concurrent use
// by multiple goroutines.
type Custom struct {
	cmdable
	statefulCmdable

	ctx  context.Context
	exec pipelineExecer
	op   CustomOperator
}

func (c *Custom) init() {
	c.cmdable = c.Process
	c.statefulCmdable = c.Process
}

func (c *Custom) Do(ctx context.Context, args ...interface{}) *Cmd {
	cmd := NewCmd(ctx, args...)
	_ = c.Process(ctx, cmd)
	return cmd
}

// Process queues the cmd for later execution.
func (c *Custom) Process(ctx context.Context, cmd Cmder) error {
	return c.op.Process(ctx, cmd)
}

// Exec passes the given
func (c *Custom) Exec(ctx context.Context, cmds []Cmder) error {
	return c.exec(ctx, cmds)
}

func (c *Custom) Pipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	return c.op.Pipelined(ctx, fn)
}

func (c *Custom) Pipeline() Pipeliner {
	return c.op.Pipeline()
}

func (c *Custom) TxPipelined(ctx context.Context, fn func(Pipeliner) error) ([]Cmder, error) {
	return c.Pipelined(ctx, fn)
}

func (c *Custom) TxPipeline() Pipeliner {
	return c.op.Pipeline()
}