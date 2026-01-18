# 技術棧 - slctl-contacts

## 1. 核心開發工具
- **程式語言**：Go
- **依賴管理**：Go Modules
- **建置工具**：GoReleaser (via Makefile wrapper)
- **CI/CD**：GitHub Actions

## 2. 關鍵程式庫
- **CLI 框架**：`github.com/spf13/cobra` - 用於構建結構清晰的命令列應用程式。
- **表格呈現**：`github.com/gosuri/uitable` - 負責格式化並對齊通訊錄資料。
- **HTTP 客戶端**：`github.com/go-resty/resty/v2` - 用於與後端 API 進行通訊以獲取員工資料。

## 3. 環境與部署
- **執行環境**：作為 `slctl` 的插件執行。
- **交叉編譯**：透過 GoReleaser 支援 Linux, Windows, 以及 macOS (Darwin amd64/arm64) 平台的打包。
