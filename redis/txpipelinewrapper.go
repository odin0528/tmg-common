package redis

func (tp *TxPipelineWrapper) Discard() error {
	return tp.pipe.Discard()
}
