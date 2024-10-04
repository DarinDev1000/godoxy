# go-proxy

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=yusing_go-proxy&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=yusing_go-proxy)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=yusing_go-proxy&metric=ncloc)](https://sonarcloud.io/summary/new_code?id=yusing_go-proxy)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=yusing_go-proxy&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=yusing_go-proxy)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=yusing_go-proxy&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=yusing_go-proxy)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=yusing_go-proxy&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=yusing_go-proxy)
[![](https://dcbadge.limes.pink/api/server/umReR62nRd)](https://discord.gg/umReR62nRd)

一個輕量化、易用且[高效]([docs/benchmark_result.md](https://github.com/yusing/go-proxy/wiki/Benchmarks)))的反向代理和端口轉發工具

## 目錄

<!-- TOC -->

- [go-proxy](#go-proxy)
  - [目錄](#目錄)
  - [重點](#重點)
  - [入門指南](#入門指南)
    - [安裝](#安裝)
    - [命令行參數](#命令行參數)
    - [環境變量](#環境變量)
    - [VSCode 中使用 JSON Schema](#vscode-中使用-json-schema)
  - [展示](#展示)
    - [idlesleeper](#idlesleeper)
  - [源碼編譯](#源碼編譯)

## 重點

-   易用
    -   不需花費太多時間就能輕鬆配置
    -   支持多個docker節點
    -   除錯簡單
-   自動配置 SSL 證書（參見[可用的 DNS 供應商](https://github.com/yusing/go-proxy/wiki/Supported-DNS%E2%80%9001-Providers)）
-   透過 Docker 容器自動配置
-   容器狀態變更時自動熱重載
-   **idlesleeper** 容器閒置時自動暫停/停止，入站時自動喚醒 (可選, 參見 [展示](#idlesleeper))
-   HTTP(s) 反向代理
-   [HTTP middleware](https://github.com/yusing/go-proxy/wiki/Middlewares)
-   [自訂 error pages](https://github.com/yusing/go-proxy/wiki/Middlewares#custom-error-pages)
-   TCP/UDP 端口轉發
-   Web 面板 (內置App dashboard)
-   支持 linux/amd64、linux/arm64 平台
-   使用 **[Go](https://go.dev)** 編寫

[🔼 返回頂部](#目錄)

## 入門指南

### 安裝

1. 抓取Docker鏡像

    ```shell
    docker pull ghcr.io/yusing/go-proxy:latest
    ```

2. 建立新的目錄，並切換到該目錄，並執行
   
   ```shell
    docker run --rm -v .:/setup ghcr.io/yusing/go-proxy /app/go-proxy setup
    ```

3. 設置 DNS 記錄，例如：

    - A 記錄: `*.y.z` -> `10.0.10.1`
    - AAAA 記錄: `*.y.z` -> `::ffff:a00:a01`

4. 配置 `docker-socket-proxy` 其他 Docker 節點（如有） (參見 [範例](docs/docker_socket_proxy.md)) 然後加到 `config.yml` 中

5. 大功告成，你可以做一些額外的配置
    - 使用文本編輯器 (推薦 Visual Studio Code [參見 VSCode 使用 schema](#vscode-中使用-json-schema))
    - 或通過 `http://localhost:3000` 使用網頁配置編輯器
    - 詳情請參閱 [docker.md](docs/docker.md)

[🔼 返回頂部](#目錄)

### 命令行參數

| 參數                      | 描述                                                                                  | 示例                                |
| ------------------------- | ------------------------------------------------------------------------------------- | ----------------------------------- |
| 空                        | 啟動代理服務器                                                                        |                                     |
| `validate`                | 驗證配置並退出                                                                        |                                     |
| `reload`                  | 強制刷新配置                                                                          |                                     |
| `ls-config`               | 列出配置並退出                                                                        | `go-proxy ls-config \| jq`          |
| `ls-route`                | 列出路由並退出                                                                        | `go-proxy ls-route \| jq`           |
| `go-proxy ls-route \| jq` |
| `ls-icons`                | 列出 [dashboard-icons](https://github.com/walkxcode/dashboard-icons/tree/main) 並退出 | `go-proxy ls-icons \| grep adguard` |
| `debug-ls-mtrace`         | 列出middleware追蹤 **(僅限於 debug 模式)**                                            | `go-proxy debug-ls-mtrace \| jq`    |

**使用 `docker exec go-proxy /app/go-proxy <參數>` 運行**

### 環境變量

| 環境變量                       | 描述             | 默認             | 格式          |
| ------------------------------ | ---------------- | ---------------- | ------------- |
| `GOPROXY_NO_SCHEMA_VALIDATION` | 禁用 schema 驗證 | `false`          | boolean       |
| `GOPROXY_DEBUG`                | 啟用調試輸出     | `false`          | boolean       |
| `GOPROXY_HTTP_ADDR`            | http 收聽地址    | `:80`            | `[host]:port` |
| `GOPROXY_HTTPS_ADDR`           | https 收聽地址   | `:443`           | `[host]:port` |
| `GOPROXY_API_ADDR`             | api 收聽地址     | `127.0.0.1:8888` | `[host]:port` |

### VSCode 中使用 JSON Schema

複製 [`.vscode/settings.example.json`](.vscode/settings.example.json) 到 `.vscode/settings.json` 並根據需求修改

[🔼 返回頂部](#目錄)


## 展示

### idlesleeper

![idlesleeper](screenshots/idlesleeper.webp)

[🔼 返回頂部](#目錄)

## 源碼編譯

1. 獲取源碼 `git clone https://github.com/yusing/go-proxy --depth=1`

2. 安裝/升級 [go 版本 (>=1.22)](https://go.dev/doc/install) 和 `make`（如果尚未安裝）

3. 如果之前編譯過（go 版本 < 1.22），請使用 `go clean -cache` 清除緩存

4. 使用 `make get` 獲取依賴項

5. 使用 `make build` 編譯

[🔼 返回頂部](#目錄)
