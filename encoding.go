package main

type Encoding int

const (
    UNKNOWN Encoding = iota
    UTF1 Encoding = iota
    UTF7 Encoding = iota
    UTF8 Encoding = iota
    UTF16_BE Encoding = iota
    UTF16_LE Encoding = iota
    UTF32_BE Encoding = iota
    UTF32_LE Encoding = iota
)

func (e Encoding) String() string {
    switch e {
    default: return "UNKNOWN"
    case UTF1: return "UTF-1"
    case UTF7: return "UTF-7"
    case UTF8: return "UTF-8"
    case UTF16_BE: return "UTF-16 Big Endian"
    case UTF16_LE: return "UTF-16 Little Endian"
    case UTF32_BE: return "UTF-32 Big Endian"
    case UTF32_LE: return "UTF-32 Little Endian"
    }
}
