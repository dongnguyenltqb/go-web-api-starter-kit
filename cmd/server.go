package cmd

import (
	"fmt"
	"ganja/pkg/infra"
	"ganja/pkg/server/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "restapi server",
	Run: func(cmd *cobra.Command, args []string) {
		boot()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func boot() {
	conf := viper.GetViper()
	port := conf.GetString(`server.port`)

	// cors config
	allow_host := conf.GetStringSlice(`server.allow_host`)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = allow_host

	// setup infra
	infra.Setup()
	app := gin.Default()
	app.Use(cors.New(corsConfig))

	// setup handler layer
	handler.Setup(app)
	fmt.Println("setup handler successfully...")

	// start server
	app.Run(fmt.Sprintf(`:%v`, port))
}
