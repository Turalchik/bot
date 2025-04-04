package invoice

func (invoiceService *DummyInvoiceService) GetSize() uint64 {
	return invoiceService.Size
}
