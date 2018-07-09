package mgologstash

import (
	"context"
	"time"

	"github.com/activecm/ipfix-rita/converter/ipfix"
)

//Reader implements ipfix.Reader
type Reader struct {
	buffer   Buffer
	pollWait time.Duration
}

//NewReader returns a new ipfix.Reader backed by a mgologstash.Buffer
func NewReader(buffer Buffer, pollWait time.Duration) ipfix.Reader {
	return Reader{
		buffer:   buffer,
		pollWait: pollWait,
	}
}

//Drain asynchronously drains a mgologstash.Buffer
func (r Reader) Drain(ctx context.Context) (<-chan ipfix.Flow, <-chan error) {
	out := make(chan ipfix.Flow)
	errors := make(chan error)

	go func(buffer Buffer, pollWait time.Duration, out chan<- ipfix.Flow, errors chan<- error) {
		pollTicker := time.NewTicker(pollWait)

		r.drainInner(ctx, buffer, out, errors)
	Loop:
		for {
			select {
			case <-ctx.Done():
				errors <- ctx.Err()
				break Loop
			case <-pollTicker.C:
				r.drainInner(ctx, buffer, out, errors)
			}
		}
		buffer.Close()
		pollTicker.Stop()
		close(errors)
		close(out)
	}(r.buffer, r.pollWait, out, errors)

	return out, errors
}

func (r Reader) drainInner(ctx context.Context, buffer Buffer, out chan<- ipfix.Flow, errors chan<- error) {
	flow := &Flow{}
	for buffer.Next(flow) {
		out <- flow
		//ensure we stop even if there is more data
		if ctx.Err() != nil {
			errors <- ctx.Err()
			break
		}
		flow = &Flow{}
	}
	if buffer.Err() != nil {
		errors <- buffer.Err()
	}
}