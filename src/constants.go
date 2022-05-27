package src

import "os"

var (
	HASH_SIZE = 16

	KEY_PATH = "/key"
	KEY_SIZE = 32

	METADATA_PATH       = "/metadata"
	WORKING_PATH        = "/encryption"
	ENCRYPTED_FILES_DIR = "/files"

	FILE_BUFF_SIZE = 2048
	SALT_SIZE      = 4096

	READ_ONLY_FILE_MODE  os.FileMode = 0444
	READ_WRITE_FILE_MODE os.FileMode = 0666
)
