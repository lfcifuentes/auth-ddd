package cmd

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(AppSecret)
}

var AppSecret = &cobra.Command{
	Use:   "make:secret",
	Short: "Generate app secret",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("App secret:")
		fmt.Println(generateBase64())
	},
}

// generateBase64 genera un valor base64 aleatorio
func generateBase64() string {
	rand.Seed(time.Now().UnixNano())

	// Genera 32 bytes aleatorios
	bytes := make([]byte, 32)
	rand.Read(bytes)

	// Codifica los bytes en base64
	base64Value := base64.StdEncoding.EncodeToString(bytes)
	return base64Value
}
