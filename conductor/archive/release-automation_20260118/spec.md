# 規格書 - release-automation_20260118

## 1. 概述
本軌道的目標是將 `slctl-contacts` 插件的發佈流程從目前的 Makefile 手動打包轉移到 GitHub Actions 並結合 GoReleaser。這將使開發者只需推送 Git Tag 即可自動完成多平台的編譯與發佈，並特別新增對 Apple Silicon (darwin/arm64) 的支援。

## 2. 功能需求
- **GoReleaser 配置**：建立 `.goreleaser.yaml`，定義編譯參數、打包格式與發佈對象。
- **支援平台**：
    - Linux (amd64)
    - Darwin (amd64, arm64)
    - Windows (amd64)
- **GitHub Actions 工作流**：建立 `.github/workflows/release.yml`，在推送 Tag 時觸發 GoReleaser。
- **打包邏輯**：維持現有的打包結構（包含執行檔與 `metadata.yaml`），並打包成 `.tgz` 格式。

## 3. 非功能需求
- **自動化**：發佈過程無需人工干預。
- **安全性**：使用 GitHub Token 進行授權，確保發佈過程安全。

## 4. 驗收標準
- [ ] 成功在專案中加入 `.goreleaser.yaml`。
- [ ] 成功加入 GitHub Actions 工作流。
- [ ] 本地執行 `goreleaser release --snapshot --clean` 能正確產生所有平台的打包檔，且包含核心執行檔與 `metadata.yaml`。
- [ ] 打包檔名稱符合原有規範（例如 `contacts-darwin-arm64-0.1.0.tgz`）。

## 5. 出範圍 (Out of Scope)
- 遷移現有的發佈紀錄。
- 修改 `contacts` 核心業務邏輯。
