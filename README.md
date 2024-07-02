# sss

rp2040で作る分割キーボード

- 46キー
- 分割
- カラムスタッガード
- ロータリーエンコーダ
- 狭ピッチ（16.5mm）
- 小型(片手13cmx9cm)
- rp2040
- 3Dプリントケース
- 3Dプリント親指キーキャップ


## 完成予定

![sss v7](https://github.com/xcd0/sss/assets/33729994/6e843bbf-9377-43de-a5ad-7c95e2e4f12b)

## 部品

### rp2040を使用した独自ボード

Raspberry Pi Pico が大きく自作キーボードで使いにくいため、同等のボードを作成した。  
とりあえず名前は安直にRPPslimとした。
[@74th](https://x.com/74th)さんの https://74th.booth.pm/items/3929664 が大変参考になった。  

![image](https://github.com/xcd0/sss/assets/33729994/2b6d9dff-1172-43ec-8f8d-49100ec63f82)

![image](https://github.com/xcd0/sss/assets/33729994/9cd8bfae-748c-45b7-9e03-e5cd26b7529e)

- 30個のGPIO (Raspberry Pi Picoは23個、ProMicroは18個)
- USB-C
- キースイッチより狭い13.3x50mm (Raspberry Pi Picoは21x51mm)
	- GPIOが余分な場合、カットして13.3x30mmまで小さくして使用できる。
	- 上の写真の左は33mmでカットしている。この状態で17個GPIOを使用できる。
- ただし、GNDの数が少なく、両面実装で分厚く(5mm)、おそらく信頼性が低い。

普通rp2040を使用する自作キーボードは、キーボード本体にrp2040をそのまま実装する場合が多い。  
このsssではrp2040を別基板とし、キーボードに対して垂直に配置することでキーボードのサイズを小さくしている。  
また、別基板とすることで、別のキーボードを作る際にも使いまわせる、rp2040の部分に不具合があっても簡単に置き換えられる、ようにした。
 
### ロータリーエンコーダ

[アリエク](https://ja.aliexpress.com/item/1005006333962313.html)で売っているEVQWGD001という水平のロータリーエンコーダを使用。

### 親指用キーキャップ

 [booth](https://xcd0.booth.pm/items/4510462)に転がしている親指キーキャップの高さを少し高くしたものを使用。

### その他のキーキャップ

[Talp](https://talpkeyboard.net/items/5f5444c380933970d139e98c)さんで売っている0.8Uのキーキャップ。  
大体いつもこれを使っている。

## 設計
handwire予定だったが基板も作ることにした。choc対応をやめたら基板がすっきりして作りやすかった。  
### 回路図
![image](https://github.com/xcd0/sss/assets/33729994/b4ac3c5c-4dd2-4a7d-a95c-e1884c2d8298)
### PCB
![image](https://github.com/xcd0/sss/assets/33729994/fdd18971-0574-4c9c-b05b-513c4f5dfe5d)

## ファームウェアのビルド

[tinygo](https://tinygo.org/)と[tinygo-keyboard](https://github.com/sago35/tinygo-keyboard)を使用している。  

Linuxはディストリビューションごとに方法があるので、<https://tinygo.org/getting-started/install/linux/>を参照。  
Macはhomebrewで入る模様。未確認。<https://tinygo.org/getting-started/install/macos/>  

windowsの場合、<https://tinygo.org/getting-started/install/windows/>にある通り、  
[scoop](https://scoop.sh/)経由でコンパイラをインストールするのが良い。

scoopはpowershellで下記のようにインストールできる。
```pwsh
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
```

後は下記のコマンドで、

```pwsh
scoop install go
scoop install tinygo
scoop install make
```
でコンパイラとmakeをインストールし、

```pwsh
cd src
make
```

でコンパイルできる。


## ファームウェアの書き込み

makefileに書いているため、

```
make flash
```

でビルドしつつ書き込むことができる。  
windowsのwslでは書き込めないので注意。  
powershellやmsysからも書き込める。  
BOOTボタンを押しながら抜き差しする必要はない。  


## キーマップ

goのソースコードに定義するとgofmtでフォーマットされて整列した状態を保てないため、  
[hjson形式](https://hjson.github.io/)で記述し、それをビルド前に解釈してキーマップを返す関数を自動生成するようにしている。  
<./src/script/prebuild/prebuild.go>にビルド前に実行されるコードがある。  

- キーマップ内のキーコードは<https://github.com/sago35/tinygo-keyboard/tree/main/keycodes>にあるものを使用する。  
- `lang`に定義した言語を表す文字列がそのままpackage名として使用される。まあ現状jpしかないけれども。  
- 短縮表現を`define`に定義できるようにした。
- `layout`にはレイヤーごとにキーコードを記載する。
	- これは単純に、`[][]string`として読み込まれ、`,`と`\n`区切りの単純な`[][][]string`に解釈される。

```hjson
# キーマップで使用する言語
lang: jp

# キーマップで使用するキーコードの短縮形定義
define: {
	LSFT: KeyLeftShift
	LCTL: KeyLeftCtrl
	LALT: KeyLeftAlt
}

# キーマップ
# レイヤーごとに記述する。
layout: [
	[ # layer 0
		KeyEsc,      Key1,        Key2,        Key3,        Key4,        Key5,        Key6,      Key7,        Key8,        Key9,        Key0,        KeyMinus,    KeyHat,        KeyBackslash,
		KeyTab,      KeyQ,        KeyW,        KeyE,        KeyR,        KeyT,        KeyY,      KeyU,        KeyI,        KeyO,        KeyP,        KeyAt,       KeyLeftBrace,  KeyBackspace,
		LCTL,        KeyA,        KeyS,        KeyD,        KeyF,        KeyG,        KeyH,      KeyJ,        KeyK,        KeyL,        KeySemicolon,KeyColon,    KeyRightBrace, KeyEnter,
		LSFT,        KeyZ,        KeyX,        KeyC,        KeyV,        KeyB,        KeyN,      KeyM,        KeyComma,    KeyPeriod,   KeySlash,    KeyBackslash2,KeyUp,        KeyRightShift,
		KeyMod1,     KeyHankaku,  KeyWindows,  LALT,        KeyMuhenkan, KeySpace,    KeySpace,  KeyHenkan,   KeyHiragana, KeyLeftAlt,  KeyMod1,     KeyLeft,     KeyDown,       KeyRight,
	]
```

