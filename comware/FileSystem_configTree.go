package comware

import "encoding/xml"

type FileSystem struct {
	/* top level
	   FileSystem
	     Files
	       []File
	     Partitions
	       []Partition
	     RecycleBin
	*/
	Files      *Files      `xml:"Files"`
	Partitions *Partitions `xml:"Partitions"`
	RecycleBin *RecycleBin `xml:"RecycleBin"`
}

type Files struct {
	XMLName xml.Name `xml:"Files"`
	Files   []File   `xml:"File"`
}

type File struct {
	XMLName     xml.Name `xml:"File"`
	Name        string   `xml:"Name"`
	Size        int      `xml:"Size"`
	Time        string   `xml:"Time"`
	IsDirectory bool     `xml:"IsDirectory"`
}

type Partitions struct {
	XMLName    xml.Name    `xml:"Partitions"`
	Partitions []Partition `xml:"Partition"`
}

type Partition struct {
	XMLName    xml.Name `xml:"Partition"`
	Name       string   `xml:"Name"`
	Total      int      `xml:"Total"`
	Used       int      `xml:"Used"`
	Free       int      `xml:"Free"`
	Bootable   bool     `xml:"Bootable"`
	MountState bool     `xml:"MountState"`
}

type RecycleBin struct {
	XMLName         xml.Name         `xml:"RecycleBin"`
	RecycleBinFiles []RecycleBinFile `xml:"RecycleBinFile"`
}

type RecycleBinFile struct {
	XMLName    xml.Name     `xml:"RecycleBinFile"`
	Medium     string       `xml:"Medium"`
	TrashFiles []TrashFiles `xml:"TrashFiles"`
}

type TrashFiles struct {
	XMLName  xml.Name `xml:"TrashFiles"`
	FileName string   `xml:"FileName"`
}
