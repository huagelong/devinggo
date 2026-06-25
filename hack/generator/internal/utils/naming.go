package utils

import (
	"strings"
	"unicode"
)

// ToSnakeCase 将字符串转换为snake_case格式
// 例如: CategoryName -> category_name, HTML -> html
func ToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	result.Grow(len(s) + 5) // 预分配空间

	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			// 检查前一个字符
			prev := rune(s[i-1])

			// 如果前一个字符不是大写或下划线，添加下划线
			if !unicode.IsUpper(prev) && prev != '_' {
				result.WriteRune('_')
			} else if i+1 < len(s) {
				// 如果当前是大写，下一个是小写（驼峰边界），添加下划线
				// 例如: HTMLParser -> html_parser
				next := rune(s[i+1])
				if unicode.IsLower(next) {
					result.WriteRune('_')
				}
			}
		}
		result.WriteRune(unicode.ToLower(r))
	}

	return result.String()
}

// ToCamelCase 将字符串转换为camelCase格式
// 例如: category_name -> categoryName, category-name -> categoryName
func ToCamelCase(s string) string {
	if s == "" {
		return ""
	}

	// 分割字符串（支持下划线和连字符）
	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}

	var result strings.Builder
	result.Grow(len(s))

	// 第一个单词小写
	result.WriteString(strings.ToLower(words[0]))

	// 其余单词首字母大写
	for i := 1; i < len(words); i++ {
		result.WriteString(capitalize(words[i]))
	}

	return result.String()
}

// ToPascalCase 将字符串转换为PascalCase格式
// 例如: category_name -> CategoryName, category-name -> CategoryName
func ToPascalCase(s string) string {
	if s == "" {
		return ""
	}

	words := splitWords(s)
	if len(words) == 0 {
		return ""
	}

	var result strings.Builder
	result.Grow(len(s))

	// 所有单词首字母大写
	for _, word := range words {
		result.WriteString(capitalize(word))
	}

	return result.String()
}

// ToConstCase 将字符串转换为CONST_CASE格式
// 例如: CategoryName -> CATEGORY_NAME, category-name -> CATEGORY_NAME
func ToConstCase(s string) string {
	if s == "" {
		return ""
	}

	// 先转换为snake_case，然后全部大写
	snake := ToSnakeCase(s)
	return strings.ToUpper(snake)
}

// ToKebabCase 将字符串转换为kebab-case格式
// 例如: CategoryName -> category-name
func ToKebabCase(s string) string {
	if s == "" {
		return ""
	}

	snake := ToSnakeCase(s)
	return strings.ReplaceAll(snake, "_", "-")
}

// splitWords 分割字符串为单词数组
// 支持下划线、连字符、驼峰命名
func splitWords(s string) []string {
	if s == "" {
		return nil
	}

	var words []string
	var word strings.Builder

	for i, r := range s {
		switch {
		case r == '_' || r == '-' || r == ' ':
			// 分隔符，保存当前单词
			if word.Len() > 0 {
				words = append(words, word.String())
				word.Reset()
			}
		case unicode.IsUpper(r):
			// 大写字母，可能是新单词的开始
			if i > 0 {
				prev := rune(s[i-1])
				// 如果前一个字符是小写或者下一个字符是小写（驼峰边界）
				if unicode.IsLower(prev) || (i+1 < len(s) && unicode.IsLower(rune(s[i+1]))) {
					if word.Len() > 0 {
						words = append(words, word.String())
						word.Reset()
					}
				}
			}
			word.WriteRune(r)
		default:
			word.WriteRune(r)
		}
	}

	// 添加最后一个单词
	if word.Len() > 0 {
		words = append(words, word.String())
	}

	return words
}

// capitalize 首字母大写
func capitalize(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	// 其余字母小写
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
}
