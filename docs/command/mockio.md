# Command to generate mockio

## 背景
[Mock](https://github.com/golang/mock)を伴うテストでは，mockのI/Oの値を管理する必要がある．
利用する各MockのI/Oをテストロジックないに記載すると，テストの見通しが悪くなる．
そのため，MockのI/Oをテストケースのフィールドとして持たせる方法がある．([参考](https://devblog.thebase.in/entry/2018/12/04/110000))
しかし，MockのI/Oのstructをテストのたびに定義するのは，コストが高く，複数のファイルで同じstructを定義してしまう可能性があり，変更コストも高くなる．

## 目的
MockのI/Oのstructを，interfaceのコードから自動生成することで，開発者の生産性向上を図る．

## 提供機能
- interfaceが定義されたファイルと，出力先のファイルパスを指定して，Mock I/O の struct が定義されたファイルを生成する．

## モジュール設計
- https://github.com/golang/mock/tree/master/mockgen をforkして作成．
- ファイルの構成は同じで，mockgen.go の一部だけを編集して，出力されるファイルの内容を変えている．
