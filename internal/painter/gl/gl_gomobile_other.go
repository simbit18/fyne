// +build !ios

package gl

func (p *glPainter) glFreeBuffer(_ Buffer) {
	// we don't free a shared buffer!
}
