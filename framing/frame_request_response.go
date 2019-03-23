package framing

import (
	"fmt"
	"github.com/rsocket/rsocket-go/common"
)

// FrameRequestResponse is frame for requesting single response.
type FrameRequestResponse struct {
	*BaseFrame
}

// Validate returns error if frame is invalid.
func (p *FrameRequestResponse) Validate() (err error) {
	return
}

func (p *FrameRequestResponse) String() string {
	m, _ := p.MetadataUTF8()
	return fmt.Sprintf("FrameRequestResponse{%s,data=%s,metadata=%s}", p.header, p.DataUTF8(), m)
}

// Metadata returns metadata bytes.
func (p *FrameRequestResponse) Metadata() ([]byte, bool) {
	return p.trySliceMetadata(0)
}

// Data returns data bytes.
func (p *FrameRequestResponse) Data() []byte {
	return p.trySliceData(0)
}

// MetadataUTF8 returns metadata as UTF8 string.
func (p *FrameRequestResponse) MetadataUTF8() (metadata string, ok bool) {
	raw, ok := p.Metadata()
	if ok {
		metadata = string(raw)
	}
	return
}

// DataUTF8 returns data as UTF8 string.
func (p *FrameRequestResponse) DataUTF8() string {
	return string(p.Data())
}

// NewFrameRequestResponse returns a new RequestResponse frame.
func NewFrameRequestResponse(id uint32, data, metadata []byte, flags ...FrameFlag) *FrameRequestResponse {
	fg := newFlags(flags...)
	bf := common.BorrowByteBuffer()
	if len(metadata) > 0 {
		fg |= FlagMetadata
		_ = bf.WriteUint24(len(metadata))
		_, _ = bf.Write(metadata)
	}
	if len(data) > 0 {
		_, _ = bf.Write(data)
	}
	return &FrameRequestResponse{
		&BaseFrame{
			header: NewFrameHeader(id, FrameTypeRequestResponse, fg),
			body:   bf,
		},
	}
}
