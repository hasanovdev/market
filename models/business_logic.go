package models

type ScanBarcodeReq struct {
	ComingTableId string `json:"coming_table_id"`
	Barcode       string `json:"barcode"`
}
