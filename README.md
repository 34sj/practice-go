# これはなに

Go 言語練習用リポジトリ  
ISUCON は Go で書かれたシステムをチューニングする競技らしいので、Go に入門して Go への抵抗を無くすことが目的

入門時の Go のバージョン： 1.19.10

```
ctsetera@DESKTOP-OR5PJU2 ~/D/practice-go (master)> go version
go version go1.19.10 linux/amd64
```

# 内容

- test01 -> "Hello World"を表示します
- test02 -> 外部ライブラリを import して「Go の格言」を表示します
- test03 -> 変数定義＆演算
- test04 -> 変数定義
- test05 -> for
- test06 -> for と array 「愛が重いにゃ～（cv. 鳶沢みさき）」
- test07 -> if と switch
- test08 -> defer
- test09 -> defer
- test10 -> ポインタ
- test11 -> 構造体
- test12 -> array と slice
- test13 -> array と slice
- test14 -> for による slice の操作
- test15 -> map (key value で値を持っとくやつ)
- test16 -> メソッド・クロージャ
- test17 -> レシーバを持ったメソッド
- test18 -> レシーバを持ったメソッド
- test19 -> インターフェースメソッド
- test20 -> エラー
- test21 -> データストリーム
- test22 -> goroutine マルチスレッド
- test23 -> channel
- test24 -> channel の close

# 参考

## test01 test02

チュートリアル: Go 言語をはじめよう | https://qiita.com/muranet/items/bbcfce2f0d31d2fd8417

## test03 ~ test24 (test06 を除く)

A Tour of Go | https://go-tour-jp.appspot.com/list

## test06

Developed by [だるぷ](https://github.com/mitixx)

## 全般

go mod init <ここには何を書く？> | https://teratail.com/questions/217859?sort=1  
Go の workspace を使ってみる | https://goodbyegangster.hatenablog.com/entry/2022/10/11/081836  
【Go】コロンイコール(:=)って何? | https://tektektech.com/go-colon-and-equal/  
Go 言語のスライスで勘違いしやすいこところ | https://qiita.com/Kashiwara/items/e621a4ad8ec00974f025  
Go - メソッドとレシーバ | https://qiita.com/Yuuki557/items/e9f5bdfbbfe92973a05e  
Go の interface がわからない人へ　| https://qiita.com/rtok/items/46eadbf7b0b7a1b0eb08  
Go の goroutine, channel をちょっと攻略！ | https://qiita.com/ta1m1kam/items/fc798cdd6a4eaf9a7d5e
