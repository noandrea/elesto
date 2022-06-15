package credential

import (
	_ "embed"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/elesto-dao/elesto/x/did"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed keeper/testdata/schema.json
	schemaOk        string
	schemaOkCompact = []uint8{0x7b, 0x22, 0x24, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x64, 0x72, 0x61, 0x66, 0x74, 0x2d, 0x30, 0x37, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x23, 0x22, 0x2c, 0x22, 0x24, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x22, 0x24, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x3a, 0x22, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x22, 0x2c, 0x22, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x31, 0x2e, 0x30, 0x22, 0x2c, 0x22, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0xf0, 0x9f, 0x85, 0xa1, 0x22, 0x2c, 0x22, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x3a, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x2c, 0x22, 0x75, 0x72, 0x69, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x6a, 0x73, 0x6f, 0x6e, 0x4c, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x50, 0x6c, 0x75, 0x73, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6c, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x70, 0x6c, 0x75, 0x73, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x22, 0x6a, 0x73, 0x6f, 0x6e, 0x4c, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6c, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x22, 0x6a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x22, 0x7d, 0x7d, 0x2c, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x20, 0x2d, 0x20, 0x41, 0x20, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x2d, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x20, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x5b, 0x22, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0x2c, 0x22, 0x69, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x2c, 0x22, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x5b, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x61, 0x72, 0x72, 0x61, 0x79, 0x22, 0x2c, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5d, 0x7d, 0x2c, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x5b, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x61, 0x72, 0x72, 0x61, 0x79, 0x22, 0x5d, 0x2c, 0x22, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x7d, 0x2c, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x5b, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5d, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x5b, 0x22, 0x69, 0x64, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x69, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x7d, 0x2c, 0x22, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x64, 0x61, 0x74, 0x65, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x7d, 0x2c, 0x22, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x5b, 0x22, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x22, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x49, 0x44, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x54, 0x68, 0x65, 0x20, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x49, 0x44, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x3a, 0x20, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x20, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x2c, 0x22, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x22, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x22, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x20, 0x65, 0x6e, 0x64, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x20, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x7a, 0x6c, 0x69, 0x62, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x20, 0x62, 0x69, 0x74, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x5b, 0x22, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5d, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x2c, 0x22, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x3a, 0x22, 0x75, 0x72, 0x69, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x7d, 0x7d, 0x7d, 0x7d, 0x7d}
	//go:embed keeper/testdata/vocab.json
	vocabOk        string
	vocabOkCompact = []uint8{0x7b, 0x22, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x77, 0x33, 0x63, 0x63, 0x72, 0x65, 0x64, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x77, 0x33, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x32, 0x30, 0x31, 0x38, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x23, 0x22, 0x2c, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x72, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x6c, 0x64, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x23, 0x22, 0x2c, 0x22, 0x52, 0x65, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x30, 0x32, 0x30, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x22, 0x7d, 0x2c, 0x22, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x77, 0x33, 0x63, 0x63, 0x72, 0x65, 0x64, 0x3a, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2c, 0x22, 0x40, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x69, 0x64, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x3a, 0x69, 0x64, 0x22, 0x2c, 0x22, 0x40, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x3a, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x22, 0x40, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x22, 0x7d, 0x2c, 0x22, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x40, 0x69, 0x64, 0x22, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2d, 0x69, 0x64, 0x3a, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2c, 0x22, 0x40, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x54, 0x65, 0x78, 0x74, 0x22, 0x7d, 0x7d, 0x7d, 0x2c, 0x22, 0x40, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x31, 0x2e, 0x31, 0x7d, 0x7d}
	//go:embed keeper/testdata/schema.invalid.json
	schemaErr string
	//go:embed keeper/testdata/vocab.invalid.json
	vocabErr string
)

