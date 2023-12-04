package main

import (
        "fmt"
        "github.com/spf13/cobra"
        "github.com/vmware-tanzu/tanzu-plugin-runtime/config/types"
        "github.com/vmware-tanzu/tanzu-plugin-runtime/log"
        "github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"
        "github.com/vmware-tanzu/tanzu-plugin-runtime/plugin/buildinfo"
        "math/rand"
        "os"
        "time"
)

var descriptor = plugin.PluginDescriptor{
        Name:        "tuna-jp",
        Description: "tuna-jp plugin.",
        Target:      types.TargetGlobal, // <<<FIXME! set the Target of the plugin to one of {TargetGlobal,TargetK8s,TargetTMC}
        Version:     buildinfo.Version,
        BuildSHA:    buildinfo.SHA,
        Group:       plugin.ManageCmdGroup, // set group
}

func main() {
        p, err := plugin.NewPlugin(&descriptor)
        if err != nil {
                log.Fatal(err, "")
        }

        // 新しいCobraコマンドを作成
        var tunaCmd = &cobra.Command{
                Use:   "get",
                Short: "Prints TUNA-JP Message",
                Long:  `Prints TUNA-JP Message to the standard output.`,
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("TUNA-JP Advent Calendar 2023")
                },
        }

        // janken コマンドの作成
        var jankenCmd = &cobra.Command{
                Use:   "janken",
                Short: "Play janken.",
                Long:  `Play janken. Use --hand with one of 'o', 'v', 'w'.`,
        }

        var hand string
        jankenCmd.Flags().StringVarP(&hand, "hand", "", "", "Your hand (o: rock, v: scissors, w: paper)")

        jankenCmd.Run = func(cmd *cobra.Command, args []string) {
                if hand != "o" && hand != "v" && hand != "w" {
                        fmt.Println("Invalid hand. Please use one of 'o', 'v', 'w'.")
                        return
                }

                // ランダムな手を生成
                rand.Seed(time.Now().UnixNano())
                opponent := []string{"o", "v", "w"}[rand.Intn(3)]

                // 結果を出力
                fmt.Printf("Your hand: %s, Tuna's fin: %s\n", hand, opponent)
        }

        // コマンドをプラグインに追加
        p.AddCommands(
                tunaCmd,
                jankenCmd,
        )
        if err := p.Execute(); err != nil {
                os.Exit(1)
        }
}
