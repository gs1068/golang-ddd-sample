# Golang DDD Sample

use DDD Architecture

## interface

クライアントからのリクエストを受け取る。
usecase 層からのレスポンスをユーザーに返す。
リクエストやレスポンスが変わってもここだけ修正すれば済むようになる

## usecase↑

interface 層からのリクエストを捌く。（参照、保存、削除など）
基本的に domain の repository を介してアクセスすることで、domain 層に依存させる

## domain↑

DDD の中心となる部分。
ドメインのルールやデータの加工を担う。
domain 層には技術的な関心ごとを実装しな
　 →DB の種類、ORM など

## infrastructure↑

技術基盤へのアクセス。
データを永続化するために DB の操作などを担う。
DB や ORM の種類が変更されてもこの層のみの変更で対応できるように実装する

# File Tree

```
src
├── api(mainファイル)
│   └── main.go
├── config
│   └── database.go
├── domain(中心となる部分)
│   ├── model:構造体やモデルがどんなものかが書かれる
│   │   └── user.go
│   └── repository:DBとのやりとりやCRUDが書かれる
│       └── user.go
├── infra
│   └── user.go
├── interface
│   ├── handler
│   │   └── user.go
│   └── router
│       └── user.go
└── usecase
    └── user.go
```

# Requirement

- Golagn 1.16
  - Echo
  - Gorm

# Installation

Docker コンテナ立ち上げ

```bash
docker-compose build
docker-compose up -d
```

Database 立ち上げと Seed 作成

```bash
make remake
make migrate
```

# Usage

Echo サーバー立ち上げ(fresh によるホットリロードが効くためファイル変更時に立ち上げ直す必要なし)

```bash
make main
```
