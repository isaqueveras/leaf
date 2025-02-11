package internal

// IPage defines the methods of a page
type IPage interface {
	// GetOffset get the offset of the current page
	GetOffset() uint64
	// GetCurrentPage get the value of the current page
	GetCurrentPage() uint64
	// GetItemsPerPage get the value of the quantity of items per page
	GetItemsPerPage() uint64

	// GetStartCursor returns the starting ID of the current page in cursor-based pagination.
	GetCursor() uint64
	// SetCursor set the starting ID of the current page in cursor-based pagination
	SetCursor(uint64)

	getPage() *page
	calculate()
}

type page struct {
	page         uint64
	itemsPerPage uint64
	offset       uint64

	cursor uint64
}

func newPage(itemsPerPage uint64) IPage {
	return &page{itemsPerPage: itemsPerPage}
}

// GetStartCursor returns the starting ID of the current page in cursor-based pagination
func (p *page) GetCursor() uint64 { return p.cursor }

// SetCursor set the starting ID of the current page in cursor-based pagination
func (p *page) SetCursor(cursor uint64) { p.cursor = cursor }

// GetOffset get the offset of the current page
func (p *page) GetOffset() uint64 { return p.offset }

// GetItemsPerPage get the value of the quantity of items per page
func (p *page) GetItemsPerPage() uint64 { return p.itemsPerPage }

// GetCurrentPage get the value of the current page
func (p *page) GetCurrentPage() uint64 { return p.page }

func (p *page) getPage() *page { return p }

func (p *page) calculate() {
	p.page++
	p.offset = (p.page - 1) * p.itemsPerPage
}
