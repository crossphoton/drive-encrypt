package src

import "os"

var HASH_SIZE = 16

var KEY_PATH = "/key"
var KEY_SIZE = 32

var METADATA_PATH = "/metadata"
var WORKING_PATH = "/encryption"
var ENCRYPTED_FILES_DIR = "/files"

var FILE_BUFF_SIZE = 2048
var SALT_SIZE = 4096

var READ_ONLY_FILE_MODE os.FileMode = 0444
var READ_WRITE_FILE_MODE os.FileMode = 0666
