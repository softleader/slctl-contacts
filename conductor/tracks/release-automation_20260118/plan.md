# 執行計劃 - release-automation_20260118

## 階段 1：基礎配置與環境準備 [checkpoint: ea68b14]

- [x] Task: 建立 GoReleaser 配置文件 (.goreleaser.yaml) <!-- bf5af4c -->
    - [ ] 建立基本的 `.goreleaser.yaml` 並設定 `builds` 區段。
    - [ ] 定義 `archive` 格式為 `tar.gz` 並包含 `metadata.yaml`。
- [x] Task: 配置多平台編譯支持 (包含 arm64) <!-- fe05777 -->
    - [ ] 在 `builds` 中加入 `darwin/arm64`。
    - [ ] 確認 LDFLAGS 版本號讀取邏輯。
- [x] Task: 建立 GitHub Actions 工作流檔案 <!-- ff2a873 -->
    - [ ] 建立 `.github/workflows/release.yml` 檔案。
    - [ ] 設定觸發條件為推送 Tag。
- [x] Task: Conductor - User Manual Verification '階段 1：基礎配置與環境準備' (Protocol in workflow.md)

## 階段 2：本地驗證與測試

- [ ] Task: 執行本地 GoReleaser 驗證
    - [ ] 使用 `--snapshot` 模式模擬發佈流程。
    - [ ] 檢查產出的 `_dist` 資料夾內容是否正確。
- [ ] Task: 驗證建置出的執行檔與 metadata
    - [ ] 手動解壓產出的 `.tgz` 檔案，確保包含執行檔與正確的 `metadata.yaml`。
- [ ] Task: Conductor - User Manual Verification '階段 2：本地驗證與測試' (Protocol in workflow.md)

## 階段 3：正式集成與清理

- [ ] Task: 更新 README 或文件（可選）
    - [ ] 更新 Release 流程說明。
- [ ] Task: 清理舊的發佈邏輯（如果不再需要 Makefile 裡的 dist）
    - [ ] 評估是否移除 `Makefile` 中的 `dist` target，或將其改為呼叫 GoReleaser。
- [ ] Task: Conductor - User Manual Verification '階段 3：正式集成與清理' (Protocol in workflow.md)
