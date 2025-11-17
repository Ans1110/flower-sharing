package utils

import (
	"html"
	"regexp"
	"strings"
)

// SanitizeString removes potentially dangerous HTML tags and scripts from user input
func SanitizeString(input string) string {
	if input == "" {
		return input
	}

	// HTML escape the entire string
	sanitized := html.EscapeString(input)

	return sanitized
}

// SanitizeHTML allows safe HTML tags while removing dangerous ones
// Use this for content that needs to support basic formatting
func SanitizeHTML(input string) string {
	if input == "" {
		return input
	}

	// Remove script tags and their contents
	scriptRegex := regexp.MustCompile(`(?i)<script[^>]*>[\s\S]*?</script>`)
	sanitized := scriptRegex.ReplaceAllString(input, "")

	// Remove iframe tags
	iframeRegex := regexp.MustCompile(`(?i)<iframe[^>]*>[\s\S]*?</iframe>`)
	sanitized = iframeRegex.ReplaceAllString(sanitized, "")

	// Remove on* event handlers (onclick, onload, etc.)
	eventHandlerRegex := regexp.MustCompile(`(?i)\s*on\w+\s*=\s*["'][^"']*["']`)
	sanitized = eventHandlerRegex.ReplaceAllString(sanitized, "")

	// Remove javascript: protocol
	jsProtocolRegex := regexp.MustCompile(`(?i)javascript:`)
	sanitized = jsProtocolRegex.ReplaceAllString(sanitized, "")

	// Remove data: protocol (can be used for XSS)
	dataProtocolRegex := regexp.MustCompile(`(?i)data:text/html`)
	sanitized = dataProtocolRegex.ReplaceAllString(sanitized, "")

	// Remove object, embed, and applet tags
	objectRegex := regexp.MustCompile(`(?i)<(object|embed|applet)[^>]*>[\s\S]*?</(object|embed|applet)>`)
	sanitized = objectRegex.ReplaceAllString(sanitized, "")

	return sanitized
}

// SanitizeURL validates and sanitizes URLs to prevent javascript: and data: schemes
func SanitizeURL(input string) string {
	if input == "" {
		return input
	}

	input = strings.TrimSpace(input)

	// Block dangerous protocols
	dangerousProtocols := []string{"javascript:", "data:", "vbscript:", "file:"}
	lowerInput := strings.ToLower(input)

	for _, protocol := range dangerousProtocols {
		if strings.HasPrefix(lowerInput, protocol) {
			return ""
		}
	}

	return input
}

// StripTags removes all HTML tags from input
func StripTags(input string) string {
	if input == "" {
		return input
	}

	// Remove all HTML tags
	tagRegex := regexp.MustCompile(`<[^>]*>`)
	stripped := tagRegex.ReplaceAllString(input, "")

	// Decode HTML entities after stripping tags
	stripped = html.UnescapeString(stripped)

	// Re-escape to prevent any remaining XSS
	stripped = html.EscapeString(stripped)

	return stripped
}

// SanitizeUsername ensures usernames don't contain dangerous characters
func SanitizeUsername(username string) string {
	if username == "" {
		return username
	}

	// Remove any HTML tags
	username = StripTags(username)

	// Remove special characters that could be used in attacks
	username = strings.TrimSpace(username)

	// Only allow alphanumeric, underscore, hyphen, and period
	validUsernameRegex := regexp.MustCompile(`[^a-zA-Z0-9_\-.]`)
	username = validUsernameRegex.ReplaceAllString(username, "")

	return username
}

// SanitizeEmail validates and sanitizes email addresses
func SanitizeEmail(email string) string {
	if email == "" {
		return email
	}

	// Remove any HTML tags
	email = StripTags(email)

	// Trim whitespace
	email = strings.TrimSpace(email)

	// Convert to lowercase
	email = strings.ToLower(email)

	return email
}

// DetectXSSPatterns checks if input contains common XSS attack patterns
func DetectXSSPatterns(input string) bool {
	if input == "" {
		return false
	}

	lowerInput := strings.ToLower(input)

	// Common XSS patterns
	xssPatterns := []string{
		"<script",
		"javascript:",
		"onerror=",
		"onload=",
		"onclick=",
		"onmouseover=",
		"<iframe",
		"<object",
		"<embed",
		"eval(",
		"expression(",
		"vbscript:",
		"data:text/html",
	}

	for _, pattern := range xssPatterns {
		if strings.Contains(lowerInput, pattern) {
			return true
		}
	}

	return false
}

// SanitizeMap sanitizes all string values in a map
func SanitizeMap(data map[string]interface{}) map[string]interface{} {
	sanitized := make(map[string]interface{})

	for key, value := range data {
		switch v := value.(type) {
		case string:
			sanitized[key] = SanitizeString(v)
		case map[string]interface{}:
			sanitized[key] = SanitizeMap(v)
		default:
			sanitized[key] = value
		}
	}

	return sanitized
}
