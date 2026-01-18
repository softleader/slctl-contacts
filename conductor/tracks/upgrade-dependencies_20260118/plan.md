# Implementation Plan: upgrade-dependencies

## Phase 1: 基礎環境檢查與備份 (Preparation)
- [x] Task: 紀錄當前依賴版本（`go.mod`）與測試結果作為基準 (Baseline)
- [x] Task: 確保所有現有測試在升級前皆能通過
- [x] Task: Conductor - User Manual Verification 'Phase 1: 基礎環境檢查與備份' (Protocol in workflow.md)

## Phase 2: 依賴包升級與適配 (Upgrade & Adaptation)
- [x] Task: 升級 `github.com/spf13/cobra` 與 `github.com/gosuri/uitable` 並驗證編譯 (11cdc14)
- [~] Task: 將 `gopkg.in/resty.v1` 升級至 `v2` 並修改相關程式碼以符合新 API (如有變更)
- [ ] Task: 執行 `go mod tidy` 升級所有間接依賴並清理 `go.mod`
- [ ] Task: Conductor - User Manual Verification 'Phase 2: 依賴包升級與適配' (Protocol in workflow.md)

## Phase 3: 最終驗證與清理 (Final Verification)
- [ ] Task: 執行完整測試套件 (`go test ./...`) 並確保測試涵蓋率符合要求 (>80%)
- [ ] Task: 執行靜態分析與 Linter 確保升級後無引入品質問題
- [ ] Task: Conductor - User Manual Verification 'Phase 3: 最終驗證與清理' (Protocol in workflow.md)
