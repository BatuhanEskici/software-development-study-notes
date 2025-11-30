# Go Strings

- A Go string is: immutable sequence of bytes

- It is not: slice of bytes, that sentence is only a representation

```go
type string struct {
    ptr *byte
    len int
}
```

## Rune

A `rune` is an alias for `int32`

UTF-8 uses variable length encoding:

| Character             | UTF-8 bytes | Rune size |
| ----------------      | ----------- | --------- |
| ASCII characters      | 1 byte      | 4 bytes   |
| Many common non-ASCII | 2-3 bytes   | 4 bytes   |
| Some rare ones        | 4 bytes     | 4 bytes   |

If Go stored strings as runes, all ASCII text would suddenly take 4× more memory.

This would kill performance for large strings, logs, JSON, network data, etc.

### Most string operations don’t need Unicode decoding

- slicing (s[10:20])
- comparing (==)
- hashing (map keys)
- len(s) -> number of bytes, not characters

Storing runes would slow all of these down.

### Many languages store strings as bytes

- Rust, Python, JavaScript, C, C++, Lua, Go -> store strings as bytes.

- Only a few languages store strings as UCS-2 or UTF-32 internally (Java, C#, older Python versions).

- Go follows the modern approach.

### Then what is the point of rune?

A rune represents one Unicode code point, not one byte and not one user-visible character.

This is needed when:

- Iterating through “characters”

- Performing Unicode transformations (upper/lower casing)

- Validating UTF-8

- Counting user-visible characters (graphemes)

Example:

```go
for i, r := range "ş" {
    fmt.Println(i, r)
    // Will result: 0 351
}
```

This automatically decodes the UTF-8 into runic code points.

Runes are for Unicode processing, not storage.

## Flow From Character We See To Go Holds

```go
Character 😂
    ↓
Unicode code point = U+1F602
    ↓ (hex -> decimal)
128514 (stored in Go rune)
    ↓ (UTF-8 encode)
F0 9F 98 82
    ↓ (hex -> decimal)
[240 159 152 130] (Go string bytes)
```

**STEP 0 — Understand UTF-8 Rules:**

```go
Range:            UTF-8 bytes:
U+0000–007F       1 byte   0xxxxxxx
U+0080–07FF       2 bytes  110xxxxx 10xxxxxx
U+0800–FFFF       3 bytes  1110xxxx 10xxxxxx 10xxxxxx
U+10000–10FFFF    4 bytes  11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
```

Our value: `U+1F602` is in range U+10000–U+10FFFF -> so UTF-8 uses 4 bytes.

```makefile
Byte1: 11110xxx
Byte2: 10xxxxxx
Byte3: 10xxxxxx
Byte4: 10xxxxxx
```

We now fill all the x positions with the binary form of our code point.

**STEP 1 — Write binary form of the code point:**

1F602 hex = 0001 1111 0110 0000 0010 binary

```makefile
1: 0001 -> 0x2³ + 0x2² + 0x2¹ + 1x2⁰ -> 1
F: 1111 -> 1x2³ + 1x2² + 1x2¹ + 1x2⁰ -> 15
6: 0110 -> 0x2³ + 1x2² + 1x2¹ + 0x2⁰ -> 6
0: 0000 -> 0x2³ + 0x2² + 0x2¹ + 0x2⁰ -> 0
2: 0010 -> 0x2³ + 0x2² + 1x2¹ + 0x2⁰ -> 2
```

Remove leading zeros: `11111011000000010` -> That's 17 bits

**STEP 2 — Fill these bits into the UTF-8 4-byte template:**

UTF-8 pattern for 4-byte chars: `11110xxx 10xxxxxx 10xxxxxx 10xxxxxx`

We must place the 21 bits `(xxx xxxxxx xxxxxx xxxxxx)` into the x positions, from right to left.

So we will add 4 zeros to the beginning of the bits and it will be: `000011111011000000010`

```scss
000 011111 011000 000010
│   │       │       │
3   6       6       6   (total = 21)
```

```makefile
Byte1: 11110xxx -> 11110 000 -> 11110000 -> 0xF0
Byte2: 10xxxxxx -> 10 011111 -> 10011111 -> 0x9F
Byte3: 10xxxxxx -> 10 011000 -> 10011000 -> 0x98
Byte4: 10xxxxxx -> 10 000010 -> 10000010 -> 0x82
```

To understand Binary -> Hex Rule

``11110000 is 8 bits (1 byte) -> 1111(1x2³ + 1x2² + 1x2¹ + 1x2⁰) 0000 (0x2³ + 0x2² + 0x2¹ + 0x2⁰) -> 0xF0``

0x is simply a prefix that means: _“The number that comes after this is written in hexadecimal (base 16)”_

It’s the same way `0b` means binary and `0o` means octal.

So the final UTF-8 bytes are:

```go
F0 9F 98 82
```

Go’s []byte prints them in decimal, like:

```go
[240 159 152 130]
```

So all we’re doing is:

```go
F0 hex -> 15x16¹ + 0x16⁰  -> 240 decimal
9F hex -> 9x16¹  + 15x16⁰ -> 159 decimal
98 hex -> 9x16¹  + 8x16⁰  -> 152 decimal
82 hex -> 8x16¹  + 2x16⁰  -> 130 decimal
```

and if we loop through byte slices:

```go
s := "😂"

for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
    // 240
    // 159
    // 152
    // 130
}
```
