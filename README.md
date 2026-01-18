# slctl-contacts

The [slctl](https://github.com/softleader/slctl) plugin to view the contacts details in SoftLeader organization

## Install

```sh
$ slctl plugin install github.com/softleader/slctl-contacts
```

## Usage

列出所有公司員工通訊錄

```sh
$ slctl contacts
```

可以使用員工姓名(模糊查詢), 或員工編號(完整查詢)過濾資料

```sh
$ slctl contacts matt
$ slctl contacts 33
```

預設是垂直顯示通訊錄資料, 傳入 `--horizontal` 可改成水平顯示

```sh
$ slctl contacts -H
```

傳入 `--all` 可以查詢包含非 active 的員工通訊錄, e.g. 已離職員工

```sh
$ slctl contacts -a
````

## Developer Guide

### Release Process

This project uses [GoReleaser](https://goreleaser.com/) and GitHub Actions for automated releases.

To release a new version:
1.  Push a new tag (e.g., `v0.1.4`).
2.  GitHub Actions will automatically build binaries for Linux, Windows, and macOS (amd64/arm64) and attach them to the GitHub Release.

Manual release via GoReleaser (for testing):
```sh
goreleaser release --snapshot --clean
```

