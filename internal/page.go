package internal

// IPage defines the methods of a page
type IPage interface {
	// GetOffset get the offset of the current page
	GetOffset() uint64
	// GetCurrentPage get the value of the current page
	GetCurrentPage() uint64
	// GetItemsPerPage get the value of the quantity of items per page
	GetItemsPerPage() uint64

	getPage() page
	calculate()
}

type page struct {
	page         uint64
	itemsPerPage uint64
	offset       uint64
}

func newPage(itemsPerPage uint64) IPage {
	return &page{itemsPerPage: itemsPerPage}
}

// GetOffset get the offset of the current page
func (p *page) GetOffset() uint64 {
	return p.offset
}

// GetItemsPerPage get the value of the quantity of items per page
func (p *page) GetItemsPerPage() uint64 {
	return p.itemsPerPage
}

// GetCurrentPage get the value of the current page
func (p *page) GetCurrentPage() uint64 {
	return p.page
}

func (p *page) calculate() {
	p.page++
	p.offset = (p.page - 1) * p.itemsPerPage
}

func (p *page) getPage() page {
	return *p
}
