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


## 予定

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


