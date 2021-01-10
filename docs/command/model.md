# Command to generate model

## 背景
DDDに乗っ取ったプロジェクトにおいて，新しいモデルを追加する際，モデルを表現する複数のファイルとモデルを操作するrepository, daoのコードが必要になる．
これらのコードには，ドメインロジックは含まれず，毎回同じ構造であるため，コードを自動生成することで，開発者の生産性を向上することで可能．

## 目的
DDDにおける モデルとそれに付随するコードを，Railsライクなインターフェースで自動生成する．

## 提供機能
- 指定したモデルの，model, entity, view, repository, dao, dao_test の Go のファイルを生成する．

## モジュール設計
```
./civgen-go/model
├── config.go // .civgen-model.yaml を読み込む
├── generator // 実際にGoのファイルを生成するパッケージ
│   ├── dao.go
│   ├── generator.go // interfaceを定義
│   ├── model.go
│   └── repository.go
├── model.go // main, generator を呼び出してコードを生成する．
└── value // 本ライブラリにおけるバリューオブジェクト
    ├── field.go
    ├── filepath.go
    ├── layer.go
    └── pakcage.go
```