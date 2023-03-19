package go_bitmap

func (m *BitMap) LoadPlain(option *Option) *BitMap {
	m.Image = pixelsToImage(option.Width, option.Height, option.BitPerPixel, option.Data)
	return m
}
