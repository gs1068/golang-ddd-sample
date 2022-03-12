# Golang DDD Sample

use DDD Architecture

### interface

クライアントからのリクエストを受け取る。
usecase 層からのレスポンスをユーザーに返す。

### usecase↑

interface 層からのリクエストを捌く。（参照、保存、削除など）

### domain↑

DDD の中心となる部分。
ドメインのルールやデータの加工を担う。

### infrastructure↑

技術基盤へのアクセス。
データを永続化するために DB の操作などを担う。

# File Tree

```
src
├── api(mainファイル)
│   └── main.go
├── config
│   └── database.go
├── domain(中心となる部分)
│   ├── model
│   │   └── task.go
│   └── repository
│       └── task.go
├── infra
│   └── task.go
├── interface
│   └── handler
│       ├── router.go
│       └── task.go
└── usecase
    └── task.go
```

# Requirement

- Golagn 1.16
  - Echo 1.0.2
  - Gorm 1.0.2

# Installation

Docker コンテナ立ち上げ

```bash
docker-compose build
docker-compose up -d
```

Database 立ち上げと Seed 作成

```bash
sh remake_databases.sh
sh migrate.sh
```

# Usage

Echo サーバー立ち上げ(fresh によるホットリロードが効くためファイル変更時に立ち上げ直す必要なし)

```bash
sh main.sh
```

# Author

作成情報を列挙する

- Genta Saito
- genta.otias@gmail.com
