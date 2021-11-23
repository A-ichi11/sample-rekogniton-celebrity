package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

const path = "image/"

var imageName1 = path + "ガッキー.png"
var imageName2 = path + "星野源.png"
var imageName3 = path + "Jeff Bezos.png"
var imageName4 = path + "Andy Jassy.png"

func main() {
	imageFiles := []string{
		imageName1,
		imageName2,
		imageName3,
		imageName4,
	}

	// AWSセッション作成
	sess := session.Must(session.NewSession())

	// Rekognitionクライアントを作成
	svc := rekognition.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

	for _, image := range imageFiles {
		// 画像ファイルを取得
		file, err := os.Open(image)
		if err != nil {
			log.Fatal(err)
		}
		// 最後に画像ファイルを閉じます
		defer file.Close()

		// 画像ファイルのデータを読み込み
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// svc.RecognizeCelebritiesに渡すパラメータを設定
		params := &rekognition.RecognizeCelebritiesInput{
			Image: &rekognition.Image{
				Bytes: bytes,
			},
		}

		// svc.RecognizeCelebritiesを実行
		result, err := svc.RecognizeCelebrities(params)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(*result.CelebrityFaces[0].Name)
	}

}
