package ui

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"

	"github.com/CosmicPredator/chibi/internal"
)

func convertToPNG(jpgData []byte) ([]byte, error) {
	img, _, err := image.Decode(bytes.NewReader(jpgData))

	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func serializeGRCommand(payload []byte, cmd map[string]string) []byte {
	seq := "\033_G"
	first := true
	for k, v := range cmd {
		if !first {
			seq += ","
		}
		first = false
		seq += fmt.Sprintf("%s=%s", k, v)
	}
	if payload != nil {
		seq += ";"
	}
	result := []byte(seq)
	if payload != nil {
		result = append(result, payload...)
	}
	result = append(result, []byte("\033\\")...)
	return result
}


func writeChunked(cmd map[string]string, data []byte) {
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, data)

	chunkSize := 4096
	for len(encoded) > 0 {
		end := chunkSize
		if end > len(encoded) {
			end = len(encoded)
		}
		chunk := encoded[:end]
		encoded = encoded[end:]

		cmdCopy := make(map[string]string)
		for k, v := range cmd {
			cmdCopy[k] = v
		}
		if len(encoded) > 0 {
			cmdCopy["m"] = "1"
		} else {
			cmdCopy["m"] = "0"
		}

		serialized := serializeGRCommand(chunk, cmdCopy)
		os.Stdout.Write(serialized)
		os.Stdout.Sync()
	}
	fmt.Println("")
}

type KGPParams struct {
	R string
	C string
}

func RenderWithImage(imageUrl string, content string, kgpParams KGPParams, numLines int) error {

	resp, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	pngData, err := convertToPNG(imageBytes)
	if err != nil {
		panic(err)
	}

	cmd := map[string]string{
		"a": "T",
		"f": "100",
		"x": "10",
		"r": kgpParams.R,
		"c": kgpParams.C,
		"z": "1",
	}

	fmt.Print("\n")
	fmt.Println(content)
	if internal.CanSupportKittyGP() {
		fmt.Printf("\033[%dA\033[2C", numLines)
		writeChunked(cmd, pngData)
		fmt.Print("\n")
	}

	return nil
}