func TestNewCredentialDefinitionFromFile(t *testing.T) {
	type args struct {
		did          did.DID
		publisherDID did.DID
		name         string
		description  string
		isPublic     bool
		isActive     bool
		schemaFile   string
		vocabFile    string
	}
	tests := []struct {
		name    string
		args    args
		want    *CredentialDefinition
		wantErr error
	}{
		{
			"PASS: can create CredentialDefinition",
			args{
				did.DID("did:cosmo:elesto:cd"),
				did.DID("did:cosmos:elesto:publisher"),
				"Credential Definition 1",
				"This is a sample credential",
				true,
				true,
				"keeper/testdata/schema.json",
				"keeper/testdata/vocab.json",
			},
			&CredentialDefinition{
				Id:           "did:cosmo:elesto:cd",
				PublisherId:  "did:cosmos:elesto:publisher",
				Schema:       schemaOkCompact,
				Vocab:        vocabOkCompact,
				Name:         "Credential Definition 1",
				Description:  "This is a sample credential",
				IsPublic:     true,
				SupersededBy: "",
				IsActive:     true,
			},
			nil,
		},
		{
			"FAIL: invalid schema file",
			args{
				did.DID("did:cosmo:elesto:cd"),
				did.DID("did:cosmos:elesto:publisher"),
				"Credential Definition 1",
				"This is a sample credential",
				true,
				true,
				"keeper/testdata/schema.invalid.json",
				"keeper/testdata/vocab.json",
			},
			&CredentialDefinition{
				Id:           "did:cosmo:elesto:cd",
				PublisherId:  "did:cosmos:elesto:publisher",
				Schema:       schemaOkCompact,
				Vocab:        vocabOkCompact,
				Name:         "Credential Definition 1",
				Description:  "This is a sample credential",
				IsPublic:     true,
				SupersededBy: "",
				IsActive:     true,
			}, fmt.Errorf("error reading schema file: unexpected end of JSON input "),
		},
		{
			"FAIL: invalid vocab file",
			args{
				did.DID("did:cosmo:elesto:cd"),
				did.DID("did:cosmos:elesto:publisher"),
				"Credential Definition 1",
				"This is a sample credential",
				true,
				true,
				"keeper/testdata/schema.json",
				"keeper/testdata/vocab.invalid.json",
			},
			&CredentialDefinition{
				Id:           "did:cosmo:elesto:cd",
				PublisherId:  "did:cosmos:elesto:publisher",
				Schema:       schemaOkCompact,
				Vocab:        vocabOkCompact,
				Name:         "Credential Definition 1",
				Description:  "This is a sample credential",
				IsPublic:     true,
				SupersededBy: "",
				IsActive:     true,
			}, fmt.Errorf("error reading vocab file: unexpected end of JSON input "),
		},
		{
			"FAIL: schema file not found",
			args{
				did.DID("did:cosmo:elesto:cd"),
				did.DID("did:cosmos:elesto:publisher"),
				"Credential Definition 1",
				"This is a sample credential",
				true,
				true,
				"keeper/testdata/non-exising.json",
				"keeper/testdata/vocab.json",
			},
			&CredentialDefinition{
				Id:           "did:cosmo:elesto:cd",
				PublisherId:  "did:cosmos:elesto:publisher",
				Schema:       schemaOkCompact,
				Vocab:        vocabOkCompact,
				Name:         "Credential Definition 1",
				Description:  "This is a sample credential",
				IsPublic:     true,
				SupersededBy: "",
				IsActive:     true,
			}, fmt.Errorf("error reading schema file: open keeper/testdata/non-exising.json: no such file or directory "),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, _ := os.Getwd()
			t.Logf("cwd is %s", w)
			got, err := NewCredentialDefinitionFromFile(tt.args.did, tt.args.publisherDID, tt.args.name, tt.args.description, tt.args.isPublic, tt.args.isActive, tt.args.schemaFile, tt.args.vocabFile)
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestNewPublicVerifiableCredential(t *testing.T) {
	type args struct {
		id   string
		opts []PublicVerifiableCredentialOption
	}
	tests := []struct {
		name string
		args args
		want *PublicVerifiableCredential
	}{
		{
			"PASS: base credential",
			args{
				id:   "did:example:1",
				opts: []PublicVerifiableCredentialOption{},
			},
			&PublicVerifiableCredential{
				Context: []string{
					"https://www.w3.org/2018/credentials/v1",
				},
				Id: "did:example:1",
				Type: []string{
					"VerifiableCredential",
				},
				Issuer:            "",
				IssuanceDate:      nil,
				ExpirationDate:    nil,
				CredentialStatus:  nil,
				CredentialSubject: nil,
				Proof:             nil,
			},
		},
		{
			"PASS: base credential",
			args{
				id: "did:example:1",
				opts: []PublicVerifiableCredentialOption{
					WithType("AnotherTypeCredential"),
					WithIssuerDID("did:example:issuer"),
					WithIssuanceDate(time.Date(2022, 2, 24, 4, 10, 41, 0, time.UTC)),
					WithExpirationDate(time.Date(2022, 12, 24, 4, 10, 41, 0, time.UTC)),
					WithContext("https://another.context/1234"),
				},
			},
			&PublicVerifiableCredential{
				Context: []string{
					"https://www.w3.org/2018/credentials/v1",
					"https://another.context/1234",
				},
				Id: "did:example:1",
				Type: []string{
					"VerifiableCredential",
					"AnotherTypeCredential",
				},
				Issuer:            "did:example:issuer",
				IssuanceDate:      func() *time.Time { v := time.Date(2022, 2, 24, 4, 10, 0, 0, time.UTC); return &v }(),
				ExpirationDate:    func() *time.Time { v := time.Date(2022, 12, 24, 4, 10, 0, 0, time.UTC); return &v }(),
				CredentialStatus:  nil,
				CredentialSubject: nil,
				Proof:             nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewPublicVerifiableCredential(tt.args.id, tt.args.opts...), "NewPublicVerifiableCredential(%v, %v)", tt.args.id, tt.args.opts)
		})
	}
}

func TestNewWrappedPublicCredentialFromFile(t *testing.T) {
	tests := []struct {
		name           string
		credentialFile string
		wantWc         *WrappedCredential
		wantErr        error
	}{
		{
			"PASS: credential file is correct",
			"keeper/testdata/dummy.credential.json",
			&WrappedCredential{
				PublicVerifiableCredential: &PublicVerifiableCredential{
					Id: "https://test.xyz/credential/1",
					Context: []string{
						"https://www.w3.org/2018/credentials/v1",
						"https://resolver.cc/context/did:cosmos:elesto:dummy",
					},
					Type: []string{
						"VerifiableCredential",
						"DummyCredential",
					},
					Issuer:            "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
					IssuanceDate:      func() *time.Time { v := time.Date(2022, 6, 2, 14, 13, 0, 0, time.UTC); return &v }(),
					CredentialSubject: []byte{123, 34, 97, 103, 101, 34, 58, 34, 52, 50, 34, 44, 34, 105, 100, 34, 58, 34, 100, 105, 100, 58, 99, 111, 115, 109, 111, 115, 58, 107, 101, 121, 58, 101, 108, 101, 115, 116, 111, 49, 55, 116, 56, 116, 51, 116, 54, 97, 54, 118, 112, 103, 107, 54, 57, 112, 101, 114, 102, 121, 113, 57, 51, 48, 53, 57, 51, 115, 97, 56, 100, 110, 102, 108, 57, 56, 109, 114, 34, 44, 34, 110, 97, 109, 101, 34, 58, 34, 65, 114, 116, 104, 117, 114, 32, 68, 101, 110, 116, 34, 125},
				},
				CredentialSubject: map[string]interface{}{
					"age":  "42",
					"id":   "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
					"name": "Arthur Dent",
				},
			},
			nil,
		},
		{
			"FAIL: file not found",
			"keeper/testdata/non-existing-file.json",
			nil,
			fmt.Errorf("open keeper/testdata/non-existing-file.json: no such file or directory"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWc, err := NewWrappedPublicCredentialFromFile(tt.credentialFile)

			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantWc, gotWc)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}

func TestNewWrappedCredential(t *testing.T) {

	tests := []struct {
		name    string
		credFn  func() (*WrappedCredential, *PublicVerifiableCredential)
		wantErr error
	}{
		{
			"PASS: credential file is correct",
			func() (*WrappedCredential, *PublicVerifiableCredential) {

				pvc := &PublicVerifiableCredential{
					Context: []string{
						"https://www.w3.org/2018/credentials/v1",
						"https://resolver.cc/context/did:cosmos:elesto:dummy",
					},
					Id: "https://test.xyz/credential/1",
					Type: []string{
						"VerifiableCredential",
						"DummyCredential",
					},
					Issuer:            "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
					IssuanceDate:      func() *time.Time { v := time.Date(2022, 6, 2, 14, 13, 0, 0, time.UTC); return &v }(),
					CredentialSubject: []byte{123, 34, 97, 103, 101, 34, 58, 34, 52, 50, 34, 44, 34, 105, 100, 34, 58, 34, 100, 105, 100, 58, 99, 111, 115, 109, 111, 115, 58, 107, 101, 121, 58, 101, 108, 101, 115, 116, 111, 49, 55, 116, 56, 116, 51, 116, 54, 97, 54, 118, 112, 103, 107, 54, 57, 112, 101, 114, 102, 121, 113, 57, 51, 48, 53, 57, 51, 115, 97, 56, 100, 110, 102, 108, 57, 56, 109, 114, 34, 44, 34, 110, 97, 109, 101, 34, 58, 34, 65, 114, 116, 104, 117, 114, 32, 68, 101, 110, 116, 34, 125},
				}

				wc := &WrappedCredential{
					PublicVerifiableCredential: pvc,
					CredentialSubject: map[string]interface{}{
						"age":  "42",
						"id":   "did:cosmos:key:elesto17t8t3t6a6vpgk69perfyq930593sa8dnfl98mr",
						"name": "Arthur Dent",
					},
				}
				return wc, pvc
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			wantWc, cred := tt.credFn()
			gotWc, err := NewWrappedCredential(cred)

			if tt.wantErr != nil {
				fmt.Println(err.Error())
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, wantWc, gotWc)
			}
		})
	}
}