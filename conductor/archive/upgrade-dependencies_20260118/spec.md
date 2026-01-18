# Track Spec: upgrade-dependencies

## Overview
此 Track 的目標是全面升級 `slctl-contacts` 專案中的 Go 依賴包。我們將採取激進升級策略，將所有依賴項提升至最新的主版本（Major Version），並確保升級後的程式碼能夠正常編譯且通過現有的測試套件。

## Functional Requirements
- 升級 `go.mod` 中列出的所有直接依賴與間接依賴。
- 特別注意 `gopkg.in/resty.v1` 到 `v2` (或其他最新版) 的潛在重大變更適配。
- 升級 `github.com/spf13/cobra`、`github.com/gosuri/uitable` 等核心函式庫。
- 如果升級導致 API 變更，需同步修改專案程式碼以維持功能運作。

## Non-Functional Requirements
- **穩定性**：升級後系統行為應與升級前保持一致。
- **現代化**：利用最新版函式庫可能帶來的效能改進或安全性修補。

## Acceptance Criteria
- [ ] `go.mod` 與 `go.sum` 已更新至最新版本。
- [ ] 執行 `go build ./...` 成功，無編譯錯誤。
- [ ] 執行 `go test ./...` 成功，所有測試案例皆通過。

## Out of Scope
- 本 Track 不包含 Go SDK 版本的升級（目前已是 1.25.6）。
- 不包含 GitHub Actions Workflow 的架構重構（除非是為了解決相容性問題）。
