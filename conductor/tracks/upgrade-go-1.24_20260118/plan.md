# 實作計畫：Go 1.24.x 升級與 Go Modules 遷移

## 階段 1：初始化 Go Modules 與依賴管理遷移
- [ ] 任務：備份原有依賴配置並移除 Glide
    - [ ] 備份 `glide.yaml` 與 `glide.lock` (作為參考後移除)
    - [ ] 移除 `vendor/` 目錄
- [ ] 任務：初始化 Go Modules
    - [ ] 執行 `go mod init github.com/softleader/slctl-contacts`
    - [ ] 設定 `go.mod` 中的 Go 版本為 1.24
- [ ] 任務：匯入依賴並進行初步全面升級
    - [ ] 執行 `go get -u ./...` 升級所有套件
    - [ ] 執行 `go mod tidy` 清理依賴
- [ ] 任務：Conductor - 使用者手動驗證 '階段 1' (Protocol in workflow.md)

## 階段 2：解決代碼相容性與編譯錯誤
- [ ] 任務：修復 `cobra` 相關的代碼變更 (若有)
- [ ] 任務：遷移 `resty` 從 v1 到最新版 (v2)
    - [ ] 搜尋 `gopkg.in/resty.v1` 的引用並更新為 `github.com/go-resty/resty/v2`
    - [ ] 調整相關的 API 讀取與設定邏輯
- [ ] 任務：修復其他編譯錯誤
    - [ ] 執行 `go build ./...` 並修正所有 syntax error
- [ ] 任務：Conductor - 使用者手動驗證 '階段 2' (Protocol in workflow.md)

## 階段 3：驗證與品質檢查
- [ ] 任務：執行單元測試
    - [ ] 執行 `go test ./...`
    - [ ] 針對因套件升級而失敗的測試進行修正
- [ ] 任務：更新建置腳本
    - [ ] 檢查並更新 `Makefile`，確保其支援 Go Modules 模式
- [ ] 任務：Lint 檢查與格式化
    - [ ] 執行 `go fmt ./...` 與 `go fix ./...`
- [ ] 任務：Conductor - 使用者手動驗證 '階段 3' (Protocol in workflow.md)
