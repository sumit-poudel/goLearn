# Go Package Errors Reference

> `var` = package exports it → use `errors.Is`
> `type` = package creates internally → use `errors.As` / `errors.AsType`

---

## database/sql
| Error | Kind | Use |
|---|---|---|
| `sql.ErrNoRows` | `var` | `errors.Is` |
| `sql.ErrTxDone` | `var` | `errors.Is` |
| `sql.ErrConnDone` | `var` | `errors.Is` |

---

## encoding/json
| Error | Kind | Use |
|---|---|---|
| `*json.SyntaxError` | `type` | `errors.As` / `errors.AsType` |
| `*json.UnmarshalTypeError` | `type` | `errors.As` / `errors.AsType` |
| `*json.InvalidUnmarshalError` | `type` | `errors.As` / `errors.AsType` |
| `*json.MarshalerError` | `type` | `errors.As` / `errors.AsType` |

---

## io/fs
| Error | Kind | Use |
|---|---|---|
| `*fs.PathError` | `type` | `errors.As` / `errors.AsType` |
| `fs.ErrNotExist` | `var` | `errors.Is` |
| `fs.ErrExist` | `var` | `errors.Is` |
| `fs.ErrPermission` | `var` | `errors.Is` |
| `fs.ErrClosed` | `var` | `errors.Is` |
| `fs.ErrInvalid` | `var` | `errors.Is` |

---

## os
| Error | Kind | Use |
|---|---|---|
| `*os.PathError` | `type` | `errors.As` / `errors.AsType` |
| `*os.LinkError` | `type` | `errors.As` / `errors.AsType` |
| `*os.SyscallError` | `type` | `errors.As` / `errors.AsType` |
| `os.ErrNotExist` | `var` | `errors.Is` |
| `os.ErrExist` | `var` | `errors.Is` |
| `os.ErrPermission` | `var` | `errors.Is` |
| `os.ErrClosed` | `var` | `errors.Is` |
| `os.ErrNoDeadline` | `var` | `errors.Is` |

---

## context
| Error | Kind | Use |
|---|---|---|
| `context.DeadlineExceeded` | `var` | `errors.Is` |
| `context.Canceled` | `var` | `errors.Is` |

---

## net
| Error | Kind | Use |
|---|---|---|
| `*net.OpError` | `type` | `errors.As` / `errors.AsType` |
| `*net.DNSError` | `type` | `errors.As` / `errors.AsType` |
| `*net.AddrError` | `type` | `errors.As` / `errors.AsType` |
| `net.ErrClosed` | `var` | `errors.Is` |

---

## net/http
| Error | Kind | Use |
|---|---|---|
| `*http.MaxBytesError` | `type` | `errors.As` / `errors.AsType` |
| `http.ErrNoCookie` | `var` | `errors.Is` |
| `http.ErrNoLocation` | `var` | `errors.Is` |
| `http.ErrBodyReadAfterClose` | `var` | `errors.Is` |
| `http.ErrHandlerTimeout` | `var` | `errors.Is` |
| `http.ErrSkipAltProtocol` | `var` | `errors.Is` |

---

## io
| Error | Kind | Use |
|---|---|---|
| `io.EOF` | `var` | `errors.Is` |
| `io.ErrUnexpectedEOF` | `var` | `errors.Is` |
| `io.ErrClosedPipe` | `var` | `errors.Is` |
| `io.ErrNoProgress` | `var` | `errors.Is` |
| `io.ErrShortBuffer` | `var` | `errors.Is` |
| `io.ErrShortWrite` | `var` | `errors.Is` |

---

## strconv
| Error | Kind | Use |
|---|---|---|
| `*strconv.NumError` | `type` | `errors.As` / `errors.AsType` |

---

## Quick Rule
```
package exports var  →  you have the address  →  errors.Is
package creates type →  you only know shape   →  errors.As / errors.AsType
```
