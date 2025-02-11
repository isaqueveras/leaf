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
	GetStartCursor() uint64
	// GetStartCursor returns the ending ID of the current page in cursor-based pagination.
	GetEndCursor() uint64
	// SetCursor set the starting ID of the current page in cursor-based pagination
	SetCursor(uint64)

	getPage() *page
	calculate()
}

type page struct {
	page         uint64
	itemsPerPage uint64
	offset       uint64

	cursorStart uint64
	cursorEnd   uint64
}

func newPage(itemsPerPage uint64) IPage { return &page{itemsPerPage: itemsPerPage} }

// GetStartCursor returns the starting id of the current page in cursor-based pagination
func (p *page) GetStartCursor() uint64 { return p.cursorStart }

// GetStartCursor returns the ending id of the current page in cursor-based pagination
func (p *page) GetEndCursor() uint64 { return p.cursorEnd }

// SetCursor set the starting id of the current page in cursor-based pagination
func (p *page) SetCursor(cursor uint64) { p.cursorEnd = cursor }

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

	p.cursorStart = p.cursorEnd
	p.cursorEnd += p.itemsPerPage
}
