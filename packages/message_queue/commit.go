package message_queue

import "time"

type CommitManager struct {
	ticker     *time.Ticker
	TickerChan <-chan time.Time
}

func NewCommitManager(commitInterval time.Duration) *CommitManager {
	ticker := time.NewTicker(commitInterval)
	return &CommitManager{
		ticker:     ticker,
		TickerChan: ticker.C,
	}
}

func NewCommitManagerForTest(testTickerChan <-chan time.Time) *CommitManager {
	return &CommitManager{
		TickerChan: testTickerChan,
	}
}

func (c *CommitManager) Stop() {
	if c.ticker != nil {
		c.ticker.Stop()
	}
}
