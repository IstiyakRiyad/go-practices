package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*(),;/"

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

var passLen int;

var generatePassword = &cobra.Command{
	Use: "gen",
	Short: "Generate a random password.",
	Long: `Generate a long password with a given length`,

	Run: func(cmd *cobra.Command, args []string) {
		pass := RandomString(passLen)

		fmt.Println(pass)
	},
}


func init() {
	rootCmd.AddCommand(generatePassword)

	generatePassword.Flags().IntVarP(&passLen, "length", "l", 10, "Length of the password")
}


