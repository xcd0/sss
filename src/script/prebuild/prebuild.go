package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/hjson/hjson-go"
)

// KeymapLoaded は、HJSONファイルから読み込まれたキーマップデータを表す構造体です。
type KeymapLoaded struct {
	Lang   string            `json:"lang"`
	Keymap [][]string        `json:"layout"`
	Define map[string]string `json:"define"`
}

// templateText は、生成されるGoコードのテンプレートを定義します。
// {{.Lang}}と{{.KeymapCode}}は、後で実際の値に置き換えられます。
const templateText = `// Code generated generated_keymap.go DO NOT EDIT.

package main

import (
	//keyboard "github.com/sago35/tinygo-keyboard"
	//"github.com/sago35/tinygo-keyboard/keycodes/{{.Lang}}"
	keyboard "github.com/xcd0/tinygo-keyboard"
	"github.com/xcd0/tinygo-keyboard/keycodes/{{.Lang}}"
)

func GetKeycodes() [][]keyboard.Keycode {
	return [][]keyboard.Keycode{
{{.KeymapCode}}	}
}
`

func main() {
	// コマンドライン引数をチェックし、HJSONファイルのパスを取得します。
	if len(os.Args) < 2 {
		panic("Usage: go run main.go <path_to_keymap.hjson>")
	}

	hjsonPath := os.Args[1]
	km := loadKeymap(hjsonPath)

	{ // 解析前処理。
		// Defineから空白を削除する。
		// またコメント//#以降を削除する。
		for i, _ := range km.Define {
			km.Define[i] = DeleteSpace(km.Define[i])
			km.Define[i] = reComment.ReplaceAllString(km.Define[i], "")
			//log.Printf("define:%v:%v", i, km.Define[i])
		}
		// キーマップから空白を削除し、行末に,をつける。すでにある場合削除する。
		for i, _ := range km.Keymap {
			for j, _ := range km.Keymap[i] {
				km.Keymap[i][j] = DeleteSpace(km.Keymap[i][j])
				km.Keymap[i][j] = strings.TrimRight(km.Keymap[i][j], ",") + ","
				//log.Printf("%v,%v:%v", i, j, km.Keymap[i][j])
			}
		}
	}
	//log.Printf("%#v", km.Keymap)

	// defineセクションに基づいてキーマップを置換する。
	km = applyDefines(km)

	// キーマップコードを生成し、ファイルに書き込みます。
	output := generateGoCode(km)
	err := os.WriteFile("generated_keymap.go", []byte(output), 0644)
	if err != nil {
		panic(fmt.Sprintf("Error writing to file: %v", err))
	}

	fmt.Println("generated_keymap.go generated successfully!")
}

