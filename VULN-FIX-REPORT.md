# Vulnerability Fix Report

**Date:** 2025-07-20
**Repo:** pujabharti7/terraform-provider-ibm
**Language:** Go (go1.26.3 → toolchain go1.26.5 required)
**Alert sources:** govulncheck (local scan — COMPSEC/Dependabot unavailable; `gh` CLI not installed)

## Fix Status

| ID | Severity | Type | Package | Found | Fixed in | Status |
|---|---|---|---|---|---|---|
| GO-2026-5856 | HIGH | stdlib | `crypto/tls` | go1.26.3 | go1.26.5 | ✅ Fixed — toolchain go1.26.5 added |
| GO-2026-5039 | HIGH | stdlib | `net/textproto` | go1.26.3 | go1.26.4 | ✅ Fixed — toolchain go1.26.5 added |
| GO-2026-5038 | MEDIUM | stdlib | `mime` | go1.26.3 | go1.26.4 | ✅ Fixed — toolchain go1.26.5 added |
| GO-2026-5037 | MEDIUM | stdlib | `crypto/x509` | go1.26.3 | go1.26.4 | ✅ Fixed — toolchain go1.26.5 added |
| GO-2026-5018 | HIGH | module | `golang.org/x/crypto` | v0.51.0 | v0.52.0 | ✅ Fixed — upgraded to v0.52.0 |
| GO-2026-4550 | HIGH | module | `github.com/cloudflare/circl` | v1.6.1 | v1.6.3 | ✅ Fixed — upgraded to v1.6.3 |

## Vulnerability Details

### GO-2026-5856 — Encrypted Client Hello privacy leak in crypto/tls
- **Advisory:** https://pkg.go.dev/vuln/GO-2026-5856
- **Impact:** ECH privacy leak via TLS handshake — affects all TLS connections
- **Fix:** `toolchain go1.26.5` in go.mod

### GO-2026-5039 — Arbitrary inputs in errors without escaping in net/textproto
- **Advisory:** https://pkg.go.dev/vuln/GO-2026-5039
- **Impact:** Unescaped user input in error messages via MIME header parsing
- **Fix:** `toolchain go1.26.5` in go.mod

### GO-2026-5038 — Quadratic complexity in mime.WordDecoder.DecodeHeader
- **Advisory:** https://pkg.go.dev/vuln/GO-2026-5038
- **Impact:** DoS via crafted MIME headers
- **Fix:** `toolchain go1.26.5` in go.mod

### GO-2026-5037 — Inefficient hostname parsing in crypto/x509
- **Advisory:** https://pkg.go.dev/vuln/GO-2026-5037
- **Impact:** DoS via certificate hostname verification
- **Fix:** `toolchain go1.26.5` in go.mod

### GO-2026-5018 — Pathological RSA/DSA DoS in golang.org/x/crypto/ssh
- **Advisory:** https://pkg.go.dev/vuln/GO-2026-5018
- **Impact:** DoS via crafted RSA/DSA parameters in SSH key parsing
- **Fix:** `golang.org/x/crypto v0.51.0` → `v0.52.0`

### GO-2026-4550 — Incorrect secp384r1 CombinedMult in github.com/cloudflare/circl
- **Advisory:** https://pkg.go.dev/vuln/GO-2026-4550
- **Impact:** Incorrect elliptic curve calculation
- **Fix:** `github.com/cloudflare/circl v1.6.1` → `v1.6.3`

## Changes Made

| File | Change |
|---|---|
| `go.mod` | Added `toolchain go1.26.5`; upgraded `golang.org/x/crypto` to v0.52.0; upgraded `github.com/cloudflare/circl` to v1.6.3 |
| `go.sum` | Updated checksums for new module versions |

## Validation

```
govulncheck ./...  → No vulnerabilities found ✅
go build ./...     → Success ✅
go vet ./...       → No issues ✅
```

## Action Required

- [ ] Ensure CI/CD environment uses Go ≥ 1.26.5 to benefit from stdlib fixes
- [ ] Update any Dockerfile `FROM golang:` base image to `golang:1.26.5` or later on next image rebuild
- [ ] Verify clean scan after merge and redeployment

## Notes

- COMPSEC/Dependabot issue tracking requires `gh` CLI pointed at `github.ibm.com` — not available in this environment. Vulnerabilities were detected and fixed via local `govulncheck` scan.
- The `go.mod` `go` directive remains at `1.25.8` (minimum language compatibility). The `toolchain` directive pins the minimum required toolchain to `go1.26.5` for security.
