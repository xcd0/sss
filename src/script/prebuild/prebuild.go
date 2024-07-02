package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/hjson/hjson-go"
)

// KeymapLoaded は、HJSONファイルから読み込まれたキーマップデータを表す構造体です。
type KeymapLoaded struct {
	Lang   string            `json:"lang"`
	Define map[string]string `json:"define"`
	Keymap [][]string        `json:"layout"`
}

// templateText は、生成されるGoコードのテンプレートを定義します。
// {{.Lang}}と{{.KeymapCode}}は、後で実際の値に置き換えられます。
const templateText = `// Code generated generated_keymap.go DO NOT EDIT.

package main

import (
	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/{{.Lang}}"
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

	// defineセクションに基づいてキーマップを置換
	km = applyDefines(km)

	// キーマップコードを生成し、ファイルに書き込みます。
	output := generateGoCode(km)
	err := os.WriteFile("generated_keymap.go", []byte(output), 0644)
	if err != nil {
		panic(fmt.Sprintf("Error writing to file: %v", err))
	}

	fmt.Println("Keymap generated successfully!")
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

// generateKeymapCode は、KeymapLoaded構造体からキーマップコードを生成します。
func generateKeymapCode(km KeymapLoaded) string {
	var sb strings.Builder
	for _, layer := range km.Keymap {
		sb.WriteString("\t\t{\n")
		for _, row := range layer {
			sb.WriteString("\t\t\t")
			keys := strings.Split(row, ",")
			for i, key := range keys {
				key = strings.TrimSpace(key)
				if len(key) != 0 {
					if key == "0" {
						sb.WriteString("0")
					} else {
						sb.WriteString(fmt.Sprintf("%s.%s", km.Lang, key))
					}
					if i < len(keys)-1 {
						sb.WriteString(", ")
					} else {
						sb.WriteString(fmt.Sprintf("%s.%s", km.Lang, key))
					}
				} else {
					// 何もしない。
				}
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\t\t},\n")
	}
	return sb.String()
}