// generateKeymapCode は、KeymapLoaded構造体からキーマップコードを生成します。
func generateKeymapCode(km KeymapLoaded) string {
	var sb strings.Builder
	outputComma := func(key string) {
		//log.Printf("key:%v", key)
		sb.WriteString(",")
	}
	for _, layer := range km.Keymap {
		sb.WriteString("\t\t{\n")
		for r, row := range layer {
			sb.WriteString("\t\t\t")
			row = DeleteSpace(row)
			keys := strings.Split(row, ",")
			for i, key := range keys {
				//key = strings.TrimSpace(key)
				if len(key) != 0 {
					//log.Printf("key:%v", key)
					if i != 0 {
						outputComma(key)
					}
					if key == "0" {
						sb.WriteString("0")
					} else if key[0] == '#' {
						// #で始まる場合数値として解釈する
						sb.WriteString(fmt.Sprintf("%s", key))
					} else if strings.Contains(key, "|") { // |がある場合解釈して書きだす。
						// PIPE: "@KeyBackslash | KeyLeftShift"   // |
						// UNDS: "@KeyBackslash2 | KeyLeftShift"  // _
						// は
						// jp.KeyBackslash | jp.KeyLeftShift
						// jp.KeyBackslash2 | jp.KeyLeftShift
						// のように解釈する。
						//UNDS: "KeyBackslash2 | KeyLeftShift"  // _
						c := strings.Split(key, "|")
						//log.Printf("keys:%#v", c)
						log.Printf("c:%#v", fmt.Sprintf("%s.%s", km.Lang, c[0]))
						for i, _ := range c {
							if i == 0 {
								sb.WriteString(fmt.Sprintf("%s.%s", km.Lang, c[i]))
							} else if len(c[i]) > 0 {
								sb.WriteString(fmt.Sprintf("|%s.%s", km.Lang, c[i]))
								//log.Printf("c:%#v", fmt.Sprintf("%s.%s", km.Lang, c[i]))
							}
						}
					} else {
						// そのまま出力する。
						sb.WriteString(fmt.Sprintf("%s.%s", km.Lang, key))
					}
					if i == len(keys)-1 {
						// レイヤーの末尾で最後に,がない場合ここに来る
						// hjson上で末尾に,がある場合は次のループで空の要素になる。
						sb.WriteString(fmt.Sprintf("%s.%s", km.Lang, key))
					}
				} else {
					if i == len(keys)-1 {
						// hjson上で末尾に,がある場合は次のループで空の要素になる。
					}
					// 何もしない。
				}
			}
			if r < len(layer)-1 {
				sb.WriteString(",\n")
				// 行末
				//log.Printf("key:#1")
			} else {
				// 行末かつレイヤーの最後
				outputComma("#2")
			}
			//log.Printf("---")
		}
		sb.WriteString("\n\t\t},\n")
	}
	return sb.String()
}

var reComment *regexp.Regexp = regexp.MustCompile(`(//|#).*`)
var reSpace *regexp.Regexp = regexp.MustCompile(`\s+`)

func DeleteSpace(str string) string {
	return reSpace.ReplaceAllString(str, "")
}

// loadKeymap は、指定されたパスのHJSONファイルを読み込み、KeymapLoaded構造体に変換します。
func loadKeymap(path string) KeymapLoaded {
	// HJSONファイルを読み込みます。
	hjsonData, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading HJSON file: %v", err))
	}

	// HJSONをJSONに変換します。
	var jsonData interface{}
	err = hjson.Unmarshal(hjsonData, &jsonData)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling HJSON: %v", err))
	}

	// JSONをバイト配列に変換します。
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		panic(fmt.Sprintf("Error marshalling to JSON: %v", err))
	}

	// JSONをKeymapLoaded構造体にアンマーシャルします。
	var km KeymapLoaded
	err = json.Unmarshal(jsonBytes, &km)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling JSON to struct: %v", err))
	}
	return km
}

// applyDefines は、defineセクションに基づいてキーマップ内の値を置換します。
func applyDefines(km KeymapLoaded) KeymapLoaded {
	for i, layer := range km.Keymap {
		for j, row := range layer {
			keys := strings.Split(row, ",")
			for k, key := range keys {
				key = strings.TrimSpace(key)
				if replacement, exists := km.Define[key]; exists {
					keys[k] = replacement
				}
			}
			km.Keymap[i][j] = strings.Join(keys, ", ")
		}
	}
	return km
}

// generateGoCode は、KeymapLoaded構造体を使用してGoコードを生成します。
func generateGoCode(km KeymapLoaded) string {
	// キーマップを返す関数を定義するコードを生成します。
	keymapCode := generateKeymapCode(km)
	// テンプレートを解析します。
	tmpl, err := template.New("keymap").Parse(templateText)
	if err != nil {
		panic(fmt.Sprintf("Error parsing template: %v", err))
	}
	// テンプレートを実行し、結果を文字列ビルダーに書き込みます。
	var buf strings.Builder
	err = tmpl.Execute(&buf, struct {
		Lang       string
		KeymapCode string
	}{
		Lang:       km.Lang,
		KeymapCode: keymapCode,
	})
	if err != nil {
		panic(fmt.Sprintf("Error executing template: %v", err))
	}
	return buf.String()
}
